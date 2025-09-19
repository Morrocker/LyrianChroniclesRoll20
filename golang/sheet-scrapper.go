package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the CSV file
	f, err := os.Open("classes.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	records = records[1:] // Skip header row

	type Class struct {
		ID   string
		Name string
		Tier int
	}

	var classes []Class

	for _, row := range records {
		tier, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}
		id := row[0]
		id = strings.ReplaceAll(id, " ", "_")
		id = strings.ToLower(id)

		c := Class{
			ID:   id,
			Name: row[0],
			Tier: tier,
		}
		classes = append(classes, c)
		fmt.Printf("id: %s, name: %s, tier: %d\n", c.ID, c.Name, tier)
	}

	// Create the HTML file
	htmlFile, err := os.Create("update-data.html")
	if err != nil {
		panic(err)
	}
	defer htmlFile.Close()

	fmt.Fprintln(htmlFile, "<!-- All Clases Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)
	fmt.Fprintln(htmlFile, `<option value="custom">Custom</option>`)
	for _, class := range classes {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, class.ID, class.Name)
	}
	fmt.Fprintln(htmlFile, "<!-- All Clases Picker List END -->")

	// ============================= BREAKTHROUGHS =============================

	f2, err := os.Open("breakthroughs.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader2 := csv.NewReader(f2)
	records2, err := reader2.ReadAll()
	if err != nil {
		panic(err)
	}

	records2 = records2[1:] // Skip header row

	type Breakthrough struct {
		ID           string
		Name         string
		Cost         int
		Requirements string
		Description  string
	}

	var breakthroughs []Breakthrough

	for _, row := range records2 {
		cost, err := strconv.Atoi(row[1])
		if err != nil {
			panic(err)
		}
		id := row[0]
		id = strings.ReplaceAll(id, " ", "_")
		id = strings.ToLower(id)

		//if id starts with a number, prepend with x
		if id[0] >= '0' && id[0] <= '9' {
			id = "x" + id
		}

		b := Breakthrough{
			ID:           id,
			Name:         row[0],
			Cost:         cost,
			Requirements: strings.ReplaceAll(row[2], `"`, "'"),
			Description:  strings.ReplaceAll(row[3], `"`, "'"),
		}
		breakthroughs = append(breakthroughs, b)
		fmt.Printf("id: %s, name: %s, cost: %d\n", b.ID, b.Name, cost)
	}

	fmt.Fprintln(htmlFile, "<!-- All Breakthroughs Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)
	fmt.Fprintln(htmlFile, `<option value="custom">Custom</option>`)

	for _, breakthrough := range breakthroughs {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, breakthrough.ID, breakthrough.Name)
	}

	fmt.Fprintln(htmlFile, "<!-- All Breakthroughs Picker List END -->")

	// ============================= BREAKTHROUGHS END =============================

	// ============================= ABILITIES START =============================

	f3, err := os.Open("abilities_processed.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader3 := csv.NewReader(f3)
	records3, err := reader3.ReadAll()
	if err != nil {
		panic(err)
	}

	records3 = records3[1:] // Skip header row

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

	var regularAbilities []Ability
	var craftingAbilities []Ability

	for _, row := range records3 {
		if row[2] == "crafting_ability" || row[2] == "gathering_ability" {
			ability := Ability{
				ID:           row[0],
				Name:         row[1],
				Type:         row[2],
				Description:  strings.ReplaceAll(row[5], `"`, "'"),
				Requirements: strings.ReplaceAll(row[6], `"`, "'"),
				Costs:        strings.ReplaceAll(row[7], `"`, "'"),
				OtherCosts:   strings.ReplaceAll(row[8], `"`, "'"),
			}

			craftingAbilities = append(craftingAbilities, ability)

		} else {
			ability := Ability{
				ID:           row[0],
				Name:         row[1],
				Type:         row[2],
				Keywords:     row[3],
				Range:        row[4],
				Description:  strings.ReplaceAll(row[5], `"`, "'"),
				Requirements: strings.ReplaceAll(row[6], `"`, "'"),
				Costs:        strings.ReplaceAll(row[7], `"`, "'"),
				OtherCosts:   strings.ReplaceAll(row[8], `"`, "'"),
				Benefits:     strings.ReplaceAll(row[12], `"`, "'"),
				AtkTypes:     row[18],
			}

			regularAbilities = append(regularAbilities, ability)

		}
	}

	type AbilityAtkCount struct {
		Name string
		LAtk bool
		MAtk bool
		HAtk bool
	}

	fmt.Fprintln(htmlFile, "<!-- All Abilities Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)
	fmt.Fprintln(htmlFile, `<option value="custom">Custom</option>`)

	for _, ability := range regularAbilities {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, ability.ID, ability.Name)
		fmt.Fprintln(htmlFile, "")
	}

	fmt.Fprintln(htmlFile, "<!-- All abilities Picker List END -->")

	// Crafting abilities
	fmt.Fprintln(htmlFile, "<!-- All Crafting Abilities Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)
	fmt.Fprintln(htmlFile, `<option value="custom">Custom</option>`)

	for _, ability := range craftingAbilities {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, ability.ID, ability.Name)
		fmt.Fprintln(htmlFile, "")
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
		arraystring2 += fmt.Sprintf(`"%s": {name: "%s",cost:%d,requirements:"%s",description:"%s"}`, breakthrough.ID, breakthrough.Name, breakthrough.Cost, breakthrough.Requirements, description)
		arraystring2 += "\n"
	}
	arraystring2 += "}"

	fmt.Fprintf(htmlFile, `const breakthroughList =%s;`, arraystring2)
	fmt.Fprintln(htmlFile, "\n")

	// Abilities array
	var arraystring3 string
	arraystring3 = "{"
	for i, ability := range regularAbilities {
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
	for i, ability := range craftingAbilities {
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
