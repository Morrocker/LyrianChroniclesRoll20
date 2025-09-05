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

	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_%s_show" hidden />`, classes[0].ID)
	fmt.Fprintln(htmlFile, `<div class="class-row">`)
	fmt.Fprintf(htmlFile, `<div class="class-name">%s</div>`, classes[0].Name)
	fmt.Fprintf(htmlFile, `<div class="class-tier">%d</div>`, classes[0].Tier)
	fmt.Fprintln(htmlFile, `<div class="class-level">`)
	fmt.Fprintf(htmlFile, `<select name="attr_%s_level" id="%s-level-select">`, classes[0].ID, classes[0].ID)
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
	fmt.Fprintln(htmlFile, `<div class="class-xp flex-row"`)
	fmt.Fprintf(htmlFile, `< <input readonly type="number" name="attr_%s_xp" value="@{%s_xp}" />`, classes[0].ID, classes[0].ID)
	fmt.Fprintln(htmlFile, `</div>`)
	fmt.Fprintln(htmlFile, `<div class="class-delete">`)
	fmt.Fprintln(htmlFile, `<div class="delete-item-btn">`)
	fmt.Fprintf(htmlFile, `<input type="checkbox" name="attr_%s_show" />`, classes[0].ID)
	fmt.Fprintln(htmlFile, `<span class="pseudo-button">X</span>`)
	fmt.Fprintln(htmlFile, `</div>`)
	fmt.Fprintln(htmlFile, `</div>`)
	fmt.Fprintln(htmlFile, `</div>`)
	//               <div class="class-delete">
	//                 <div class="delete-item-btn">
	//                   <input type="checkbox" name="attr_sword_show" />
	//                   <span class="pseudo-button">X</span>
	//                 </div>
	//               </div>
	//             </div>

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
