package main

import (
	"bufio"
	"fmt"
	"os"
)

type Perso struct {
	nom    string
	classe string
	grade  int
	pvMAX  float64
	pv     float64
	inv    []string
}
type Item struct {
	Name     string
	Price    int
	Quantity int
}

func (p *Perso) Init(nom string, classe string, grade int, pvMAX float64, pv float64, inv []string) {
	p.nom = nom
	p.classe = classe
	p.grade = grade
	p.pv = pv
	p.pvMAX = pvMAX
	p.inv = inv

}

var p1 Perso
var p2 Perso

func main() {

	p1.Init("AHMAD ABOUSAMRA", "Teroriste", 1, 100, 70, []string{"AK47", "Armure légère", "Potion de vie"})
	p2.Init("SCOTT ALDEN", "ANTI terroriste", 1, 100, 70, []string{"M4", "Armure légère", "Potion de vie"})
	Menu()
}
func Menu() {
	fmt.Println("。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。゜+゜。。+゜゜")
	fmt.Println("     ")
	fmt.Println("➡️Menu:")
	fmt.Println("	👺1.Afficher les informations des TERORISTE👺")
	fmt.Println("	👮🏻2.Afficher les informations des LAT (👮🏻)")
	fmt.Println("	🗃️3. Accéder au contenu de l'inventaire DES TERORISTE👺🗃️")
	fmt.Println("	🗃️4. Accéder au contenu de l'inventaire DES des LAT (👮🏻)🗃️")
	fmt.Println("	💰5. Marchand💰")
	fmt.Println("	🔴6. Quitter🔴")
	fmt.Println("。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。゜+゜。。+゜゜")
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	fmt.Print("choisissez une option")
	scanner.Scan()          // lancement du scanner
	choix := scanner.Text() // stockage du résultat du scanner dans une variable
	switch choix {
	case "1":

		p1.displayInfoTERO()

	case "2":
		p2.displayInfoLAT()
	case "3":
		fmt.Println("💂💂INVENTAIRE:💂‍♀️💂‍♀️")
		fmt.Println("👺👺")
		p1.accessInventory()
	case "4":
		fmt.Println("💂💂INVENTAIRE:💂‍♀️💂‍♀️")
		fmt.Println("👮🏻👮🏻")
		p2.accessInventory()

	case "5":
		fmt.Println("🛒💰SHOP TIMEE !!💰🛒")
		p1.displaymarchand()

	case "6":
		fmt.Println("GOOD BYE!!🖖🖖🖖")
	default:
		fmt.Println("")
	}
}

func (p Perso) displayInfoLAT() {
	fmt.Println("------👮🚨🚓🚨👮--------LUTE ANTI TERRORISTE--------👮🚨🚓🚨👮--------")
	fmt.Println("Nom :", p.nom)
	fmt.Println("classe :", p.classe)
	fmt.Println("Grade:", p.grade)
	fmt.Println("Points de vie actuels :", p.pv)
	fmt.Println("Points de vie maximum :", p.pvMAX)
	fmt.Println("inventaire :", p.inv, "potions")
	fmt.Println("")
	var choix string
	fmt.Println("🔙0. Retour au menu précédent🔙")
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	fmt.Print("choisissez une option")
	scanner.Scan()         // lancement du scanner
	choix = scanner.Text() // stockage du résultat du scanner dans une variable
	if choix == "0" {
		Menu()
	}

}
func (p *Perso) displaymarchand() {
	fmt.Println("Bienvenue chez le marchand !:")
	fmt.Println("voici la liste des ")
	fmt.Println("1. Potion de vie (gratuitement)")
	fmt.Println("2. Grenade")
	fmt.Println("3 . couteau")
	fmt.Println("0. Quitter")

	var choix int
	fmt.Print("Choisissez un item à acheter : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		p.Addinv("Potion de vie")
		fmt.Println("Vous avez acheté une Potion de vie.")
		Menu()
	case 2:
		p.Addinv("Grenade")
		Menu()
	case 3:
		p.Addinv("Couteau")
	case 0:
		Menu()
	default:
		fmt.Println("Choix invalide.")
	}
}

func (p *Perso) displayInfoTERO() {
	fmt.Println("☠-----▄︻デ══━一💥----👳🏽‍♂️-----TERORISTE-----▄︻デ══━一💥-----👳🏽‍♂️------☠")
	fmt.Println("Nom :", p.nom)
	fmt.Println("classe :", p.classe)
	fmt.Println("Grade:", p.grade)
	fmt.Println("Points de vie actuels :", p.pv)
	fmt.Println("Points de vie maximum :", p.pvMAX)
	fmt.Println("inventaire :", p.inv, "potions")
	fmt.Println("")
	var choix string
	fmt.Println("🔙0. Retour au menu précédent🔙")
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	fmt.Print("choisissez une option")
	scanner.Scan()         // lancement du scanner
	choix = scanner.Text() // stockage du résultat du scanner dans une variable
	if choix == "0" {
		Menu()
	}

}
func (p *Perso) accessInventory() {
	fmt.Println("Inventaire du Personnage :")
	for i, item := range p.inv {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	var choix string
	fmt.Println("🔙0. Retour au menu précédent🔙")
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	fmt.Print("choisissez une option")
	scanner.Scan()         // lancement du scanner
	choix = scanner.Text() // stockage du résultat du scanner dans une variable
	if choix == "0" {
		Menu()
	}
}
