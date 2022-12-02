package day_01

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
