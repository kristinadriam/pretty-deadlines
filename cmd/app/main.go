package main

import (
	"os"
	"strconv"
	"time"

	db "pretty-deadlines/internal/db/deadline"
	"pretty-deadlines/internal/models"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

var database *db.Database // Предполагается, что у вас есть структура Database в вашем пакете db

func main() {
	var err error
	database, err = db.InitDb() // Инициализация базы данных
	if err != nil {
		dialog.ShowInformation("Ошибка", "Не удалось подключиться к базе данных.", nil)
		os.Exit(1)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Deadline Manager")

	// Стартовое окно
	startWindow := container.NewVBox(
		widget.NewButton("Создать дедлайн", func() {
			createDeadlineWindow(myApp)
		}),
		widget.NewButton("Просмотреть дедлайны", func() {
			viewDeadlines(myApp)
		}),
		widget.NewButton("Выход", func() {
			myApp.Quit()
		}),
	)

	myWindow.SetContent(startWindow)
	myWindow.ShowAndRun()
}

// Функция для создания окна создания дедлайна
func createDeadlineWindow(myApp fyne.App) {
	win := myApp.NewWindow("Создание дедлайна")

	titleEntry := widget.NewEntry()
	titleEntry.SetPlaceHolder("Название дедлайна")

	descEntry := widget.NewEntry()
	descEntry.SetPlaceHolder("Описание дедлайна")

	// Dropdowns для выбора дня, месяца и года
	dayEntry := widget.NewSelect(getDays(), func(day string) {})
	monthEntry := widget.NewSelect(getMonths(), func(month string) {})
	yearEntry := widget.NewSelect(getYears(), func(year string) {})

	// Dropdowns для выбора часа и минуты
	hourEntry := widget.NewSelect(getHours(), func(hour string) {})
	minuteEntry := widget.NewSelect(getMinutes(), func(minute string) {})

	createButton := widget.NewButton("Создать", func() {
		title := titleEntry.Text
		description := descEntry.Text

		day, _ := strconv.Atoi(dayEntry.Selected)
		month, _ := strconv.Atoi(monthEntry.Selected)
		year, _ := strconv.Atoi(yearEntry.Selected)
		hour, _ := strconv.Atoi(hourEntry.Selected)
		minute, _ := strconv.Atoi(minuteEntry.Selected)

		if title == "" || description == "" || day == 0 || month == 0 || year == 0 {
			dialog.ShowInformation("Ошибка", "Пожалуйста, заполните все поля.", win)
			return
		}

		dueDate := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

		deadlineEntry := models.Deadline{
			Title:       title,
			Description: description,
			DueDate:     dueDate,
		}

		// Вставка в базу данных
		err := database.Insert(deadlineEntry)
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
		widget.NewLabel("Выберите дату:"),
		dayEntry,
		monthEntry,
		yearEntry,
		widget.NewLabel("Выберите время:"),
		hourEntry,
		minuteEntry,
		createButton,
		widget.NewButton("Назад", func() {
			win.Close()
		}),
	))
	win.Show()
}

// Функция для отображения всех дедлайнов
func viewDeadlines(myApp fyne.App) {
	// Создание окна для просмотра дедлайнов
	viewWin := myApp.NewWindow("Список дедлайнов")
	// deadlinesList := widget.NewList(
	//	func() int {
	//		deadlines, _ := database.GetAll() // Получите дедлайны из базы данных
	//		return len(deadlines)
	//	},
	//	func() fyne.CanvasObject {
	//		return widget.NewLabel("")
	//	},
	//	func(i int, o fyne.CanvasObject) {
	//		deadlines, _ := database.GetAll() // Получите дедлайны из базы данных
	//		o.(*widget.Label).SetText(deadlines[i].Title + " - " + deadlines[i].DueDate)
	//	},
	//)

	viewWin.SetContent(container.NewVBox(
		widget.NewLabel("Список дедлайнов"),
		// deadlinesList,
		widget.NewButton("Закрыть", func() {
			viewWin.Close()
		}),
	))
	viewWin.Show()
}

// Функции для получения дней, месяцев и лет
func getDays() []string {
	days := make([]string, 31)
	for i := 1; i <= 31; i++ {
		days[i-1] = strconv.Itoa(i)
	}
	return days
}

func getMonths() []string {
	months := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12",
	}
	return months
}

func getYears() []string {
	years := make([]string, 100)
	currentYear := time.Now().Year()
	for i := 0; i < 100; i++ {
		years[i] = strconv.Itoa(currentYear + i)
	}
	return years
}

func getHours() []string {
	hours := make([]string, 24)
	for i := 0; i < 24; i++ {
		hours[i] = strconv.Itoa(i)
	}
	return hours
}

func getMinutes() []string {
	minutes := make([]string, 60)
	for i := 0; i < 60; i++ {
		minutes[i] = strconv.Itoa(i)
	}
	return minutes
}
