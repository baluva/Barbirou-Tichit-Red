package main

import (
	"fmt"
	"time"
)

func (p *Perso) takeapot() {
	var potionIdx = -1
	for i, item := range p.inv {
		if item == "Potion de vie" {
			potionIdx = i
			break
		}
	}
	p.pv += 50
	///verif pv
	if p.pv > p.pvMAX {
		p.pv = p.pvMAX
	}
	p.inv = append(p.inv[:potionIdx], p.inv[potionIdx+1:]...)

	fmt.Println("Vous avez utilisé une Potion de vie.")
	fmt.Printf("Nouveaux points de vie actuels : %d / %d\n", p.pv, p.pvMAX)
}
func (p *Perso) acheterItem(item string) {
	p.inv = append(p.inv, item)
	fmt.Printf("%s a acheté : %s\n", p.nom, item)
}
func Dead(p *Perso) {
	if p.pv <= 0 {
		fmt.Printf("%s est mort !\n", p.nom)
		// Resurrection avec 50% de points de vie maximum
		p.pv = p.pvMAX * 0.5
		fmt.Printf("%s est ressuscité avec %.2f points de vie\n", p.nom, p.pv)
	}
}
func poisonPot(p *Perso) {
	fmt.Printf("une Potion de poison \n")
	// Infliger des dégâts pendant 3 secondes
	for i := 0; i < 3; i++ {
		p.pv -= 10
		fmt.Printf("Points de vie actuels : %.2f / %.2f\n", p.pv, p.pvMAX)
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("L'effet du poison a disparu.\n")
}

func removeInventory(p *Perso, item string) {
	for i, invItem := range p.inv {
		if invItem == item {
			p.inv = append(p.inv[:i], p.inv[i+1:]...)
			fmt.Printf("%s a vendu : %s\n", p.nom, item)
			return
		}
	}
}

func (p *Perso) Addinv(item string) {
	p.inv = append(p.inv, item)
	fmt.Printf("%s a acheté : %s\n", p.nom, item)
	fmt.Println(p.inv)
}
