package move

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type SubmarinePosition struct {
	depth      int
	horizontal int
	aim        int
}

func MoveSubmarine(commands []string) int {
	submarinePosition := SubmarinePosition{depth: 0, horizontal: 0}

	for _, val := range commands {
		commandAndChange := GetCommandAndChangeFromDataLine(val)

		if strings.ToLower(strings.Trim(commandAndChange.command, " \n")) == "forward" {
			submarinePosition.horizontal += commandAndChange.change
		} else if strings.ToLower(strings.Trim(commandAndChange.command, " \n")) == "up" {
			submarinePosition.depth -= commandAndChange.change
		} else if strings.ToLower(strings.Trim(commandAndChange.command, " \n")) == "down" {
			submarinePosition.depth += commandAndChange.change
		}
	}

	return submarinePosition.depth * submarinePosition.horizontal
}

func MoveSubmarineIncludingAim(commands []string) int {
	submarinePosition := SubmarinePosition{depth: 0, horizontal: 0, aim: 0}

	for _, val := range commands {
		commandAndChange := GetCommandAndChangeFromDataLine(val)

		if strings.ToLower(strings.Trim(commandAndChange.command, " \n")) == "forward" {
			submarinePosition.horizontal += commandAndChange.change
			submarinePosition.depth += submarinePosition.aim * commandAndChange.change
		} else if strings.ToLower(strings.Trim(commandAndChange.command, " \n")) == "up" {
			submarinePosition.aim -= commandAndChange.change
		} else if strings.ToLower(strings.Trim(commandAndChange.command, " \n")) == "down" {
			submarinePosition.aim += commandAndChange.change
		}
	}

	return submarinePosition.depth * submarinePosition.horizontal
}

func GetCommandAndChangeFromDataLine(dataLine string) (result struct {
	command string
	change  int
}) {
	commandChangeArray := strings.Split(dataLine, " ")
	if len(commandChangeArray) != 2 {
		log.Fatal("The given line is not a correct submarine command. Correct submarine command goes like\ncommand change")
	}

	result.command = commandChangeArray[0]
	changeAsString := commandChangeArray[1]

	change, err := strconv.Atoi(changeAsString)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	result.change = change

	return
}
