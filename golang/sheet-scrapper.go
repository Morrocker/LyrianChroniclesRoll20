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
	for _, class := range classes {
		fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_cls_%s_show" hidden />`, class.ID)
		fmt.Fprintln(htmlFile, `<div class="class-row">`)
		fmt.Fprintf(htmlFile, `<div class="class-name">%s</div>`, class.Name)
		fmt.Fprintf(htmlFile, `<div class="class-tier">%d</div>`, class.Tier)
		fmt.Fprintln(htmlFile, `<div class="class-level">`)
		fmt.Fprintf(htmlFile, `<select name="attr_%s_level" id="%s-level-select">`, class.ID, class.ID)
		fmt.Fprintln(htmlFile, `<option value="1">1</option>`)
		fmt.Fprintln(htmlFile, `<option value="2">2</option>`)
		fmt.Fprintln(htmlFile, `<option value="3">3</option>`)
		fmt.Fprintln(htmlFile, `<option value="4">4</option>`)
		fmt.Fprintln(htmlFile, `<option value="5">5</option>`)
		fmt.Fprintln(htmlFile, `<option value="6">6</option>`)
		fmt.Fprintln(htmlFile, `<option value="7">7</option>`)
		fmt.Fprintln(htmlFile, `<option value="8">8</option>`)
		fmt.Fprintln(htmlFile, `</select>`)
		fmt.Fprintln(htmlFile, `</div>`)
		fmt.Fprintln(htmlFile, `<div class="class-xp flex-row">`)
		fmt.Fprintf(htmlFile, `<input readonly type="number" name="attr_%s_xp" value="@{%s_xp}" />`, class.ID, class.ID)
		fmt.Fprintln(htmlFile, `</div>`)
		fmt.Fprintln(htmlFile, `<div class="class-delete">`)
		fmt.Fprintln(htmlFile, `<div class="delete-item-btn">`)
		fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_cls_%s_show" />`, class.ID)
		fmt.Fprintln(htmlFile, `<span class="pseudo-button">X</span>`)
		fmt.Fprintln(htmlFile, `</div>`)
		fmt.Fprintln(htmlFile, `</div>`)
		fmt.Fprintln(htmlFile, `</div>`)
	}

	fmt.Fprintln(htmlFile, "<!-- All Clases Rows END -->")

	fmt.Fprintln(htmlFile, "<!-- All Clases Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)
	for _, class := range classes {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, class.ID, class.Name)
	}
	fmt.Fprintln(htmlFile, "<!-- All Clases Picker List END -->")

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

	for _, breakthrough := range breakthroughs {
		//

		fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_bt_%s_show" hidden />`, breakthrough.ID)
		fmt.Fprintf(htmlFile, `<div class="breakthrough-row">
                    <div class="breakthrough-name">
                      <button
                        type="roll"
                        name="roll_bt_%s"
                        value="&{template:default} {{name=@{name} ❖ Breakthrough: %s}}{{Cost=%d}} {{Requirements=%s}} {{ Description=%s}}"
                        tabindex="-1"
                      >
                        %s
                      </button>
                    </div>`, breakthrough.ID, breakthrough.Name, breakthrough.Cost, breakthrough.Requirements, breakthrough.Description, breakthrough.Name)
		fmt.Fprintf(htmlFile, `<div class="breakthrough-cost">%d</div>`, breakthrough.Cost)
		fmt.Fprintf(htmlFile, `<div class="breakthrough-requirements">%s</div>`, breakthrough.Requirements)
		fmt.Fprintln(htmlFile, `<div class="breakthrough-delete">`)
		fmt.Fprintln(htmlFile, `<div class="delete-item-btn">`)
		fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_bt_%s_show" />`, breakthrough.ID)
		fmt.Fprintln(htmlFile, `<span class="pseudo-button">X</span>`)
		fmt.Fprintln(htmlFile, `</div>`)
		fmt.Fprintln(htmlFile, `</div>`)
		fmt.Fprintln(htmlFile, `</div>`)

	}

	fmt.Fprintln(htmlFile, "<!-- All Breakthroughs Table List END -->")

	fmt.Fprintln(htmlFile, "<!-- All Breakthroughs Picker List START -->")

	fmt.Fprintln(htmlFile, `<option value="">--Select--</option>`)

	for _, breakthrough := range breakthroughs {
		fmt.Fprintf(htmlFile, `<option value="%s">%s</option>`, breakthrough.ID, breakthrough.Name)
	}

	fmt.Fprintln(htmlFile, "<!-- All Breakthroughs Picker List END -->")

	fmt.Fprintln(htmlFile, "<!-- SCRIPTS TO REPLACE START -->")

	var arraystring string
	arraystring = "["
	for i, class := range classes {
		if i > 0 {
			arraystring += ","
		}
		arraystring += fmt.Sprintf("{ class:\"%s\", tier:%d}", class.ID, class.Tier)
	}
	arraystring += "]"

	fmt.Fprintf(htmlFile, `const classList =%s;`, arraystring)
	fmt.Fprintln(htmlFile, "\n")

	// fmt.Fprintln(htmlFile, `// Udates the class total XP when the level is changed`)
	// fmt.Fprintln(htmlFile, `classList.forEach(({ class: className, tier}) => {
	// 	on(`+"`change:${className}_level`"+`, function () {
	// 	  getAttrs([`+"`${className}_level`"+`], function (values) {
	// 		let level = parseInt(values[`+"`${className}_level`"+`], 10) || 1;
	// 		let xp = (level - 1) * 100 + tier*100;
	// 		let update = {};
	// 		update[`+"`${className}_xp`"+`] = xp;
	// 		setAttrs(update);
	// 	  });
	// 	});
	//   });
	//   `)

	// fmt.Fprintf(htmlFile, "\n")
	// fmt.Fprintln(htmlFile, "// Reset class levels and XP when the class is deleted")
	// fmt.Fprintln(htmlFile, `classList.forEach(({ class: className, tier}) => {
	//   on(`+"`change:${className}_show`"+`, function () {
	// 	getAttrs([`+"`${className}_show`"+`], function (values) {
	// 	  let update = {};
	// 	  update[`+"`${className}_level`"+`] = 1;
	// 	  update[`+"`${className}_xp`"+`] = tier*100;
	// 	  setAttrs(update);
	// 	});
	//   });
	// });
	// `)

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
