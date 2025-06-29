package main

import "fmt"

func FindLongestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]

	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}

	return longest
}

func main() {
	// Примеры использования
	fmt.Println(FindLongestString("car", "fbi", "tiger"))
	fmt.Println(FindLongestString("banana", "orange", "cherry"))
	fmt.Println(FindLongestString("two"))
	fmt.Println(FindLongestString("glhf", "short", "jungle"))
}
