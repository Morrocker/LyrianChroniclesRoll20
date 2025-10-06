package main

import (
	"encoding/json"
	"maps"
	"os"
)

type TrueAbility struct {
	ID          string
	Name        string
	Keywords    string
	Range       string
	Requirement string
	Description string
	RPcost      string
	APcost      string
	MPcost      string
	Othercost   string
}

func main() {
	abilities := make(map[string]TrueAbility, 0)
	abilities2 := make(map[string]TrueAbility, 0)

	abytes, err := os.ReadFile("true_abilities.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(abytes, &abilities)
	if err != nil {
		panic(err)
	}

	abytes2, err := os.ReadFile("races_true_abilities.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(abytes2, &abilities2)
	if err != nil {
		panic(err)
	}

	// merge the two maps
	maps.Copy(abilities, abilities2)

	// export the data to a csv file
	file, err := os.Create("true_abilities.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// write the header
	_, err = file.WriteString("ID,Macro\n")
	if err != nil {
		panic(err)
	}
	for _, ability := range abilities {
		// print the ability ID and an empty column for the macro
		line := ability.ID + ",\n"
		_, err = file.WriteString(line)
		if err != nil {
			panic(err)
		}
	}
}
