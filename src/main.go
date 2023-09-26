package main

import (
	"fmt"
	"red/perso"

	"github.com/mbndr/figlet4go"
)

func main() {
	perso.ClearScreen()
	DisplayCSGOWelcome()
	personnage := perso.CharCreation()
	fmt.Println("Personnage créé avec succès ! ")
	personnage.Menu()
}

func DisplayCSGOWelcome() {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("Welcome to CS:GO")
	greenText := "\x1b[32m" + renderStr + "\x1b[0m"
	fmt.Println(greenText)
}
