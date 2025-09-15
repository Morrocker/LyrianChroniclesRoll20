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

	fmt.Fprintln(htmlFile, `<!-- All Clases Rows START -->`)
	// for _, class := range classes {
	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_cls_%s_show" hidden />`, class.ID)
	// 	fmt.Fprintln(htmlFile, `<div class="class-row">`)
	// 	fmt.Fprintf(htmlFile, `<div class="class-name">%s</div>`, class.Name)
	// 	fmt.Fprintf(htmlFile, `<div class="class-tier">%d</div>`, class.Tier)
	// 	fmt.Fprintln(htmlFile, `<div class="class-level">`)
	// 	fmt.Fprintf(htmlFile, `<select name="attr_%s_level" id="%s-level-select">`, class.ID, class.ID)
	// 	fmt.Fprintln(htmlFile, `<option value="1">1</option>`)
	// 	fmt.Fprintln(htmlFile, `<option value="2">2</option>`)
	// 	fmt.Fprintln(htmlFile, `<option value="3">3</option>`)
	// 	fmt.Fprintln(htmlFile, `<option value="4">4</option>`)
	// 	fmt.Fprintln(htmlFile, `<option value="5">5</option>`)
	// 	fmt.Fprintln(htmlFile, `<option value="6">6</option>`)
	// 	fmt.Fprintln(htmlFile, `<option value="7">7</option>`)
	// 	fmt.Fprintln(htmlFile, `<option value="8">8</option>`)
	// 	fmt.Fprintln(htmlFile, `</select>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintln(htmlFile, `<div class="class-xp flex-row">`)
	// 	fmt.Fprintf(htmlFile, `<input readonly type="number" name="attr_%s_xp" value="@{%s_xp}" />`, class.ID, class.ID)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintln(htmlFile, `<div class="class-delete">`)
	// 	fmt.Fprintln(htmlFile, `<div class="delete-item-btn">`)
	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_cls_%s_show" />`, class.ID)
	// 	fmt.Fprintln(htmlFile, `<span class="pseudo-button">✖</span>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// }

	fmt.Fprintln(htmlFile, "<!-- All Clases Rows END -->")

	fmt.Fprintln(htmlFile, "<!-- All Clases Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)
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

	fmt.Fprintln(htmlFile, "<!-- All Breakthroughs Table List START -->")

	// for _, breakthrough := range breakthroughs {
	// 	//

	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_bt_%s_show" hidden />`, breakthrough.ID)
	// 	fmt.Fprintln(htmlFile, `<div class="breakthrough-container">`)
	// 	fmt.Fprintf(htmlFile, `<div class="breakthrough-row">
	//                 <div class="breakthrough-name">
	//                   <button
	//                     type="roll"
	//                     name="roll_bt_%s"
	//                     value="&{template:default} {{name=@{name} ❖ Breakthrough: %s}}{{Cost=%d}} {{Requirements=%s}} {{ Description=%s}}"
	//                     tabindex="-1"
	//                   >
	//                     %s
	//                   </button>
	//                 </div>`, breakthrough.ID, breakthrough.Name, breakthrough.Cost, breakthrough.Requirements, breakthrough.Description, breakthrough.Name)
	// 	fmt.Fprintf(htmlFile, `<div class="breakthrough-cost">%d</div>`, breakthrough.Cost)
	// 	// fmt.Fprintf(htmlFile, `<div class="breakthrough-requirements">%s</div>`, breakthrough.Requirements)
	// 	fmt.Fprintln(htmlFile, `<div class="breakthrough-delete">`)
	// 	fmt.Fprintln(htmlFile, `<div class="delete-item-btn thin">`)
	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_bt_%s_details_show" />`, breakthrough.ID)
	// 	fmt.Fprintln(htmlFile, `<span class="view-details">ℹ</span>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintln(htmlFile, `<div class="delete-item-btn thin">`)
	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_bt_%s_show" />`, breakthrough.ID)
	// 	fmt.Fprintln(htmlFile, `<span class="pseudo-button">✖</span>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_bt_%s_details_show" hidden />`, breakthrough.ID)
	// 	fmt.Fprintln(htmlFile, `<div class="ability-details">`)
	// 	fmt.Fprintln(htmlFile, `<div class="detail-row">`)
	// 	fmt.Fprintln(htmlFile, `<div class="ability-label">XP Cost:</div>`)
	// 	fmt.Fprintf(htmlFile, `<div class="ability-value">%d</div>`, breakthrough.Cost)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	var req string = "--"
	// 	if breakthrough.Requirements != "" {
	// 		req = breakthrough.Requirements
	// 	}
	// 	fmt.Fprintln(htmlFile, `<div class="detail-row">`)
	// 	fmt.Fprintln(htmlFile, `<div class="ability-label">Requirement:</div>`)
	// 	fmt.Fprintf(htmlFile, `<div class="ability-value">%s</div>`, req)
	// 	fmt.Fprintln(htmlFile, `</div>`)

	// 	fmt.Fprintln(htmlFile, `<div class="detail-row">`)
	// 	fmt.Fprintln(htmlFile, `<div class="ability-label">Description:</div>`)
	// 	fmt.Fprintf(htmlFile, `<div class="ability-value">%s</div>`, breakthrough.Description)
	// 	fmt.Fprintln(htmlFile, `</div>`)

	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)

	// }

	fmt.Fprintln(htmlFile, "<!-- All Breakthroughs Table List END -->")

	fmt.Fprintln(htmlFile, "<!-- All Breakthroughs Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)

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

	for _, row := range records3 {
		if row[2] == "crafting_ability" {
			continue
		}
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
		fmt.Printf("id: %s, name: %s, type: %s\n", ability.ID, ability.Name, ability.Type)

		regularAbilities = append(regularAbilities, ability)
	}

	type AbilityAtkCount struct {
		Name string
		LAtk bool
		MAtk bool
		HAtk bool
	}

	fmt.Fprintln(htmlFile, "<!-- All Abilities List START -->")

	// for _, ability := range regularAbilities {
	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_abil_%s_show" hidden />`, ability.ID)
	// 	fmt.Fprintln(htmlFile, `<div class="ability-container">`)
	// 	fmt.Fprintln(htmlFile, `<div class="ability-row">`)
	// 	fmt.Fprintln(htmlFile, `<div class="ability-name">`)
	// 	fmt.Fprintln(htmlFile, `<button type="action" name="act_post_ability_info"`)
	// 	fmt.Fprintf(htmlFile, `data-info-label="%s"`, ability.Name)
	// 	fmt.Fprintf(htmlFile, `data-roll-label="%s"`, ability.Name)
	// 	if ability.Type == "true_ability" {
	// 		fmt.Fprintf(htmlFile, `data-info-type="true"`)
	// 		fmt.Fprintf(htmlFile, `data-info-keywords="%s"`, ability.Keywords)
	// 		fmt.Fprintf(htmlFile, `data-info-range="%s"`, ability.Range)
	// 		fmt.Fprintf(htmlFile, `data-info-description="%s"`, ability.Description)
	// 		fmt.Fprintf(htmlFile, `data-info-requirement="%s"`, ability.Requirements)
	// 		fmt.Fprintf(htmlFile, `data-info-cost="%s"`, ability.Costs)
	// 	} else if ability.Type == "key_ability" {
	// 		fmt.Fprintf(htmlFile, `data-info-type="key"`)
	// 		fmt.Fprintf(htmlFile, `data-info-benefits="%s"`, ability.Benefits)
	// 	}
	// 	// Determine attack types
	// 	if strings.Contains(ability.AtkTypes, "Light / Heavy / Precise") {
	// 		fmt.Fprintln(htmlFile, `data-roll-light-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-heavy-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-precise-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-choice="true"`)
	// 	} else if strings.Contains(ability.AtkTypes, "Light / Heavy") {
	// 		fmt.Fprintln(htmlFile, `data-roll-light-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-heavy-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-precise-atk="false"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-choice="true"`)
	// 	} else if strings.Contains(ability.AtkTypes, "Heavy / Precise") {
	// 		fmt.Fprintln(htmlFile, `data-roll-light-atk="false"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-heavy-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-precise-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-choice="true"`)
	// 	} else if strings.Contains(ability.AtkTypes, "Light + Heavy") {
	// 		fmt.Fprintln(htmlFile, `data-roll-light-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-heavy-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-precise-atk="false"`)
	// 	} else if strings.Contains(ability.AtkTypes, "Light + Precise") {
	// 		fmt.Fprintln(htmlFile, `data-roll-light-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-heavy-atk="false"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-precise-atk="true"`)
	// 	} else if strings.Contains(ability.AtkTypes, "Heavy + Precise") {
	// 		fmt.Fprintln(htmlFile, `data-roll-light-atk="false"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-heavy-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-precise-atk="true"`)
	// 	} else if strings.Contains(ability.AtkTypes, "Light") {
	// 		fmt.Fprintln(htmlFile, `data-roll-light-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-heavy-atk="false"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-precise-atk="false"`)
	// 	} else if strings.Contains(ability.AtkTypes, "Heavy") {
	// 		fmt.Fprintln(htmlFile, `data-roll-light-atk="false"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-heavy-atk="true"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-precise-atk="false"`)
	// 	} else if strings.Contains(ability.AtkTypes, "Precise") {
	// 		fmt.Fprintln(htmlFile, `data-roll-light-atk="false"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-heavy-atk="false"`)
	// 		fmt.Fprintln(htmlFile, `data-roll-precise-atk="true"`)
	// 	}
	// 	fmt.Fprintf(htmlFile, `tabindex="-1">%s</button>`, ability.Name)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	if ability.Range == "Melee Weapon Range" {
	// 		ability.Range = "MWR"
	// 	} else if ability.Range == "Ranged Weapon Range" {
	// 		ability.Range = "RWR"
	// 	}

	// 	fmt.Fprintf(htmlFile, `<div class="ability-range">%s</div>`, ability.Range)
	// 	fmt.Fprintf(htmlFile, `<div class="ability-costs">%s</div>`, ability.Costs)
	// 	fmt.Fprintln(htmlFile, `<div class="ability-delete flex-row">`)
	// 	fmt.Fprintln(htmlFile, `<div class="delete-item-btn thin">`)
	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_abil_%s_favorite"/>`, ability.ID)
	// 	fmt.Fprintln(htmlFile, `<span class="favorite">★</span></div>`)
	// 	fmt.Fprintln(htmlFile, `<div class="delete-item-btn thin">`)
	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_%s_details_show"/>`, ability.ID)
	// 	fmt.Fprintln(htmlFile, `<span class="view-details">ℹ</span></div>`)
	// 	fmt.Fprintln(htmlFile, `<div class="delete-item-btn thin">`)
	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_abil_%s_show"/>`, ability.ID)
	// 	fmt.Fprintln(htmlFile, `<span class="pseudo-button">✖</span></div>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_%s_details_show" hidden />`, ability.ID)
	// 	fmt.Fprintln(htmlFile, `<div class="ability-details">`)
	// 	if ability.Type == "key_ability" {
	// 		fmt.Fprintln(htmlFile, `<div class="detail-row">`)
	// 		fmt.Fprintln(htmlFile, `<div class="ability-label">Benefits:</div>`)
	// 		fmt.Fprintln(htmlFile, `</div>`)
	// 		fmt.Fprintln(htmlFile, `<div class="detail-row">`)
	// 		fmt.Fprintln(htmlFile, `<ul class="ability-value">`)
	// 		splitBenefits := strings.Split(ability.Benefits, ".,")
	// 		for _, benefit := range splitBenefits {
	// 			fmt.Fprintf(htmlFile, `<li>%s</li>`, benefit)

	// 		}
	// 		fmt.Fprintln(htmlFile, `</ul>`)
	// 		fmt.Fprintln(htmlFile, `</div>`)
	// 		// fmt.Fprintln(htmlFile, `<div class="ability-value">%s</div>`, ability.Benefits)
	// 	}
	// 	if ability.Type == "true_ability" {
	// 		if ability.Keywords != "" {
	// 			fmt.Fprintln(htmlFile, `<div class="detail-row">`)
	// 			fmt.Fprintln(htmlFile, `<div class="ability-label">Keywords:</div>`)
	// 			fmt.Fprintf(htmlFile, `<div class="ability-value">%s</div>`, ability.Keywords)
	// 			fmt.Fprintln(htmlFile, `</div>`)
	// 		}
	// 		if ability.Range != "" {
	// 			fmt.Fprintln(htmlFile, `<div class="detail-row">`)
	// 			fmt.Fprintln(htmlFile, `<div class="ability-label">Range:</div>`)
	// 			fmt.Fprintf(htmlFile, `<div class="ability-value">%s</div>`, ability.Range)
	// 			fmt.Fprintln(htmlFile, `</div>`)
	// 		}
	// 		fmt.Fprintln(htmlFile, `<div class="detail-row">`)
	// 		fmt.Fprintln(htmlFile, `<div class="ability-label">Description:</div>`)
	// 		fmt.Fprintf(htmlFile, `<div class="ability-value">%s</div>`, ability.Description)
	// 		fmt.Fprintln(htmlFile, `</div>`)
	// 		if ability.Requirements != "" {
	// 			fmt.Fprintln(htmlFile, `<div class="detail-row">`)
	// 			fmt.Fprintln(htmlFile, `<div class="ability-label">Requirement:</div>`)
	// 			fmt.Fprintf(htmlFile, `<div class="ability-value">%s</div>`, ability.Requirements)
	// 			fmt.Fprintln(htmlFile, `</div>`)
	// 		}
	// 		if ability.Costs != "" {
	// 			fmt.Fprintln(htmlFile, `<div class="detail-row">`)
	// 			fmt.Fprintln(htmlFile, `<div class="ability-label">Costs:</div>`)
	// 			fmt.Fprintf(htmlFile, `<div class="ability-value">%s</div>`, ability.Costs)
	// 			fmt.Fprintln(htmlFile, `</div>`)
	// 		}
	// 	}
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// 	fmt.Fprintln(htmlFile, `</div>`)
	// }

	fmt.Fprintln(htmlFile, "<!-- All Abilities List END -->")

	fmt.Fprintln(htmlFile, "<!-- All Abilities Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)

	for _, ability := range regularAbilities {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, ability.ID, ability.Name)
		fmt.Fprintln(htmlFile, "")
	}

	fmt.Fprintln(htmlFile, "<!-- All abilities Picker List END -->")

	// ============================= ABILITIES END =============================

	fmt.Fprintln(htmlFile, "<!-- SCRIPTS TO REPLACE START -->")

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

	fmt.Fprintln(htmlFile, `<!-- SCRIPTS TO REPLACE END -->`)
	// 	// Write basic HTML structure
	// 	fmt.Fprintln(htmlFile, "<!DOCTYPE html>")
	// 	fmt.Fprintln(htmlFile, "<html><body><table border='1'>")

	// 	// Write table rows
	// 	for _, row := range records {
	// 		fmt.Fprint(htmlFile, "<tr>")
	// 		for _, col := range row {
	// 			fmt.Fprintf(htmlFile, "<td>%s</td>", col)
	// 		}
	// 		fmt.Fprintln(htmlFile, "</tr>")
	// 	}

	// fmt.Fprintln(htmlFile, "</table></body></html>")
	// fmt.Println("HTML file generated: output.html")
}
