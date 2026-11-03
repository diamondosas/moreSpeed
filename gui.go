package main

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"sync"
	"time"

	"github.com/go-gui-org/go-gui/gui"
	"github.com/go-gui-org/go-gui/gui/backend"
)
type Process struct{ 
	Name string
	Upload int
	Download int
}

type App struct{
	Processes []Process
}

var appMU sync.Mutex

func startUI() {
	app := initProcesses()
	w := gui.SimpleWindow("moreSpeed", 450, 300, app, func(w *gui.Window) {
		w.SetView(mainView)
		go updateValues(w)
		go updateProcesses(w)
	})
	

	backend.Run(w)
}

func initProcesses() *App {
	initState := &App{
	Processes: []Process{
		{Name: "Telegram", Upload: 10,Download: 10},
		{Name: "Whatsapp", Upload: 10, Download: 10},
	},
	}	
	return initState
}

func updateProcesses(w *gui.Window){
	// runtime.LockOSThread()
	// defer runtime.UnlockOSThread()
	programs := []string{"Telegram","Whatsapp","Discord","Youtube","Reddit","Hackernews"}
	ticker := time.NewTicker(time.Second)
	getNoOfProg:= func() int{
		noOfProg := rand.IntN(len(programs) - 1); if noOfProg <=0 {noOfProg = 1};
		return noOfProg
	}

	for{
		select {
		case<-w.Ctx().Done():
			return
		case<-ticker.C:
			w.QueueCommand(
				func(w *gui.Window) {
					app := gui.State[App](w)
					// app.Processes = nil
					fmt.Println(len(programs))
					fmt.Println(rand.IntN(len(programs) - 1))
					// for i:=0; i < len(app.Processes); i++{
						// 
					for i := range app.Processes{
						noOfProg := getNoOfProg()
						fmt.Println("Put this ", programs[noOfProg])
						app.Processes[i].Name = programs[noOfProg]
						fmt.Println("Putted:", app.Processes[i].Name)
					}
						// process := Process{app.Processes[i].Name,0,0}
						// app.Processes = append(app.Processes, process)

					// }
				})
			w.InvalidateLayout()
		}
	}
} 

func updateValues(w *gui.Window){
	ticker := time.NewTicker(time.Second)
	for{
		select{
		case <-w.Ctx().Done():
			return
		case <-ticker.C:
			w.QueueCommand(func(w *gui.Window) {
				app := gui.State[App](w)
				for i := range app.Processes{
					app.Processes[i].Upload, app.Processes[i].Download = 10 + rand.IntN(400),10 +  rand.IntN(1000)
				}

				//Tell the ui to redrwa everthin in the
				w.InvalidateLayout()
			})
		}
	}
}

func mainView(w *gui.Window) gui.View {
	//Initialise Applicaton State
	app := gui.State[App](w)


	//Update the rows per render
	rows := [][]string{
			{"Proc Name", "Upload", "Downlaod"},
	}

	for _, p := range app.Processes{
		rows = append(rows,[]string{
			p.Name,
			fmt.Sprintf("%d Kbps", p.Upload),
			fmt.Sprintf("%d Kbps", p.Download),
		})	
	}

	altBg := gui.RGB(35, 37, 43)

	//Return the Table 
	return w.Table(gui.TableCfg{
		TextStyleHead: gui.TextStyle{
			Size: 25,
			Family: "Arial",
			Color: gui.Red,
		},
		TextStyle: gui.TextStyle{
			Size: 20,
			Color: gui.White,
			Family: " Arial",
		}.Bold(),

		Sizing:           gui.FillFill,
		BorderStyle:      gui.TableBorderHorizontal,
		SizeBorderHeader: 5,
		FreezeHeader:     true,
		ColumnAlignments: []gui.HorizontalAlign{
			gui.HAlignStart,
			gui.HAlignCenter,
			gui.HAlignCenter,
		},
		ColorRowAlt: &altBg,
		Focusable:   true,
		
		// OnSelect: func(selected map[int]bool, row int, ctx gui.EventCtx),

		ID: "moreSpeed",
		RawData: rows,
	})
}

