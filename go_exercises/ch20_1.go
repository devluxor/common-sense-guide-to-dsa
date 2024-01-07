package main

import (
	"fmt"

)

// You are to write a function that accepts two arrays of players and returns
// an array of the players who play in both sports. In this case, that would be:

func main() {
	basketballPlayers := []map[string]string{
		{"first_name": "Jill", "last_name": "Huang", "team": "Gators"},
		{"first_name": "Janko", "last_name": "Barton", "team": "Sharks"},
		{"first_name": "Wanda", "last_name": "Vakulskas", "team": "Sharks"},
		{"first_name": "Jill", "last_name": "Moloney", "team": "Gators"},
		{"first_name": "Luuk", "last_name": "Watkins", "team": "Gators"}}

	footballPlayers := []map[string]string{
		{"first_name": "Hanzla", "last_name": "Radosti", "team": "32ers"},
		{"first_name": "Tina", "last_name": "Watkins", "team": "Barleycorns"},
		{"first_name": "Alex", "last_name": "Patel", "team": "32ers"},
		{"first_name": "Jill", "last_name": "Huang", "team": "Barleycorns"},
		{"first_name": "Wanda", "last_name": "Vakulskas", "team": "Barleycorns"}}

	fmt.Println(commonPlayers(&basketballPlayers, &footballPlayers))
}

func commonPlayers(basketPlayers, footballPlayers *[]map[string]string) []string {
	basketNameList := createNameList(basketPlayers)

	output := []string{}
	playerName := ""
	for _, p := range *footballPlayers {
		playerName = parseName(p)
		if basketNameList[playerName] {
			output = append(output, playerName)
		}
	}

	return output
}

func createNameList(playerList *[]map[string]string) map[string]bool {
	list := make(map[string]bool)

	for _, p := range *playerList {
		list[parseName(p)] = true
	}

	return list
}

func parseName(player map[string]string) string {
	return fmt.Sprintf("%s %s", player["first_name"], player["last_name"])
}

/*

for every player in basket:

	create a string, first + last name

	store stringName into new hash, basketPlayersNameList

for every player in football players

	create a string, first + last name

	if the player exists in basketPlayersNameList,

		add to output array

return o. array

*/