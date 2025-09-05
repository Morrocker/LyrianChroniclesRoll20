# LyrianChronicles sheet-to-roll20 .csv converter

This is a simple Go program that reads a CSV file containing class data for the Lyrian Chronicles RPG and converts it into HTML code needed to update the custom Roll20 character sheet for the game.

## Disclaimer
This program relies on the asumption that the input CSV file is well-formed and contains the expected columns. It does not include extensive error handling for malformed input.
Any change on the basic structure of the CSV file or in the way the classes, items, abilities, etc. are represented in the Roll20 sheet may break this program.

Also, its output is intended to be used as a part of a larger HTML document (the Roll20 character sheet). It does not produce a complete HTML document on its own.

The HTML will be un-indented for simplicity, but you may want to format it for better readability.

## Prerequisites
- Go programming language installed on your machine.
- A CSV file named `classes.csv` in the same directory as the Go program, containing the class data.
- A CSV file named `items.csv` in the same directory as the Go program, containing the item data.
- A CSV file named `abilities.csv` in the same directory as the Go program, containing the ability data.
