package art

func Scan(s string) []string {
	str1 := []string{}
	str2 := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			str1 = append(str1, str2)
			str2 = ""
			continue
		} else if i == 0 {
			continue
		} else {
			str2 = str2 + string(s[i])
		}
	}
	return str1
}
