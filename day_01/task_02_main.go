package day_01

import "bufio"
import "strings"
import "sort"

func solve02(input string) int {
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

	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})
	total := 0
	for i := 0; i < 3; i++ {
		total = total + elves[i].calories
	}

	return total
}
