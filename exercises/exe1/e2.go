package main

func getgrade(score int) string {
	switch {
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score <= 89:
		return "B"
	case score >= 70 && score <= 79:
		return "C"
	case score < 70:
		return "D"
	default:
		return "Invalid"
	}

}
