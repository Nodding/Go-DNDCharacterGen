// First Go program
package main
  
import (
	"fmt"
	"os"
	"flag"
)
// Main function
func main() {
  
	// variables declaration  
    var player_name string  
    var character_name string

	// flags declaration using flag. The default values will be noname and we need to check if it is.
	flag.StringVar(&player_name, "pname", "noname", "Specify player's name. Required.")
	flag.StringVar(&character_name, "cname", "noname", "Specify character's name. Required.")
	flag.Parse()  // after declaring flags we need to call it


	if player_name == "noname" {
		fmt.Println("No player name!")
		os.Exit(0)
	}
	if character_name == "noname" {
		fmt.Println("No character name!")
		os.Exit(0)
	}

    fmt.Println("Golang Character Generator!")
	fmt.Println("Player Name: " + player_name)
	fmt.Println("Character Name: " + character_name)

}