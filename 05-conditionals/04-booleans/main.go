package conditionals

func In20thCentury(year int) bool {
	if year > 1900 && year <= 2000 {
		return true
	}
	return false
}
