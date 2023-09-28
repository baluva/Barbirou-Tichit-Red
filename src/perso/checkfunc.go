package perso

import (
	"fmt"
	"strings"
)

// BuyItem permet au personnage d'acheter un objet auprès du marchand.
func (p *Perso) BuyItem(item ItemType, price int) {
	if p.po >= price {
		p.AddToInventory(item, 1)
		p.po -= price
		fmt.Printf("Vous avez acheté %s pour %d pièces d'or.\n", item, price)
	} else {
		fmt.Println("Vous n'avez pas assez de pièces d'or pour acheter cet item.")
	}
}

// AddToInventory ajoute un objet à l'inventaire du personnage.
func (p *Perso) AddToInventory(item ItemType, quantity int) {
	currentQuantity, exists := p.inv[item]
	if exists {
		p.inv[item] = currentQuantity + quantity
	} else {
		p.inv[item] = quantity
	}
}

// CraftItem permet au personnage de fabriquer un objet chez le forgeron.
func (p *Perso) CraftItem(result ItemType, cost int, ingredients ...interface{}) {
	for i := 0; i < len(ingredients); i += 2 {
		ingredient, ok := ingredients[i].(ItemType)
		if !ok {
			fmt.Println("Erreur de données de l'ingrédient.")
			return
		}

		quantity, ok := ingredients[i+1].(int)
		if !ok {
			fmt.Println("Erreur de données de la quantité d'ingrédient.")
			return
		}

		currentQuantity, exists := p.inv[ingredient]
		if !exists || currentQuantity < quantity {
			fmt.Printf("Vous n'avez pas assez de %s pour fabriquer cet équipement.\n", ingredient)
			return
		}
	}

	// Vérifiez si le personnage a suffisamment d'or
	if p.po < cost {
		fmt.Println("Vous n'avez pas assez de pièces d'or pour fabriquer cet équipement.")
		return
	}

	// Déduire le coût et les ingrédients du personnage
	p.po -= cost
	for i := 0; i < len(ingredients); i += 2 {
		ingredient := ingredients[i].(ItemType)
		quantity := ingredients[i+1].(int)
		p.inv[ingredient] -= quantity
		if p.inv[ingredient] == 0 {
			delete(p.inv, ingredient)
		}
	}

	// Ajoutez l'objet fabriqué à l'inventaire du personnage
	p.AddToInventory(result, 1)
	fmt.Printf("Vous avez fabriqué %s avec succès !\n", result)
}

// FormatNom formate le nom du personnage.
func FormatNom(nom string) string {
	return strings.Title(strings.ToLower(nom))
}

// ValideNom vérifie si le nom du personnage ne contient que des lettres.
func ValideNom(nom string) bool {
	return strings.TrimSpace(nom) != "" && isAlpha(nom)
}

func isAlpha(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return false
		}
	}
	return true
}

// EstValideClasse vérifie si la classe du personnage est valide.
func EstValideClasse(classe string) bool {
	return classe == "LAT" || classe == "TERRO"
}
func ClearScreen() {
	// Clear the console screen
	fmt.Print("\033[H\033[2J")
}
func (p *Perso) Checkinv() bool {
	count := 0
	for _, quantity := range p.inv {
		count += quantity
		if count >= 10 {
			return false
		}
	}
	return true

}
func (p *Perso) dead() {
	if p.pv <= 0 {
		fmt.Println("Vous êtes mort !")
		p.pv = p.pvMAX * 0.5
		fmt.Printf("Vous avez été ressuscité avec %.1f points de vie.\n", p.pv)
	}
}
func (p *Perso) takePot(item ItemType) {
	if quantity, exists := p.inv[item]; exists && quantity > 0 {
		fmt.Printf("Vous utilisez une %s.\n", item)
		p.inv[item]--
		hpToAdd := 50.0
		if p.pv+hpToAdd > p.pvMAX {
			hpToAdd = p.pvMAX - p.pv
		}
		p.pv += hpToAdd
		fmt.Printf("Points de vie actuels : %.1f / %.1f\n", p.pv, p.pvMAX)
	} else {
		fmt.Printf("Vous n'avez pas de %s dans votre inventaire.\n", item)
	}
}
