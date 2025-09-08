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
- A CSV file named `processed_abilities.csv` in the same directory as the Go program, containing the ability data that was pre-processed by chat GPT to recognize from the description which abilities contains attacks. The query was as follows:
- Here I got a .csv for you. The data that is relevant is on the 3rd and 6th columns. the third has one of three values. If the value is true_ability then the 6th column must be processed. The 6th column describes a TTRPG mechanic where there can be Light, Heavy and/or Precise attacks. I want you to return an .xls from this .csv that would add at the end whether a Light, Heavy or Precise attacks are performed and if more than one is stated whether it must be decided between the two or if both happen simultaneously.
- ok, so the first attempt was quite good. I'd like you to try again (from scratch) but pay special attention to the cases where there is more than 1 type of attack in order to not confuse simultaneous with choosing. Also pay attention to not confuse Light with Lightning. Besides that there may be cases where there are attributes like Focus, Power, Toughness and Agility mentioned. Sometimes indicating a multiplication. I'd like you to also isolate those cases in a separate column.
