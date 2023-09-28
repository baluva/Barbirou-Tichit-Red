package perso

import (
	"fmt"
	"time"
)

func playmenu(p *Perso) {
	monstre := InitGoblin()
	DisplayCSGOWelcomefight()
	var choix string
	fmt.Println(redText + "🥋1/train" + reset)
	fmt.Println(redText + "🥋2/ fight a monster" + reset)
	fmt.Println(redText + "🧙3/ combat magique" + reset)
	fmt.Scanln(&choix)
	switch choix {
	case "1":
		trainingFight(p)
	case "2":
		combat(p, &monstre)
	case "3":
		combatmagique(p, &monstre)
	}

}
func trainingFight(player *Perso) {
	ClearScreen()
	monstre := InitGoblin()
	tourDeCombat := 1
	for player.pv > 0 && monstre.pv > 0 {
		ClearScreen()
		fmt.Println("visez bien!!!")
		fmt.Printf(redText+"\n-- Tour de combat %d --\n"+reset, tourDeCombat)
		fmt.Printf(redText+"Points de vie du joueur : %.1f / %.1f\n"+reset, player.pv, player.pvMAX)
		fmt.Printf(redText+"Points de vie du monstre : %.1f / %.1f\n"+reset, monstre.pv, monstre.pvMAX)
		fmt.Println(redText + "C'est votre tour. Que voulez-vous faire ?" + reset)
		fmt.Println(redText + "1. Attaquer" + reset)
		fmt.Println(redText + "2. Quitter le combat" + reset)
		var choix int
		fmt.Print(redText + "Choisissez une option : " + reset)
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			monstre.pv -= 5
			fmt.Printf("Vous infligez 5 points de dégâts au monstre.\n")
		case 2:
			fmt.Println("Vous avez quitté le combat.")
			return
		default:
			fmt.Println("Option invalide. Choisissez une option valide.")
			continue
		}
		if monstre.pv <= 0 {
			ClearScreen()
			fmt.Println("😵😵😵😵😵😵😵😵😵")
			fmt.Println(yellowText + "Vous avez vaincu le monstre !" + resety)
			playerAugmentStats(player)
			return
		}
		fmt.Println("\nC'est le tour du monstre.")
		player.pv -= 5
		fmt.Printf("Le monstre vous inflige 5 points de dégâts.\n")
		if player.pv <= 0 {
			fmt.Println("Vous avez été vaincu par le monstre.")
			player.dead() 
			return
		}
		tourDeCombat++
		time.Sleep(1 * time.Second)
	}
}
func monsterAttack(player *Perso, monster *Monstre) {
	damage := monster.pointsAttaque
	fmt.Printf("%s inflige %d dégâts à %s\n", monster.nom, damage, player.nom)
	player.pv -= float64(damage)
	fmt.Printf("%s PV : %.2f / %.2f\n", player.nom, player.pv, player.pvMAX)
}
func charTurn(player *Perso, monster *Monstre) {
	fmt.Printf("Tour du joueur (%s)\n", player.nom)
	fmt.Println("Menu:")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")

	var choice int
	fmt.Print("Choisissez une option : ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		damage := 5
		fmt.Printf("%s inflige %d dégâts à %s\n", player.nom, damage, monster.nom)
		monster.pv -= float64(damage)
		fmt.Printf("%s PV : %.2f / %.2f\n", monster.nom, monster.pv, monster.pvMAX)
	case 2:
		fmt.Println("Inventaire :")
		for item, quantity := range player.inv {
			fmt.Printf("%s x%d\n", item, quantity)
		}
		var inventoryChoice ItemType
		fmt.Print("Choisissez un objet de l'inventaire (ou '0' pour revenir) : ")
		fmt.Scanln(&inventoryChoice)

		if inventoryChoice != "0" {
			if quantity, ok := player.inv[inventoryChoice]; ok && quantity > 0 {
				fmt.Printf("Vous utilisez %s\n", inventoryChoice)
				player.inv[inventoryChoice]--
			} else {
				fmt.Println("Objet non disponible.")
			}
		}
	}
	fmt.Printf("Tour du monstre (%s)\n", monster.nom)
	monsterAttack(player, monster)
	if player.pv <= 0 {
		fmt.Println("Vous avez perdu la partie.")
		player.dead() 
	} else if monster.pv <= 0 {
		fmt.Println("Vous avez vaincu le monstre !")
	}
}
func combat(player *Perso, monster *Monstre) {
	fmt.Println("Vous entrez dans un combat d'entraînement contre un Gobelin d'entraînement !")

	for player.pv > 0 && monster.pv > 0 {
		fmt.Println("\nNouveau tour de jeu d'entraînement :")
		charTurn(player, monster)
		if player.pv <= 0 {
			fmt.Println("Vous avez perdu la partie.")
			break
		} else if monster.pv <= 0 {
			fmt.Println("Vous avez vaincu le monstre !")
			playerAugmentStats(player)
			break
		}
		var continueChoice string
		fmt.Print("Voulez-vous continuer l'entraînement ? (O/N) : ")
		fmt.Scanln(&continueChoice)
		if continueChoice != "O" && continueChoice != "o" {
			fmt.Println("Vous quittez l'entraînement.")
			break
		}
	}
}
func playerAugmentStats(player *Perso) {
	switch player.Niveau {
	case 2:
		player.pvMAX += 10
	case 3:
		player.pvMAX += 15
	case 4:
		player.pvMAX += 20
	case 5:
		player.pvMAX += 30
	case 6:
		fmt.Println("niveau max")
	}
}
func combatmagique(player *Perso, monster *Monstre) {

}
