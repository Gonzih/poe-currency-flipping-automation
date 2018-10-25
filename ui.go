package main

import "github.com/marcusolsson/tui-go"

func renderUI() {
	main := tui.NewHBox()
	left := tui.NewVBox()
	right := tui.NewVBox()
	main.Append(left)
	main.Append(right)
	left.SetBorder(true)
	right.SetBorder(true)

	pairs := getPairs()

	list := tui.NewList()
	left.Append(list)

	for i := 0; i < len(pairs); i++ {
		list.AddItems(pairs[i].String())
	}

	ui := tui.New(main)
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("q", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
