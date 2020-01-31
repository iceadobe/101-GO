package mypackage

type Weekday int
const (
	_ Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (day Weekday) ToString() string {
	names := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	if !day.IsValid() {
		return "Unknown Day"
	} else {
		return names[day-1]
	}
}

func (day Weekday) IsWeekend() bool {
	switch day {
	case Sunday, Saturday:
		return true
	default:
		return false
	}
}

func (day Weekday) IsValid() bool {
	if day > Sunday || day < Monday {
		return false
	}
	return true
}
