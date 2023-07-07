package main

import (
	"github.com/rivo/tview"
)

type ListItem struct {
	Text    string
	Checked bool
}

func main() {
	app := tview.NewApplication()

	// Create a slice of ListItems with checkboxes.
	items := []*ListItem{
		{Text: "Option 1", Checked: false},
		{Text: "Option 2", Checked: false},
		{Text: "Option 3", Checked: false},
		{Text: "Option 4", Checked: false},
	}

	// Create a new list.
	list := tview.NewList()

	// Function to populate the list based on the current state of the items.
	var populateList func()
	populateList = func() {
		list.Clear()
		for i, item := range items {
			index := i // Capture index for the closure.
			prefix := "☐ "
			if item.Checked {
				prefix = "✔ "
			}
			list.AddItem(prefix+item.Text, "", 0, func() {
				// Toggle the checkbox when the item is selected.
				items[index].Checked = !items[index].Checked
				populateList()
			})
		}
	}

	populateList() // Populate the list at the start.

	// Start the application.
	if err := app.SetRoot(list, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
