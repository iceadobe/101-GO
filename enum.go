package main

import (
	enums "./mypackage"
	f "fmt"
)

func main() {
	var value int
	f.Print("Enter value between 1-7: ")
	f.Scanln(&value)

	day := enums.Weekday(value)

	if day.IsValid() {
		f.Printf("Your selected day is %s", day.ToString())
		if day.IsWeekend() {
			f.Println(" , enjoy your WEEKEND")
		} else {
			f.Println(" , don't be late to work")
		}
	} else {
		f.Println("There are only Seven days in a week.")
	}

	f.Println(enums.Monday.ToString(), enums.Monday)
	f.Println(enums.Tuesday.ToString(), enums.Tuesday)
	f.Println(enums.Wednesday.ToString(), enums.Wednesday)
	f.Println(enums.Thursday.ToString(), enums.Thursday)
	f.Println(enums.Friday.ToString(), enums.Friday)
	f.Println(enums.Saturday.ToString(), enums.Saturday)
	f.Println(enums.Sunday.ToString(), enums.Sunday)
}
