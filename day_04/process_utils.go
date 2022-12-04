package day_04

func getItemPriority(item rune) int {
	itemPriority := 0
	itemAsciiValue := int(item)
	if itemAsciiValue >= 65 && itemAsciiValue <= 90 {
		itemPriority = itemAsciiValue - 38
	}
	if itemAsciiValue >= 97 && itemAsciiValue <= 122 {
		itemPriority = itemAsciiValue - 96
	}
	return itemPriority
}
