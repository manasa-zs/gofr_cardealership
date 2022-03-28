package LeapYear

/*LeapYear is used find whether a given input is leap year or not*/
func LeapYear(input int) string {
	if input%4 == 0 && input%100 != 0 {
		return "Leap year"
	} else if input%100 == 0 && input%400 == 0 {
		return "Leap year"
	}
	return "Non Leap Year"
}
