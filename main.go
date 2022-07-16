// First Go program
package main
  
import (
	"fmt"
	"flag"
)
// Main function
func main() {
  
	// variables declaration  
    var player_name string  
    var character_name string

	// flags declaration using flag package
	flag.StringVar(&player_name, "pname", "test1", "Specify player's name. Required.")
	flag.StringVar(&character_name, "cname", "test2", "Specify character's name. Required.")

	flag.Parse()  // after declaring flags we need to call it

    fmt.Println("Golang Character Generator!")
	fmt.Println("Player Name: " + player_name)
	fmt.Println("Character Name: " + character_name)

}