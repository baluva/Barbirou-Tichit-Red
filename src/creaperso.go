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
	inv    map[string]int
	po     int
	skill  string
	Tete  string
	Torse string
	Pieds string
	Equipement Equipement
}

type Item struct {
	Name     string
	Price    int
	Quantity int
}

func (p *Perso) Init(nom string, classe string, grade int, pvMAX float64, pv float64, inv map[string]int, po int, skill string,Tete string, Torse string, Pieds string) {
	p.nom = nom
	p.classe = classe
	p.grade = grade
	p.pv = pv
	p.pvMAX = pvMAX
	p.inv = inv
	p.po = po
	p.skill = skill
	p.Tete = Tete
	p.Torse = Torse
	p.Pieds = Pieds
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
	fmt.Printf("Tete : %s\n", personnage.Tete)
	fmt.Printf("Torse : %s\n", personnage.Torse)
	fmt.Printf("Pieds : %s\n", personnage.Pieds)
	p2.Init("SCOTT ALDEN", "ANTI terroriste", 1, 100, 70, map[string]int{"M4":1, "Armure légère":1, "Potion de vie":1}, 100, "skill", "Casque", "Gilet", "Bottes")
	Menu()
}
func Menu() {
	fmt.Println("。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。+゜゜。。゜+゜。。+゜゜")
	fmt.Println("     ")
	fmt.Println("➡️Menu: ")
	fmt.Println("	👺1. Afficher les informations des TERORISTE👺 ")
	fmt.Println("	👮🏻2. Afficher les informations des LAT (👮🏻) ")
	fmt.Println("	🗃️3. Accéder au contenu de l'inventaire DES TERORISTE👺🗃️ ")
	fmt.Println("	🗃️4. Accéder au contenu de l'inventaire DES des LAT (👮🏻)🗃️ ")
	fmt.Println("	💰5. Marchand💰 ")
	fmt.Println("	  6. Forgeron") // ATTENTION icone du forgeron a changer
	fmt.Println("	🔴7. Quitter🔴 ")
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
			fmt.Println("🛒💰CRAFT TIME !!💰🛒") // ATTENTION icone du forgeron a changer
			p1.displayForgeron()

		case "7":
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
	fmt.Println("1 . Potion de vie")
	fmt.Println("2 . Grenade")
	fmt.Println("3 . Couteau")
	fmt.Println("4 . Livre de sort : Rage")
	fmt.Println("5 . Potion de poison")
	fmt.Println("6 . Etoffe militaire")
	fmt.Println("7 . Kevlar")
	fmt.Println("8 . Cuir")
	fmt.Println("9 . Caméra")
	fmt.Println("10 . Augmentation d'inventaire")
	fmt.Println("0 . Quitter")

	var choix int
	fmt.Print("Choisissez un item à acheter : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		if p.po >= 3 { // Vérifiez si le personnage a assez d'argent pour acheter la potion de vie
			p.Addinv("Potion de vie")
			fmt.Println("Vous avez acheté une Potion de vie.")
			p.po = p.po - 3 // Soustrayez le coût de la potion à l'argent du personnage
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cela.")
		}
	case 2:
		if p.po >= 6 { // Vérifiez si le personnage a assez d'argent pour acheter la Grenade
			p.Addinv("Grenade")
			fmt.Println("Vous avez acheté une Grenade.")
			p.po = p.po - 6 // Soustrayez le coût de la potion à l'argent du personnage
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cela.")
		}
	case 3:
		if p.po >= 8 { // Vérifiez si le personnage a assez d'argent pour acheter le Couteau
			p.Addinv("Couteau")
			fmt.Println("Vous avez acheté un Couteau.")
			p.po = p.po - 8 // Soustrayez le coût de la potion à l'argent du personnage
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cela.")
		}
	case 4: /* case 4 a finir de determiner l'utilisation de cet objet
		if p.po >= 3 { // Vérifiez si le personnage a assez d'argent pour acheter la potion de vie
		p.Spellbook()
		p.po = p.po - 25
		}																								*/
	case 5:
		if p.po >= 6 { // Vérifiez si le personnage a assez d'argent pour acheter la Potion de poison
			p.Addinv("Potion de poison")
			fmt.Println("Vous avez acheté une Potion de poison.")
			p.po = p.po - 6
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cela.")
		}
	case 6:
		if p.po >= 4 { // Vérifiez si le personnage a assez d'argent pour acheter l'Etoffe militaire
			p.Addinv("Etoffe militaire")
			fmt.Println("Vous avez acheté une Etoffe militaire.")
			p.po = p.po - 4
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cela.")
		}
	case 7:
		if p.po >= 7 { // Vérifiez si le personnage a assez d'argent pour acheter le Kevlar
			p.Addinv("Kevlar")
			fmt.Println("Vous avez acheté un Kevlar.")
			p.po = p.po - 7
		}
	case 8:
		if p.po >= 3 { // Vérifiez si le personnage a assez d'argent pour acheter le Cuir
			p.Addinv("Cuir")
			fmt.Println("Vous avez acheté un Cuir.")
			p.po = p.po - 3
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cela.")
		}
	case 9:
		if p.po >= 1 { // Vérifiez si le personnage a assez d'argent pour acheter la Camera
			p.Addinv("Caméra")
			fmt.Println("Vous avez acheté un Caméra.")
			p.po = p.po - 1
		}
	case 10:
		if p.po >= 30 {
			p.Addinv("Augmentation d'inventaire")
			fmt.Println("Vous avez acheté une augmentation d'inventaire.")
			p.po = p.po - 30
		} else {
			fmt.Println("Vous n'avez pas assez d'argent pour acheter cela.")
		}
	default:
		fmt.Println("Choix invalide.")
	}
	Menu()
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
		i := 1
	for item, quantity := range p.inv {
		fmt.Printf("%d. %s (Quantité : %d)\n", i, item, quantity)
		i++
	}

	fmt.Println("🔙0. Retour au menu précédent🔙")
	var choix string
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
	var inv map[string]int
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

func (p *Perso) displayForgeron() { // creation du menu forgeron
	fmt.Println("Bienvenue chez le forgeron !:")
	fmt.Println("voici la liste des marchandises")
	fmt.Println("1 . Casque de protection lv1. ")
	fmt.Println("2 . Gilet pare-balle lv1.")
	fmt.Println("3 . Bottes tactiques lv1.")
	fmt.Println("4 . Casque de protection lv2. ")
	fmt.Println("5 . Gilet pare-balle lv2.")
	fmt.Println("6 . Bottes tactiques lv2.")
	var choix int
	fmt.Print("Choisissez un item à fabriquer : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		if p.inv["Caméra"] >= 1 && p.inv["Cuir"] >= 1 {
			p.inv["Caméra"] -= 1
			p.inv["Cuir"] -= 1
			p.Addinv("Casque de protection lv1.")
			fmt.Println("Vous avez fabriqué un Casque de protection lv1.")
			p.po = p.po - 5
		} else {
			fmt.Println("Vous n'avez pas suffisamment de ressources pour fabriquer cela. ")
		}
	case 2:
		if p.inv["Etoffe militaire"] >= 2 && p.inv["Kevlar"] >= 1 {
			p.inv["Etoffe militaire"] -= 2
			p.inv["Kevlar"] -= 1
			p.Addinv("Gilet pare-balle lv1.")
			fmt.Println("Vous avez fabriqué un Gilet pare-balle lv1.")
			p.po = p.po - 5
		} else {
			fmt.Println("Vous n'avez pas suffisamment de ressources pour fabriquer cela. ")
		}
	case 3:
		if p.inv["Etoffe militaire"] > 1 && p.inv["Cuir"] > 1 {
			p.inv["Etoffe militaire"] -= 1
			p.inv["Cuir"] -= 1
			p.Addinv("Bottes tactiques lv1.")
			fmt.Println("Vous avez fabriqué des Bottes tactiques lv1.")
			p.po = p.po - 5
		} else {
			fmt.Println("Vous n'avez pas suffisamment de ressources pour fabriquer cela. ")
		}
	case 4:
		if p.inv["Caméra"] >= 2 && p.inv["Cuir"] >= 2 {
			p.inv["Caméra"] -= 2
			p.inv["Cuir"] -= 2
			p.Addinv("Casque de protection lv2.")
			fmt.Println("Vous avez fabriqué un Casque de protection lv2.")
			p.po = p.po - 10
		} else {
			fmt.Println("Vous n'avez pas suffisamment de ressources pour fabriquer cela. ")
		}
	case 5:
		if p.inv["Etoffe militaire"] >= 4 && p.inv["Kevlar"] >= 2 {
			p.inv["Etoffe militaire"] -= 4
			p.inv["Kevlar"] -= 2
			p.Addinv("Gilet pare-balle lv2.")
			fmt.Println("Vous avez fabriqué un Gilet pare-balle lv2.")
			p.po = p.po - 10
		} else {
			fmt.Println("Vous n'avez pas suffisamment de ressources pour fabriquer cela. ")
		}
	case 6:
		if p.inv["Etoffe militaire"] > 2 && p.inv["Cuir"] > 2 {
			p.inv["Etoffe militaire"] -= 2
			p.inv["Cuir"] -= 2
			p.Addinv("Bottes tactiques lv2.")
			fmt.Println("Vous avez fabriqué des Bottes tactiques lv2.")
			p.po = p.po - 10
		} else {
			fmt.Println("Vous n'avez pas suffisamment de ressources pour fabriquer cela. ")		
		}
	}
	Menu()
}

type EquipementInfo struct {
	Nom     string
	BonusPv int
}

type Equipement struct {
	Tete  EquipementInfo
	Torse EquipementInfo
	Pieds EquipementInfo
}


func (p *Perso) equiperEquipement(emplacement string, equip EquipementInfo) {
	switch emplacement {
	case "Tete":
		p.inv[equip.Nom]--
		ancienEquip := p.Equipement.Tete
		if ancienEquip.Nom != "" {
			var choix string
			fmt.Printf("Il y a déjà un équipement à l'emplacement %s. Voulez-vous le remplacer ? (O/N) :", emplacement)
			fmt.Scan(&choix)
			if choix == "O" || choix == "o" {
				p.inv[ancienEquip.Nom]++
			} else {
				fmt.Println("Opération annulée")
				return
			}
		}
		p.Equipement.Tete = equip
	case "Torse":
		p.inv[equip.Nom]--
		ancienEquip := p.Equipement.Torse
		if ancienEquip.Nom != "" {
			var choix string
			fmt.Printf("Il y a déjà un équipement à l'emplacement %s. Voulez-vous le remplacer ? (O/N) :", emplacement)
			fmt.Scan(&choix)
			if choix == "O" || choix == "o" {
				p.inv[ancienEquip.Nom]++
			} else {
				fmt.Println("Opération annulée")
				return
			}
		}
		p.Equipement.Tete = equip	
	case "Pieds":
			p.inv[equip.Nom]--
			ancienEquip := p.Equipement.Pieds
			if ancienEquip.Nom != "" {
				var choix string
				fmt.Printf("Il y a déjà un équipement à l'emplacement %s. Voulez-vous le remplacer ? (O/N) :", emplacement)
				fmt.Scan(&choix)
				if choix == "O" || choix == "o" {
					p.inv[ancienEquip.Nom]++
				} else {
					fmt.Println("Opération annulée")
					return
				}		
		}
		p.Equipement.Tete = equip
		
	}
}

func (p *Perso) GetEquipement(emplacement string) EquipementInfo {
	switch emplacement {
	case "Tete":
		return p.Equipement.Tete
	case "Torse":
		return p.Equipement.Torse
	case "Pieds":
		return p.Equipement.Pieds
	default:
		return EquipementInfo{}
	}
}

