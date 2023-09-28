package perso

import (
	"fmt"
	"time"

	"github.com/mbndr/figlet4go"
)

func DisplayCSGOWelcomefight() {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("Game started")
	greenText := "\x1b[32m" + renderStr + "\x1b[0m"
	fmt.Println(greenText)
}
func goodbye() {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("GOOD BYE!")
	greenText := "\x1b[32m" + renderStr + "\x1b[0m"
	fmt.Println(greenText)
	time.Sleep(5 * time.Second)
}
func displaymagique() {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("LAT VS TERRO")
	greenText := "\x1b[32m" + renderStr + "\x1b[0m"
	fmt.Println(greenText)
	time.Sleep(5 * time.Second)
}
