package art

func MakeMap(values []string) map[rune][]string {
	j := 1
	final := make(map[rune][]string)
	for i := ' '; i <= '~'; i++ {

		final[i] = values[j : j+8]
		j = j + 9
	}
	return final
}
