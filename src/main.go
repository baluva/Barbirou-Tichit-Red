package main

import (
	"fmt"
	"red/perso"
	"time"

	"github.com/fatih/color"
	"github.com/mbndr/figlet4go"
)

func main() {
	perso.ClearScreen()
	displayhist()
	time.Sleep(10 * time.Second)
	perso.ClearScreen()
	DisplayCSGOWelcome()
	personnage := perso.CharCreation()
	Displaysuc()
	time.Sleep(3 * time.Second)
	personnage.Menu()
}

func DisplayCSGOWelcome() {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("Welcome to CS:GO")
	greenText := "\x1b[32m" + renderStr + "\x1b[0m"
	fmt.Println(greenText)
}

func displayhist() {
	gameTitle := "Counter-Strike: Global Offensive (CS:GO)"
	developers := "développé avec passion par Louey BARBIROU et Johan TIPAU."
	teams := "Trois équipes se font face : les Terroristes, les Contre-Terroristes et les monstrueuses créatures."
	gameplay := "Global Offensive a été accueilli favorablement par les critiques à sa sortie, recevant des éloges pour son gameplay addictif qui reste fidèle à l'esprit de la série légendaire Counter-Strike. Le jeu offre une expérience de jeu riche en stratégie, en coordination d'équipe et en action palpitante."
	yellow := color.New(color.FgYellow).SprintFunc()

	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("%s\n", green("Bienvenue dans l'univers de"))
	time.Sleep(2 * time.Second)
	fmt.Printf("%s\n", yellow(gameTitle))
	time.Sleep(2 * time.Second)
	fmt.Printf("%s\n", yellow(developers))
	time.Sleep(2 * time.Second)
	fmt.Printf("%s\n", blue(teams))
	time.Sleep(2 * time.Second)
	fmt.Printf("%s\n", gameplay)
	time.Sleep(2 * time.Second)
}
func Displaysuc() {
	perso.ClearScreen()
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("personnage créé")
	greenText := "\x1b[32m" + renderStr + "\x1b[0m"
	fmt.Println(greenText)

}
