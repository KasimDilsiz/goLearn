package main
// cases can be presented in comma-separated lists.//
import "fmt"

func main() {
	n := "Bond"

	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("miss money or bond or do no")

	case "m":
		fmt.Println("james bond")
	case "q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}
}
