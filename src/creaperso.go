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
	po     int
	skill  string
}
type Item struct {
	Name     string
	Price    int
	Quantity int
}

func (p *Perso) Init(nom string, classe string, grade int, pvMAX float64, pv float64, inv []string, po int, skill string) {
	p.nom = nom
	p.classe = classe
	p.grade = grade
	p.pv = pv
	p.pvMAX = pvMAX
	p.inv = inv
	p.po = po
	p.skill = skill
}

var p1 Perso
var p2 Perso

func main() {

	personnage := charCreation()
	fmt.Println("Personnage créé avec succès ! Voici ses détails :")
	fmt.Printf("Nom : %s\n", personnage.nom)
	fmt.Printf("Classe : %s\n", personnage.classe)
	fmt.Printf("Grade : %d\n", personnage.grade)
	fmt.Printf("Points de vie max : %.2f\n", personnage.pvMAX)
	fmt.Printf("Points de vie actuels : %.2f\n", personnage.pv)
	fmt.Printf("Inventaire : %v\n", personnage.inv)
	fmt.Printf("Pièces d'or : %d\n", personnage.po)
	fmt.Printf("Compétence : %s\n", personnage.skill)
	p2.Init("SCOTT ALDEN", "ANTI terroriste", 1, 100, 70, []string{"M4", "Armure légère", "Potion de vie"}, 100, "")
	Menu()
}
func Menu() {
	fmt.Println("。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。゜+゜。。+゜゜")
	fmt.Println("     ")
	fmt.Println("➡️Menu: ")
	fmt.Println("	👺1.Afficher les informations des TERORISTE👺 ")
	fmt.Println("	👮🏻2.Afficher les informations des LAT (👮🏻) ")
	fmt.Println("	🗃️3. Accéder au contenu de l'inventaire DES TERORISTE👺🗃️ ")
	fmt.Println("	🗃️4. Accéder au contenu de l'inventaire DES des LAT (👮🏻)🗃️ ")
	fmt.Println("	💰5. Marchand💰 ")
	fmt.Println("	🔴6. Quitter🔴 ")
	fmt.Println("。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。゜+゜。。+゜゜")
	scanner := bufio.NewScanner(os.Stdin) // création du scanner capturant une entrée utilisateur
	var choix string
	for choix != "1" || choix != "2" || choix != "3" || choix != "4" || choix != "5" || choix != "6" {
		fmt.Print("choisissez une option ")
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
	fmt.Println("voici la liste des marchandises")
	fmt.Println("1 . Potion de vie ")
	fmt.Println("2 . Grenade")
	fmt.Println("3 . Couteau")
	fmt.Println("4 . Livre de sort : Rage")
	fmt.Println("5 . Potion de poison")
	fmt.Println("6 . Etoffe militaire")
	fmt.Println("7 . Kevlar")
	fmt.Println("8 . Cuir")
	fmt.Println("9 . Caméra")
	fmt.Println("0 . Quitter")

	var choix int
	fmt.Print("Choisissez un item à acheter : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		if p.CheckInv(){
		p.Addinv("Potion de vie")
		fmt.Println("Vous avez acheté une Potion de vie.")
		p.po = p.po - 3
		Menu()
		}
	case 2:
		if p.CheckInv(){
		p.Addinv("Grenade")
		fmt.Println("Vous avez acheté une Grenade.")
		Menu()
		}
	case 3:
		if p.CheckInv(){
		p.Addinv("Couteau")
		fmt.Println("Vous avez acheté un Couteau.")
		Menu()
		}
	case 4:
		if p.CheckInv(){
		p.Spellbook()
		p.po = p.po - 25
		Menu()
		}
	case 5:
		if p.CheckInv(){
		p.Addinv("Potion de poison")
		fmt.Println("Vous avez acheté une Potion de poison.")
		p.po = p.po - 6
		Menu()
		}
	case 6:
		if p.CheckInv(){
		p.Addinv("Etoffe militaire")
		fmt.Println("Vous avez acheté une Etoffe militaire.")
		p.po = p.po - 4
		Menu()
		}
	case 7:
		if p.CheckInv(){
		p.Addinv("Kevlar")
		fmt.Println("Vous avez acheté un Kevlar.")
		p.po = p.po - 7
		Menu()
		}
	case 8:
		if p.CheckInv(){
		p.Addinv("Cuir")
			fmt.Println("Vous avez acheté un Cuir.")
			p.po = p.po - 3
			Menu()
		}

	case 9:
		p.Addinv("Caméra")
		fmt.Println("Vous avez acheté un Caméra.")

		p.po = p.po - 1
		Menu()
	case 0:
		Menu()
	default:
		fmt.Println("Choix invalide.")
	}
}

func (p *Perso) displayInfoTERO() {
	fmt.Println("☠-----▄︻デ══━一💥----👳🏽‍♂️-----TERORISTE-----▄︻デ══━一💥-----👳🏽‍♂️------☠")
	fmt.Println("Nom : ", p.nom)
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
func charCreation() *Perso {
	var nom string
	for {
		fmt.Print("Entrez votre nom  : ")
		fmt.Scan(&nom)
		if ValideNom(nom) {
			break
		} else {
			fmt.Println("Le nom doit contenir uniquement des lettres.")
		}
	}

	nom = formatNom(nom)

	var classe string
	for {
		fmt.Print("Choisissez votre classe : entre LAT ET TERRO ")
		fmt.Scan(&classe)
		if estValideClasse(classe) {
			break
		} else {
			fmt.Println("Classe invalide. Veuillez choisir une classe valide.")
		}
	}

	var grade int
	for {
		fmt.Print("Entrez votre grade : ")
		_, err := fmt.Scan(&grade)
		if err == nil {
			break
		} else {
			fmt.Println("Grade invalide. Veuillez entrer un nombre entier.")
		}
	}

	pvMAX := 100.0
	switch classe {
	case "LAT":
		pvMAX = 100.0
	case "TERO":
		pvMAX = 80.0
	}
	pv := pvMAX / 2
	var inv []string
	var po int
	var skill string
	return &Perso{
		nom:    nom,
		classe: classe,
		grade:  grade,
		pvMAX:  pvMAX,
		pv:     pv,
		inv:    inv,
		po:     po,
		skill:  skill,
	}
}
func (p *Perso) CheckInv() bool { // creation de la limite de l'inventaire.
	if len(p.inv) < 10 {
		return true
	} else {
		fmt.Println("L'inventaire est plein. Impossible d'ajouter un nouvel objet. ")
		return false
	}
}
