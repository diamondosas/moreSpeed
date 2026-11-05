//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type NetStats struct {
	RxBytes uint64 // Download
	TxBytes uint64 // Upload
}
func main(){
	MeasureSpeed()
}
func ReadProcNetDev() (NetStats, error) {
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		return NetStats{}, err
	}
	defer file.Close()

	var total NetStats
	scanner := bufio.NewScanner(file)
	
	// Skip the first 2 header lines
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		if lineCount <= 2 {
			continue
		}
		
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 10 {
			continue
		}
		
		iface := strings.TrimSuffix(parts[0], ":")
		if iface == "lo" { // Ignore loopback
			continue
		}
		
		rx, _ := strconv.ParseUint(parts[1], 10, 64)
		tx, _ := strconv.ParseUint(parts[9], 10, 64)
		
		total.RxBytes += rx
		total.TxBytes += tx
	}
	return total, scanner.Err()
}

// Calculate rate:
func MeasureSpeed() {
	prev, _ := ReadProcNetDev()
	prevTime := time.Now()

	for range time.Tick(1 * time.Second) {
		curr, _ := ReadProcNetDev()
		now := time.Now()
		elapsed := now.Sub(prevTime).Seconds()

		downloadBps := float64(curr.RxBytes-prev.RxBytes) / elapsed
		uploadBps := float64(curr.TxBytes-prev.TxBytes) / elapsed
		fmt.Printf("Download: %d KBps\n", (downloadBps))
		fmt.Printf("Upload: %d KBps\n", uploadBps/1000)
		prev = curr
		prevTime = now
		_ = downloadBps
		_ = uploadBps
	}
}

