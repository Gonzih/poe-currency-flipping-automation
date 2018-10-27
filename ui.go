package main

import (
	"fmt"

	"github.com/marcusolsson/tui-go"
)

const (
	LeftPanelID = iota
	RightPanelID
	CommentPanelID
)

type UIState struct {
	leftPanelID  int
	rightPanelID int
	focusedPanel int
	pairs        []Pair
}

var state = UIState{}

func drawLeftPanel(list *tui.List) {
	list.RemoveItems()

	for _, pair := range state.pairs {
		list.AddItems(pair.of.String())
	}

	list.SetSelected(state.leftPanelID)
}

func drawRightPanel(list *tui.List) {
	list.RemoveItems()

	for _, offer := range state.pairs[state.leftPanelID].ofs {
		label := fmt.Sprintf("%s profit %.2f", offer.String(), state.pairs[state.leftPanelID].of.Profit(offer))
		list.AddItems(label)
	}

	list.SetSelected(state.rightPanelID)
}

func drawCommentPanel(list *tui.List) {
	list.RemoveItems()

	list.AddItems(
		state.pairs[state.leftPanelID].of.ToMessage(leagueToSearch),
		"",
		state.pairs[state.leftPanelID].ofs[state.rightPanelID].ToMessage(leagueToSearch),
	)
}

func renderUI() {
	root := tui.NewVBox()
	main := tui.NewHBox()
	left := tui.NewVBox()
	right := tui.NewVBox()
	comment := tui.NewVBox()

	root.Append(main)
	root.Append(comment)
	main.Append(left)
	main.Append(right)
	left.SetBorder(true)
	right.SetBorder(true)
	comment.SetBorder(true)

	state.pairs = getPairs()

	leftList := tui.NewList()
	left.Append(leftList)
	drawLeftPanel(leftList)

	rightList := tui.NewList()
	right.Append(rightList)
	drawRightPanel(rightList)

	comments := tui.NewList()
	comment.Append(comments)
	drawCommentPanel(comments)

	focusFN := func() {
		switch state.focusedPanel {
		case LeftPanelID:
			state.focusedPanel++
		case RightPanelID:
			state.focusedPanel = 0
		}
	}

	ui := tui.New(root)
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("q", func() { ui.Quit() })
	ui.SetKeybinding("Enter", focusFN)
	ui.SetKeybinding("Right", focusFN)
	ui.SetKeybinding("Left", focusFN)

	ui.SetKeybinding("Up", func() {
		adjustIDs(-1)

		drawLeftPanel(leftList)
		drawRightPanel(rightList)
		drawCommentPanel(comments)
	})

	ui.SetKeybinding("Down", func() {
		adjustIDs(1)

		drawLeftPanel(leftList)
		drawRightPanel(rightList)
		drawCommentPanel(comments)
	})

	if err := ui.Run(); err != nil {
		panic(err)
	}
}

func adjustIDs(direction int) {
	switch state.focusedPanel {
	case LeftPanelID:
		newID := state.leftPanelID + direction
		state.rightPanelID = 0
		if newID >= len(state.pairs) {
			newID = 0
		}
		if newID < 0 {
			newID = len(state.pairs) - 1
		}
		state.leftPanelID = newID
	case RightPanelID:
		newID := state.rightPanelID + direction
		if newID >= len(state.pairs[state.leftPanelID].ofs) {
			newID = 0
		}
		if newID < 0 {
			newID = len(state.pairs[state.leftPanelID].ofs) - 1
		}
		state.rightPanelID = newID
	}
}