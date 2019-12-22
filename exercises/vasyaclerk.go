package exercises

//Tickets ...
func Tickets(peopleInLine []int) string {
	var quarter, half int = 0, 0

	for _, bill := range peopleInLine {
		if bill == 25 {
			quarter++
		}
		if bill == 50 {
			if quarter > 0 {
				half++
				quarter--
			} else {
				return "NO"
			}
		}
		if bill == 100 {
			if half > 0 && quarter > 0 {
				half--
				quarter--
			} else if quarter > 2 {
				quarter = quarter - 3
			} else {
				return "NO"
			}
		}
	}
	return "YES"
}
