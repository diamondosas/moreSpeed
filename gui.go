package main

import (
	"github.com/go-gui-org/go-gui/gui"
	"github.com/go-gui-org/go-gui/gui/backend"
)

func startUI() {
	w := gui.SimpleWindow("IneedSpeed", 450, 300, &App{}, func(w *gui.Window) {
		w.SetView(mainView)
	})
	backend.Run(w)
}

func mainView(w *gui.Window) gui.View {
	// app := gui.State[App](w)
	altBg := gui.RGB(35, 37, 43)

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
		RawData: [][]string{
			{"Proc Name", "Upload"},
			{"Telegram", "12 + Kbps"},
			{"Youtube", "12 + Kbps"},
		},
	})
}
