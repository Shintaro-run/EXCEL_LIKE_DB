package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/xuri/excelize/v2"
)

func main() {
	// Create a new Fyne application
	myApp := app.New()
	myWindow := myApp.NewWindow("Excel Search")

	// Create a slice to store selected file paths
	var selectedFiles []string

	// Create a file picker widget
	filePicker := widget.NewButton("Select Excel file(s)", func() {
		fileDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				// User cancelled
				return
			}
			selectedFiles = append(selectedFiles, reader.URI().Path())
		}, myWindow)
		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".xlsx"}))
		fileDialog.Show()
	})

	// Create input field for entering the query word
	queryEntry := widget.NewEntry()
	queryEntry.SetPlaceHolder("Enter query word")

	// Create a list to display search results
	var listData []string
	resultList := widget.NewList(
		func() int {
			return len(listData) // リストデータの項目数
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("") // リスト項目のプレースホルダー
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(listData[lii]) // リストデータを表示
		},
	)

	// Create a button to initiate search
	searchButton := widget.NewButton("Search", func() {
		listData = []string{} // リストデータをクリア

		// Get the query word
		query := queryEntry.Text

		// Iterate through selected files
		for _, file := range selectedFiles {
			// Read Excel file
			f, err := excelize.OpenFile(file)
			if err != nil {
				fmt.Println("Error opening Excel file:", err)
				continue
			}

			// Iterate through all worksheets in the file
			for _, sheet := range f.GetSheetList() {
				rows, err := f.GetRows(sheet)
				if err != nil {
					fmt.Println("Error reading sheet:", err)
					continue
				}

				// Iterate through rows and cells, find matches, and display results
				for rowIndex, row := range rows {
					for cellIndex, cellValue := range row {
						// Check if cell contains the query word
						if strings.Contains(strings.ToLower(cellValue), strings.ToLower(query)) {
							colName, _ := excelize.ColumnNumberToName(cellIndex + 1) // 修正
							result := fmt.Sprintf("File: %s - Sheet: %s - Row: %d - Cell: %s", file, sheet, rowIndex+1, colName)
							listData = append(listData, result) // リストデータに結果を追加
						}
					}
				}
			}
		}

		resultList.Refresh() // リストを更新
	})

	// Create a layout container to arrange widgets
	content := container.NewVBox(
		filePicker,
		queryEntry,
		searchButton,
		resultList,
		// Add additional widgets as needed
	)

	// Set content and show window
	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}
