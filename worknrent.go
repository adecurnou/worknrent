package main

import (
	"fmt"
)

var rent, hour_rate, tax_rate, total_hours, weekly_hours float64 = 0.0, 0.0, 0.0, 0.0, 0.0


func main() {
	welcome := "Welcome to the Work-n-Rent calculator."
	ques_rent := "What is your monthly rent?: "
	ques_hour := "How much do you earn hourly?: "
	ques_tax := "What is your average tax rate as a percentage? (no sign please): "
	result := "You need to work approximately %.2f hours for the month, and %.2f hours each week. Good luck.\n"
	
	fmt.Println(welcome)
	fmt.Print(ques_rent)
	fmt.Scanln(&rent)
	fmt.Print(ques_hour)
	fmt.Scanln(&hour_rate)
	fmt.Print(ques_tax)
	fmt.Scanln(&tax_rate)

	total_hours = rent / (hour_rate * ((1-(tax_rate * 0.01))))
	weekly_hours = total_hours / 4.0

	fmt.Printf(result, total_hours, weekly_hours)
}