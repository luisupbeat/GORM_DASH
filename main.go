package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()

	// Cargar el ícono desde un archivo
	resource, err := fyne.LoadResourceFromPath("icon.png")
	if err != nil {
		panic(err)
	}
	myApp.SetIcon(resource)

	myWindow := myApp.NewWindow("Dashboard")

	// Crear las páginas del contenido
	homeContent := widget.NewLabel("Home Content")
	settingsContent := widget.NewLabel("Settings Content")
	profileContent := widget.NewLabel("Profile Content")

	// Contenedor para cambiar el contenido
	content := container.NewMax(homeContent)
	// Como hago para que no me aparezca el icono y el nombre de la ventana arre
	// Crear el menú lateral
	menu := container.NewVBox(
		widget.NewButton("Home", func() {
			content.Objects = []fyne.CanvasObject{homeContent}
			content.Refresh()
		}),
		widget.NewButton("Settings", func() {
			content.Objects = []fyne.CanvasObject{settingsContent}
			content.Refresh()
		}),
		widget.NewButton("Profile", func() {
			content.Objects = []fyne.CanvasObject{profileContent}
			content.Refresh()
		}),
		widget.NewButton("Form", func() {
			formContent := createForm()
			content.Objects = []fyne.CanvasObject{formContent}
			content.Refresh()
		}),
		widget.NewButton("Form Haciendas", func() {
			formContent := createFormHaciendas()
			content.Objects = []fyne.CanvasObject{formContent}
			content.Refresh()
		}),
	)

	// Dividir la ventana en dos partes
	split := container.NewHSplit(menu, content)
	split.Offset = 0.2 // Configurar el tamaño del menú lateral

	myWindow.SetContent(split)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

func createForm() fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter your name")
	emailEntry := widget.NewEntry()
	emailEntry.SetPlaceHolder("Enter your email")

	submitButton := widget.NewButton("Submit", func() {
		// Manejar la lógica de envío del formulario aquí
		println("Name:", nameEntry.Text)
		println("Email:", emailEntry.Text)
	})

	form := container.NewVBox(
		widget.NewLabel("Please fill out the form below:"),
		widget.NewForm(
			widget.NewFormItem("Name", nameEntry),
			widget.NewFormItem("Email", emailEntry),
		),
		submitButton,
	)

	return form
}

func createFormHaciendas() fyne.CanvasObject {
	// Crear los campos del formulario
	options := []string{"Sí", "No"}
	selectOption := widget.NewSelect(options, nil)

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter the name")
	locationEntry := widget.NewEntry()
	locationEntry.SetPlaceHolder("Enter the location")

	dateEntry := widget.NewEntry()
	dateEntry.SetPlaceHolder("Enter the date")
	dateEntry.Hide()

	harvestDateEntry := widget.NewEntry()
	harvestDateEntry.SetPlaceHolder("Enter the start date of harvest")
	harvestDateEntry.Hide()

	// Configurar el comportamiento del select
	selectOption.OnChanged = func(value string) {
		if value == "Sí" {
			nameEntry.Show()
			locationEntry.Show()
			dateEntry.Hide()
			harvestDateEntry.Hide()
		} else if value == "No" {
			nameEntry.Hide()
			locationEntry.Hide()
			dateEntry.Show()
			harvestDateEntry.Show()
		}
	}

	submitButton := widget.NewButton("Submit", func() {
		if selectOption.Selected == "Sí" {
			println("Name:", nameEntry.Text)
			println("Location:", locationEntry.Text)
		} else if selectOption.Selected == "No" {
			println("Date:", dateEntry.Text)
			println("Harvest Start Date:", harvestDateEntry.Text)
		}
	})

	form := container.NewVBox(
		widget.NewLabel("Form Haciendas"),
		widget.NewForm(
			widget.NewFormItem("Option", selectOption),
			widget.NewFormItem("Name", nameEntry),
			widget.NewFormItem("Location", locationEntry),
			widget.NewFormItem("Date", dateEntry),
			widget.NewFormItem("Start Date of Harvest", harvestDateEntry),
		),
		submitButton,
	)

	// Inicialmente ocultar los campos
	nameEntry.Hide()
	locationEntry.Hide()

	return form
}
