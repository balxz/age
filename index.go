package main

import (
	"fmt"
	"time"
)

func h(Year int, Month int, Datee int) int {
	now := time.Now()
	DateeTime := time.Date(Year, time.Month(Month), Datee, 0, 0, 0, 0, time.UTC)

	Age := now.Year() - DateeTime.Year()
	if now.YearDay() < DateeTime.YearDay() {
		Age--
	}
	return Age
}

func main() {
	var Year, Month, Datee int

	fmt.Print("Years? (2000 - 2024): ")
	fmt.Scan(&Year)
	fmt.Print("Month? (1-12): ")
	fmt.Scan(&Month)
	fmt.Print("Date? 1-31: ")
	fmt.Scan(&Datee)

	Age := h(Year, Month, Datee)
	fmt.Printf("Your age is %d years.\n", Age)
}
