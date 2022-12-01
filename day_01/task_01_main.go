package day_01

import "bufio"
import "strings"
import "strconv"

func calculateElfValue(input []string) int {
	total := 0
	for _, strValue := range input {
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			panic(err)
		}
		total = total + intValue
	}
	return total
}

func findGreatestElfValue(input []elf) elf {
	tmpElf := elf{
		index:    0,
		calories: 0,
	}
	for index, element := range input {
		if element.calories > tmpElf.calories {
			tmpElf = input[index]
		}
	}
	return tmpElf
}

func solve(input string) elf {
	elves := []elf{}

	elfCount := 0
	elfBuffer := make([]string, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			elfValue := calculateElfValue(elfBuffer)
			temp_elf := elf{
				index:    elfCount,
				calories: elfValue,
			}
			elves = append(elves, temp_elf)
			elfBuffer = nil
			elfCount++
			continue
		}
		elfBuffer = append(elfBuffer, trimmed)
	}

	if len(elfBuffer) > 0 {
		elfValue := calculateElfValue(elfBuffer)
		temp_elf := elf{
			index:    elfCount,
			calories: elfValue,
		}
		elves = append(elves, temp_elf)
	}

	// iterate through the elves and find the best elf, extract int a map
	bestElf := findGreatestElfValue(elves)

	// return the best elf
	return bestElf
}
