package art

import "strings"

func Split(s string) []string {
	arg1 := strings.ReplaceAll(s, "\n", "\\n")
	arg := strings.Split(arg1, "\\n")
	count := 0
	for _, v := range arg {
		if v == "" {
			count++
		}
	}
	if count == len(arg) {
		arg = arg[:len(arg)-1]
	}
	return arg
}
