package perso

import (
	"fmt"

	"github.com/mbndr/figlet4go"
)

func DisplayCSGOWelcomefight() {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("Game started")
	greenText := "\x1b[32m" + renderStr + "\x1b[0m"
	fmt.Println(greenText)
}
