package perso

import (
	"fmt"
	"time"
)

func playmenu(p *Perso) {
	ClearScreen()
	DisplayCSGOWelcomefight()
	var choix string
	fmt.Println(redText + "🥋1/trainingFight" + reset)
	fmt.Println(redText + "🥋2/" + reset)
	fmt.Scan(&choix)
	switch choix {
	case "1":
		trainingFight(p)
	}

}
func trainingFight(player *Perso) {
	ClearScreen()
	monstre := InitGoblin()
	tourDeCombat := 1
	for player.pv > 0 && monstre.pv > 0 {
		fmt.Printf(redText+"\n-- Tour de combat %d --\n"+reset, tourDeCombat)
		fmt.Printf(redText+"Points de vie du joueur : %.1f / %.1f\n"+reset, player.pv, player.pvMAX)
		fmt.Printf(redText+"Points de vie du monstre : %.1f / %.1f\n"+reset, monstre.pv, monstre.pvMAX)
		fmt.Println(redText + "C'est votre tour. Que voulez-vous faire ?" + reset)
		fmt.Println(redText + "1. Attaquer" + reset)
		fmt.Println(redText + "2. Quitter le combat" + reset)
		var choix int
		fmt.Print(redText + "Choisissez une option : " + reset)
		fmt.Scan(&choix)
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
			fmt.Println(yellowText + "Vous avez vaincu le monstre !" + resety)
			return
		}
		fmt.Println("\nC'est le tour du monstre.")
		player.pv -= 5
		fmt.Printf("Le monstre vous inflige 5 points de dégâts.\n")
		if player.pv <= 0 {
			fmt.Println("Vous avez été vaincu par le monstre.")
			return
		}
		tourDeCombat++
		time.Sleep(1 * time.Second)
	}
}
func goblinPattern(player *Perso) {
    goblin := InitGoblin()

    for tour := 1; player.pv > 0; tour++ {
        var DegatsInfliges int

        if tour%3 == 0 {
            DegatsInfliges = goblin.pointsAttaque * 2 
        } else {
            DegatsInfliges = goblin.pointsAttaque 

        player.pv -= float64(DegatsInfliges)
        fmt.Printf("%s inflige à %s %d de dégâts.\n", goblin.nom, player.nom, DegatsInfliges)
        fmt.Printf("Points de vie actuels de %s : %.1f/%.1f\n", player.nom, player.pv, player.pvMAX)

        if player.pv <= 0 {
            fmt.Printf("%s a été vaincu par %s !\n", player.nom, goblin.nom)
            break
        }	
    }
}
}