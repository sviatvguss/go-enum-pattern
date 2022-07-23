package go_enum_pattern

type Month int

const (
	January Month = iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func (m Month) String() string {
	switch m {
	case January:
		return "January"
	case February:
		return "February"
	case March:
		return "March"
	case April:
		return "April"
	case May:
		return "May"
	case June:
		return "June"
	case July:
		return "July"
	case August:
		return "August"
	case September:
		return "September"
	case October:
		return "October"
	case November:
		return "November"
	case December:
		return "December"
	default:
		return ""
	}
}
