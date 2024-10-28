package main

import (
	"os"

	db "pretty-deadlines/internal/db/deadline"
	"pretty-deadlines/internal/models"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

var database *db.Database // Assuming you have a Database struct in your db package

func main() {
	var err error
	database, err = db.InitDb() // Initialize the database
	if err != nil {
		dialog.ShowInformation("Ошибка", "Не удалось подключиться к базе данных.", nil)
		os.Exit(1)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Deadline Manager")

	// Стартовое окно
	startWindow := container.NewVBox(
		widget.NewButton("Создать дедлайн", func() {
			// Создание окошка для создания дедлайна
			win := myApp.NewWindow("Создание дедлайна")

			titleEntry := widget.NewEntry()
			titleEntry.SetPlaceHolder("Название дедлайна")

			descEntry := widget.NewEntry()
			descEntry.SetPlaceHolder("Описание дедлайна")

			dateEntry := widget.NewEntry()
			dateEntry.SetPlaceHolder("Дата дедлайна (YYYY-MM-DD)")

			createButton := widget.NewButton("Создать", func() {
				title := titleEntry.Text
				description := descEntry.Text
				dateStr := dateEntry.Text

				if title == "" || description == "" || dateStr == "" {
					dialog.ShowInformation("Ошибка", "Пожалуйста, заполните все поля.", win)
					return
				}

				deadlineEntry := models.Deadline{
					Title:       title,
					Description: description,
					DueDate:     dateStr,
				}

				// Вставка в базу данных
				err = database.Insert(deadlineEntry)
				if err != nil {
					dialog.ShowInformation("Ошибка", "Не удалось сохранить дедлайн в базе данных.", win)
					return
				}

				dialog.ShowInformation("Успех", "Дедлайн успешно создан!", win)
				win.Close()
			})

			win.SetContent(container.NewVBox(
				widget.NewLabel("Создание дедлайна"),
				titleEntry,
				descEntry,
				dateEntry,
				createButton,
			))
			win.Show()
		}),
	)

	myWindow.SetContent(startWindow)
	myWindow.ShowAndRun()
}
