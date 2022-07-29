// First Go program
package main

import (
	// for printing out things
	"fmt"
	"time"

	// for exiting the program + command line args
	"os"
	// this allows users to flag their command line args (like -pname Lucca)
	"flag"
	// allows the logging of our errors (currently just for the api call)
	"log"
	// used to make the http.get request to the api
	"net/http"

	//used for random rolling
	"math/rand"
	//allows us to parse the json the api sends over
	"encoding/json"
	//allows us to take the httpresponse and make the body (or information) into a byte array
	"bytes"
	"io/ioutil"
)

//doesn't round down properly on the lower ones? shouldn't matter when stat rolling is un-simplified
func get_mod(x int) int {
	return ((x - 10) / 2)
}

// Main function
func main() {

	// variables declaration
	var player_name string
	var character_name string
	var character_race string

	//not 4d6 drop lowest, but trying to be simple first
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var STR int = (r1.Intn((20 - 3 + 1) + 3))
	var DEX int = (r1.Intn((20 - 3 + 1) + 3))
	var CON int = (r1.Intn((20 - 3 + 1) + 3))
	var INT int = (r1.Intn((20 - 3 + 1) + 3))
	var WIS int = (r1.Intn((20 - 3 + 1) + 3))
	var CHA int = (r1.Intn((20 - 3 + 1) + 3))

	var modSTR int = (get_mod(STR))
	var modDEX int = (get_mod(DEX))
	var modCON int = (get_mod(CON))
	var modINT int = (get_mod(INT))
	var modWIS int = (get_mod(WIS))
	var modCHA int = (get_mod(CHA))

	var str_line = fmt.Sprintf("Strength: %d, Mod: %d", STR, modSTR)
	var dex_line = fmt.Sprintf("Dexterity: %d, Mod: %d", DEX, modDEX)
	var con_line = fmt.Sprintf("Constitution: %d, Mod: %d", CON, modCON)
	var int_line = fmt.Sprintf("Intelligence: %d, Mod: %d", INT, modINT)
	var wis_line = fmt.Sprintf("Wisdom: %d, Mod: %d", WIS, modWIS)
	var cha_line = fmt.Sprintf("Charisma: %d, Mod: %d", CHA, modCHA)

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
	fmt.Println(str_line)
	fmt.Println(dex_line)
	fmt.Println(con_line)
	fmt.Println(int_line)
	fmt.Println(wis_line)
	fmt.Println(cha_line)

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

	//this grabs the body of information, and attempts to read it.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	//this converts the byte array into a string and prints it. This contains all the race information.
	toprint := bytes.NewBuffer(body).String()
	fmt.Println(toprint)
}

// RACE API CALL CONSTRUCTORs
type Race struct {
}

type RaceAPIResponse struct {
	ExecutionTime string `json:"executionTime"`
	RaceInfoList  []Race `json:"RaceInfoList"`
}

// this parses the body of the httpresponse, and the error if it gets one.
func getRaceInfo(body []byte) (*RaceAPIResponse, error) {
	var s = new(RaceAPIResponse)
	//takes the json to a readable format
	//TODO look into json unmarshal command to see what it does.
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

//API DOCS
//Api to bytearray https://blog.josephmisiti.com/parsing-json-responses-in-golang
//byte array to json? https://stackoverflow.com/questions/45756011/parse-json-http-response-using-golang
//byte array to string https://golangdocs.com/golang-byte-array-to-string
