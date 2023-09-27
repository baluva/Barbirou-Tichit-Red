package perso

import (
	"bufio"
	"fmt"
	"os"
)

// tous les item disponible dans le jeu
const (
	// Consomable
	PotionDeVie     ItemType = "Potion de vie"
	Grenade         ItemType = "Grenade"
	Couteau         ItemType = "Couteau"
	LivreDeSortRage ItemType = "Livre de sort : Rage"
	PotionDePoison  ItemType = "Potion de poison"
	redText                  = "\x1b[31m"
	reset                    = "\x1b[0m"
	yellowText               = "\x1b[33m"
	resety                   = "\x1b[0m"

	// Craftable
	EtoffeMilitaire ItemType = "Etoffe militaire"
	Kevlar          ItemType = "Kevlar"
	Cuir            ItemType = "Cuir"
	Camera          ItemType = "Caméra"

	// Equipment
	CasqueProtectionLv1    ItemType = "Casque de protection lv1"
	GiletPareBalleLv1      ItemType = "Gilet pare-balle lv1"
	BottesTactiquesLv1     ItemType = "Bottes tactiques lv1"
	CasqueProtectionLv2    ItemType = "Casque de protection lv2"
	GiletPareBalleLv2      ItemType = "Gilet pare-balle lv2"
	BottesTactiquesLv2     ItemType = "Bottes tactiques lv2"
	AugmentationInventaire ItemType = "Augmentation d'inventaire"
)

// max inventaire
const (
	MaxPistol  = 1
	MaxGrenade = 5
)

// Perso struct
type Perso struct {
	nom           string
	classe        string
	grade         int
	pvMAX         float64
	pv            float64
	inv           map[ItemType]int
	po            int
	skill         string
	Tete          string
	Torse         string
	Pieds         string
	Equipement    Equipement
	pointsAttaque int
}
type Monstre struct {
	nom           string
	pvMAX         float64
	pv            float64
	pointsAttaque int
}

// ItemType represents the type of an item.
type ItemType string

// Equipement Info représente des informations sur un élément d'équipement.
type EquipementInfo struct {
	Nom     string
	BonusPv int
}

// L'équipement représente les emplacements d'équipement
type Equipement struct {
	Tete  EquipementInfo
	Torse EquipementInfo
	Pieds EquipementInfo
}

// Initialisez la map des articles avec les types d'articles et leurs prix
var itemPrices = map[ItemType]int{
	PotionDeVie:            3,
	Grenade:                6,
	Couteau:                8,
	LivreDeSortRage:        25,
	PotionDePoison:         6,
	EtoffeMilitaire:        4,
	Kevlar:                 7,
	Cuir:                   3,
	Camera:                 1,
	AugmentationInventaire: 30,
}

// Créer une map pour les quantités maximales d'inventaire
var maxInventory = map[ItemType]int{
	PotionDeVie: 10,
	Grenade:     MaxGrenade,
}

// Initialiser un Perso
func CharCreation() *Perso {
	var nom string
	for {
		fmt.Print(redText + "Entrez votre nom : " + reset)
		fmt.Scan(&nom)
		if ValideNom(nom) {
			break
		} else {
			fmt.Println("⚠️Le nom doit contenir uniquement des lettres.")
		}
	}

	nom = FormatNom(nom)

	var classe string
	for {
		fmt.Print(redText + "Choisissez votre classe (LAT👮 ou TERRO👹) : " + reset)
		fmt.Scan(&classe)
		if EstValideClasse(classe) {
			break
		} else {
			fmt.Println("⚠️Classe invalide. Veuillez choisir LAT ou TERRO.")
		}
	}

	var grade int
	for {
		fmt.Print(redText + "Entrez votre grade : " + reset)
		_, err := fmt.Scanln(&grade)
		if err == nil {
			break
		} else {
			fmt.Println("⚠️Grade invalide. Veuillez entrer un nombre entier.")
		}
	}

	pvMAX := 100.0
	switch classe {
	case "LAT":
		pvMAX = 100.0
	case "TERRO":
		pvMAX = 80.0
	}
	pv := pvMAX / 2
	inv := make(map[ItemType]int)
	po := 100
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

// affiche le menu du personnage et permet au joueur de choisir les options.
func (p *Perso) Menu() {
	for true {
		fmt.Println(".-~-.-~-.-~.-~-.-~-.-~.-~-.-~-.-~.-~-.-~-.-~-")
		fmt.Println("-----------------MENU-----------------------:")
		fmt.Println("👨‍💻1. Afficher les informations du personnage")
		fmt.Println("🎒2. Accéder au contenu de l'inventaire")
		fmt.Println("💲3. Marchand")
		fmt.Println("🪓4. Forgeron")
		fmt.Println("5.Play")
		fmt.Println("6. Qui sont-ils ?")
		fmt.Println("◀️7. Quitter")

		scanner := bufio.NewScanner(os.Stdin)
		var choix string
		fmt.Print("Choisissez une option : ")
		scanner.Scan()
		choix = scanner.Text()
		fmt.Println("valeur choix : " + choix)

		switch choix {
		case "1":
			p.DisplayInfo()
		case "2":
			p.DisplayInventory()
		case "3":
			fmt.Println("Bienvenue chez le marchand !")
			fmt.Println(".-~-.-~-.-~.-~-.-~-.-~.-~-.-~-.-~.-~-.-~-.-~-")
			p.DisplayMarchand()
		case "4":
			fmt.Println("Bienvenue chez le forgeron !")
			fmt.Println(".-~-.-~-.-~.-~-.-~-.-~.-~-.-~-.-~.-~-.-~-.-~-")
			p.DisplayForgeron()
		case "5":
			playmenu(p)
		case "6":
			fmt.Println("ABBA")	
			fmt.Println("Steven Spielberg")	
			fmt.Println("QUEEN (Freddie Mercury)")	
		case "7":
			fmt.Println("Au revoir !")
			fmt.Println(".-~-.-~-.-~.-~-.-~-.-~.-~-.-~-.-~.-~-.-~-.-~-")
			return
		default:
			fmt.Println("Option invalide.")
		}
	}
}

// DisplayInventory affiche l'inventaire du personnage et permet au joueur d'interagir avec lui.
func (p *Perso) DisplayInventory() {
	fmt.Println("Inventaire du Personnage :")
	i := 1
	for item, quantity := range p.inv {
		fmt.Printf("%d. %s (Quantité : %d)\n", i, item, quantity)
		i++
	}

	fmt.Println("0. Retour au menu précédent")
	var choix string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Choisissez une option : ")
	scanner.Scan()
	choix = scanner.Text()
	if choix == "0" {
		p.Menu()
	}
}

// DisplayInfo affiche des informations sur le personnage.
func (p *Perso) DisplayInfo() {
	if p.classe == "TERRO" {
		fmt.Println("☠-----▄︻デ══━一💥----👳🏽‍♂️-----TERORISTE-----▄︻デ══━一💥-----👳🏽‍♂️------☠")
	} else {
		fmt.Println("------👮🚨🚓🚨👮--------LUTE ANTI TERRORISTE--------👮🚨🚓🚨👮--------")
	}

	fmt.Println(".-~-.-~-.-~.-~-.-~-.-~.-~-.-~-.-~.-~-.-~-.-~-")
	fmt.Println("Informations sur le personnage:")
	fmt.Println("Nom :", p.nom)
	fmt.Println("Classe :", p.classe)
	fmt.Println("Grade :", p.grade)
	fmt.Println("Points de vie actuels :", p.pv)
	fmt.Println("Points de vie maximum :", p.pvMAX)
	fmt.Println("Inventaire :", p.inv)
	fmt.Println("Pièces d'or :", p.po)
	fmt.Println("Compétence :", p.skill)
	fmt.Println("Tête :", p.Tete)
	fmt.Println("Torse :", p.Torse)
	fmt.Println("Pieds :", p.Pieds)
}

// DisplayMarchand affiche les articles disponibles à l'achat chez le commerçant.
func (p *Perso) DisplayMarchand() {
	fmt.Println("Bienvenue chez le marchand !")
	fmt.Println("Voici la liste des marchandises :")
	fmt.Println("1. Potion de vie - 3 pièces d'or")
	fmt.Println("2. Grenade - 6 pièces d'or")
	fmt.Println("3. Couteau - 8 pièces d'or")
	fmt.Println("4. Livre de sort : Rage - 25 pièces d'or")
	fmt.Println("5. Potion de poison - 6 pièces d'or")
	fmt.Println("6. Etoffe militaire - 4 pièces d'or")
	fmt.Println("7. Kevlar - 7 pièces d'or")
	fmt.Println("8. Cuir - 3 pièces d'or")
	fmt.Println("9. Caméra - 1 pièce d'or")
	fmt.Println("10. Augmentation d'inventaire - 30 pièces d'or")
	fmt.Println("0. Quitter")

	var choix int
	fmt.Print("Choisissez un item à acheter : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		if p.Checkinv() {
			p.BuyItem(PotionDeVie, 3)
		}
	case 2:
		if p.Checkinv() {
			p.BuyItem(Grenade, 6)
		}
	case 3:
		if p.Checkinv() {
			p.BuyItem(Couteau, 8)
		}
	case 4:
		if p.Checkinv() {
			p.BuyItem(LivreDeSortRage, 25)
		}
	case 5:
		if p.Checkinv() {
			p.BuyItem(PotionDePoison, 6)
		}
	case 6:
		if p.Checkinv() {
			p.BuyItem(EtoffeMilitaire, 4)
		}
	case 7:
		if p.Checkinv() {
			p.BuyItem(Kevlar, 7)
		}
	case 8:
		if p.Checkinv() {
			p.BuyItem(Cuir, 3)
		}
	case 9:
		if p.Checkinv() {
			p.BuyItem(Camera, 1)
		}
	case 10:
		if p.Checkinv() {
			fmt.Println("Vous ne pouvez pas acheter une augmentation d'inventaire directement.")
		}
	default:
		fmt.Println("Choix invalide.")
	}
}

// DisplayForgeron affiche les objets disponibles pour la fabrication chez le forgeron.
func (p *Perso) DisplayForgeron() {
	fmt.Println(".-~-.-~-.-~.-~-.-~-.-~")
	fmt.Println("Bienvenue chez le forgeron !")
	fmt.Println("Voici la liste des fabrications :")
	fmt.Println("1. Casque de protection lv1 - 5 pièces d'or, 1 Caméra, 1 Cuir")
	fmt.Println("2. Gilet pare-balle lv1 - 5 pièces d'or, 2 Etoffe militaire, 1 Kevlar")
	fmt.Println("3. Bottes tactiques lv1 - 5 pièces d'or, 1 Etoffe militaire, 1 Cuir")
	fmt.Println("4. Casque de protection lv2 - 10 pièces d'or, 2 Caméras, 2 Cuir")
	fmt.Println("5. Gilet pare-balle lv2 - 10 pièces d'or, 4 Etoffe militaire, 2 Kevlar")
	fmt.Println("6. Bottes tactiques lv2 - 10 pièces d'or, 2 Etoffe militaire, 2 Cuir")
	fmt.Println("0. Quitter")

	var choix int
	fmt.Print("Choisissez un équipement à fabriquer : ")
	fmt.Scan(&choix)

	switch choix {
	case 1:
		p.CraftItem(CasqueProtectionLv1, 5, Camera, 1, Cuir, 1)
	case 2:
		p.CraftItem(GiletPareBalleLv1, 5, EtoffeMilitaire, 2, Kevlar, 1)
	case 3:
		p.CraftItem(BottesTactiquesLv1, 5, EtoffeMilitaire, 1, Cuir, 1)
	case 4:
		p.CraftItem(CasqueProtectionLv2, 10, Camera, 2, Cuir, 2)
	case 5:
		p.CraftItem(GiletPareBalleLv2, 10, EtoffeMilitaire, 4, Kevlar, 2)
	case 6:
		p.CraftItem(BottesTactiquesLv2, 10, EtoffeMilitaire, 2, Cuir, 2)
	case 0:
		fmt.Println("Au revoir !")
		return
	default:
		fmt.Println("Choix invalide.")
	}
}
func InitGoblin() Monstre {
	return Monstre{
		nom:           "Gobelin d'entraînement",
		pvMAX:         40,
		pv:            40,
		pointsAttaque: 5,
	}
}
