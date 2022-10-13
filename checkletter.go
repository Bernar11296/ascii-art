package art

func CheckLetter(s string) bool {
	for _, v := range s {
		if (v < ' ' || v > '~') && v != '\n' {
			return false
		}
	}
	return true
}
