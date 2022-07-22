// First Go program
package main

import (
	// for printing out things
	"fmt"
	// for exiting the program + command line args
	"os"
	// this allows users to flag their command line args (like -pname Lucca)
	"flag"
	// allows the logging of our errors (currently just for the api call)
	"log"
	// used to make the http.get request to the api
	"net/http"
)

// Main function
func main() {

	// variables declaration
	var player_name string
	var character_name string
	var character_race string

	// flags declaration using flag. The default values will be none (needs to be a string to match the type) and we need to check if it is.
	flag.StringVar(&player_name, "pname", "none", "Specify player's name. Required.")
	flag.StringVar(&character_name, "cname", "none", "Specify character's name. Required.")
	flag.StringVar(&character_race, "crace", "none", "Specify character's race. Chosen at random if nothing is provided.")
	flag.Parse() // after declaring flags we need to call it

	if player_name == "none" {
		fmt.Println("No player name!")
		os.Exit(0)
	}
	if character_name == "none" {
		fmt.Println("No character name!")
		os.Exit(0)
	}
	if character_race == "none" {

		// for now, character race is defaulted to dwarf. need to instead choose one from random list later.
		character_race = "dwarf"
	}

	//printing out information give to us!
	fmt.Println("Golang Character Generator!")
	fmt.Println("Player Name: " + player_name)
	fmt.Println("Character Name: " + character_name)
	fmt.Println("Character Race: " + character_race)

	//calls the api function
	api("https://www.dnd5eapi.co/api/races/" + character_race)
}

//our call an api function.
//
// right now all it does is see if it can get information, does nothing with it yet.
func api(url string) {

	//assign the response variable, and if there is an error, the error variable is filled with information from the site
	//http.Get(url) does the actual call to the server
	response, err := http.Get(url)

	//if there is information (like an error) we log it.
	if err != nil {
		log.Fatal(err)
	}

	//since there was no error, we can go ahead and print out the information.
	//in the future, we need to actually parse what they give us to extract the information out.
	fmt.Println(response)
}
