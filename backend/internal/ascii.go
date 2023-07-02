package internal

import (
	"fmt"
)

func Create(content []string) (map[rune][]string, error) {
	if len(content) != 855 {
		return nil, fmt.Errorf("non-valid number of lines")
	}
	maps := make(map[rune][]string)
	element := rune(32)
	prev, current := 1, 9
	for current <= len(content) {
		maps[element] = content[prev:current]
		element++
		prev = current + 1
		current += 9
	}

	return maps, nil
}

// func Print(maps map[rune][]string, words []string) error {
// 	if len(os.Args) != 2 && len(os.Args) != 3 {
// 		return fmt.Errorf("Incorrect number of arguments")
// 	}
// 	for _, word := range words {
// 		if word == "" {
// 			fmt.Println()
// 			continue
// 		}
// 		for i := 0; i < 8; i++ {
// 			for _, value := range word {
// 				fmt.Print(maps[value][i])
// 			}
// 			if i != 8 {
// 				fmt.Println()
// 			}
// 		}
// 	}
// 	return nil
// }
