package art

import "fmt"

func Print(s []string, final map[rune][]string) {
	for _, val := range s {
		runes := []rune(val)

		str := []string{}

		for i := 0; i < 8; i++ {
			for _, v := range runes {
				str = append(str, final[v][i])
			}
		}
		for i := 0; i < len(str); i++ {
			if i == 0 {
				fmt.Print(str[i])
			} else if i%len(runes) == 0 && i != 0 {
				fmt.Println()
				fmt.Print(str[i])
			} else {
				fmt.Print(str[i])
			}
		}
		fmt.Println()
	}
}
