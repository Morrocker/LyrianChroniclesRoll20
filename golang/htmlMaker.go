package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type ClassInfo struct {
	Name string
	ID   string
	Tier int
}

type KeyAbility struct {
	ID       string
	Name     string
	Benefits []string
}

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
}

type CraftingAbility struct {
	ID          string
	Name        string
	Keywords    string
	Cost        string
	Description string
}

type Breakthrough struct {
	ID           string
	Name         string
	Cost         string
	Requirement  string
	Requirements string
	Description  string
}

func main() {
	breakthroughsMap := make(map[string]Breakthrough)
	trueAbilities := make(map[string]TrueAbility)
	craftingAbilities := make(map[string]CraftingAbility)
	keyAbilities := make(map[string]KeyAbility)
	classesMap := make(map[string]ClassInfo)

	// Loading all JSONS

	bytes, err := os.ReadFile("./webscrapper/breakthroughs.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(bytes, &breakthroughsMap); err != nil {
		panic(err)
	}

	bytes, err = os.ReadFile("./webscrapper/true_abilities.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(bytes, &trueAbilities); err != nil {
		panic(err)
	}

	bytes, err = os.ReadFile("./webscrapper/crafting_abilities.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(bytes, &craftingAbilities); err != nil {
		panic(err)
	}

	bytes, err = os.ReadFile("./webscrapper/key_abilities.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(bytes, &keyAbilities); err != nil {
		panic(err)
	}

	bytes, err = os.ReadFile("./webscrapper/classes.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(bytes, &classesMap); err != nil {
		panic(err)
	}
	var classes []ClassInfo

	for _, class := range classesMap {

		c := ClassInfo{
			ID:   class.ID,
			Name: class.Name,
			Tier: class.Tier,
		}
		classes = append(classes, c)
		fmt.Printf("id: %s, name: %s, tier: %d\n", c.ID, c.Name, c.Tier)
	}

	sort.Slice(classes, func(i, j int) bool {
		return classes[i].ID < classes[j].ID
	})

	// Create the HTML file
	htmlFile, err := os.Create("update-data.html")
	if err != nil {
		panic(err)
	}
	defer htmlFile.Close()

	fmt.Fprintln(htmlFile, "<!-- All Clases Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)
	fmt.Fprintln(htmlFile, `<option value="custom">Custom</option>`)
	for _, class := range classesMap {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, class.ID, class.Name)
	}
	fmt.Fprintln(htmlFile, "<!-- All Clases Picker List END -->")

	// ============================= BREAKTHROUGHS =============================

	var breakthroughs []Breakthrough

	for _, bt := range breakthroughsMap {

		if bt.ID[0] >= '0' && bt.ID[0] <= '9' {
			bt.ID = "x" + bt.ID
		}

		b := Breakthrough{
			ID:           bt.ID,
			Name:         bt.Name,
			Cost:         bt.Cost,
			Requirements: strings.ReplaceAll(bt.Requirement, `"`, "'"),
			Description:  strings.ReplaceAll(bt.Description, `"`, "'"),
		}
		breakthroughs = append(breakthroughs, b)
		fmt.Printf("id: %s, name: %s, cost: %s\n", b.ID, b.Name, b.Cost)
	}
	sort.Slice(breakthroughs, func(i, j int) bool {
		return breakthroughs[i].Name < breakthroughs[j].Name
	})

	fmt.Fprintln(htmlFile, "<!-- All Breakthroughs Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)
	fmt.Fprintln(htmlFile, `<option value="custom">Custom</option>`)

	for _, breakthrough := range breakthroughs {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, breakthrough.ID, breakthrough.Name)
	}

	fmt.Fprintln(htmlFile, "<!-- All Breakthroughs Picker List END -->")

	// ============================= BREAKTHROUGHS END =============================

	// ============================= ABILITIES START =============================

	type Ability struct {
		ID           string
		Name         string
		Type         string
		Keywords     string
		Range        string
		Description  string
		Requirements string
		Costs        string
		OtherCosts   string
		Benefits     string
		AtkTypes     string
	}

	var regularAbilitiesArray []Ability
	var craftingAbilitiesArray []Ability

	for _, ab := range craftingAbilities {
		ability := Ability{
			ID:          ab.ID,
			Name:        ab.Name,
			Type:        "crafting_ability",
			Description: strings.ReplaceAll(ab.Description, `"`, "'"),
			Costs:       strings.ReplaceAll(ab.Cost, `"`, "'"),
		}

		craftingAbilitiesArray = append(craftingAbilitiesArray, ability)
	}
	sort.Slice(craftingAbilitiesArray, func(i, j int) bool {
		return craftingAbilitiesArray[i].ID < craftingAbilitiesArray[j].ID
	})

	for _, ab := range trueAbilities {
		var costs string = ""
		if ab.APcost != "" {
			costs += ab.APcost
		}
		if ab.MPcost != "" {
			if costs != "" {
				costs += ", "
			}
			costs += ab.MPcost
		}
		if ab.RPcost != "" {
			if costs != "" {
				costs += ", "
			}
			costs += ab.RPcost
		}

		ability := Ability{
			ID:           ab.ID,
			Name:         ab.Name,
			Type:         "true_ability",
			Keywords:     ab.Keywords,
			Range:        ab.Range,
			Description:  strings.TrimSpace(strings.ReplaceAll(ab.Description, `"`, "'")),
			Requirements: strings.ReplaceAll(ab.Requirement, `"`, "'"),
			Costs:        strings.ReplaceAll(costs, `"`, "'"),
		}
		regularAbilitiesArray = append(regularAbilitiesArray, ability)
	}

	for _, ab := range keyAbilities {
		benefits := strings.Join(ab.Benefits, " ")
		benefits2 := fmt.Sprintf(`%s`, benefits)

		ability := Ability{
			ID:       ab.ID,
			Name:     ab.Name,
			Type:     "key_ability",
			Benefits: strings.TrimSpace(strings.ReplaceAll(benefits2, `"`, "'")),
		}
		regularAbilitiesArray = append(regularAbilitiesArray, ability)
	}

	sort.Slice(regularAbilitiesArray, func(i, j int) bool {
		return regularAbilitiesArray[i].ID < regularAbilitiesArray[j].ID
	})

	fmt.Fprintln(htmlFile, "<!-- All Abilities Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)
	fmt.Fprintln(htmlFile, `<option value="custom">Custom</option>`)

	for _, ability := range regularAbilitiesArray {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, ability.ID, ability.Name)
	}

	fmt.Fprintln(htmlFile, "<!-- All abilities Picker List END -->")

	// Crafting abilities
	fmt.Fprintln(htmlFile, "<!-- All Crafting Abilities Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)
	fmt.Fprintln(htmlFile, `<option value="custom">Custom</option>`)

	for _, ability := range craftingAbilitiesArray {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, ability.ID, ability.Name)
	}

	fmt.Fprintln(htmlFile, "<!-- All Crafting abilities Picker List END -->")

	// ============================= ABILITIES END =============================

	fmt.Fprintln(htmlFile, "<!-- SCRIPTS TO REPLACE START -->")

	// Classes array
	var arraystring string
	arraystring = "{"
	for i, class := range classes {
		if i > 0 {
			arraystring += ","
		}
		arraystring += fmt.Sprintf(`"%s": {name: "%s",tier:%d}`, class.ID, class.Name, class.Tier)
	}
	arraystring += "}"

	fmt.Fprintf(htmlFile, `const classList =%s;`, arraystring)
	fmt.Fprintln(htmlFile, "\n")

	// Breakthroughs array
	var arraystring2 string
	arraystring2 = "{"
	for i, breakthrough := range breakthroughs {
		if i > 0 {
			arraystring2 += ","
		}
		description := strings.Replace(breakthrough.Description, "\n", " ", -1)
		arraystring2 += fmt.Sprintf(`"%s": {name: "%s",cost:%s,requirements:"%s",description:"%s"}`, breakthrough.ID, breakthrough.Name, breakthrough.Cost, breakthrough.Requirements, description)
		arraystring2 += "\n"
	}
	arraystring2 += "}"

	fmt.Fprintf(htmlFile, `const breakthroughList =%s;`, arraystring2)
	fmt.Fprintln(htmlFile, "\n")

	// Abilities array
	var arraystring3 string
	arraystring3 = "{"
	for i, ability := range regularAbilitiesArray {
		if i > 0 {
			arraystring3 += ","
		}
		description := strings.Replace(ability.Description, "\n", " ", -1)
		arraystring3 += fmt.Sprintf(`"%s": {name: "%s",type:"%s",keywords:"%s",range:"%s",description:"%s",requirements:"%s",costs:"%s",benefits:"%s"}`, ability.ID, ability.Name, ability.Type, ability.Keywords, ability.Range, description, ability.Requirements, ability.Costs, ability.Benefits)
		arraystring3 += "\n"
	}
	arraystring3 += "}"

	fmt.Fprintf(htmlFile, `const abilityList =%s;`, arraystring3)
	fmt.Fprintln(htmlFile, "\n")

	// Crafting Abilities array
	var arraystring4 string
	arraystring4 = "{"
	for i, ability := range craftingAbilitiesArray {
		if i > 0 {
			arraystring4 += ","
		}
		description := strings.Replace(ability.Description, "\n", " ", -1)
		arraystring4 += fmt.Sprintf(`"%s": {name: "%s",type:"%s",description:"%s",requirements:"%s",costs:"%s",othercosts:"%s"}`, ability.ID, ability.Name, ability.Type, description, ability.Requirements, ability.Costs, ability.OtherCosts)
		arraystring4 += "\n"
	}
	arraystring4 += "}"

	fmt.Fprintf(htmlFile, `const craftingAbilityList =%s;`, arraystring4)
	fmt.Fprintln(htmlFile, "\n")

	// End of scripts to replace

	fmt.Fprintln(htmlFile, `<!-- SCRIPTS TO REPLACE END -->`)
}
