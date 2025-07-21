package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name and surname: ")
	sc.Scan()
	name := sc.Text()

	fmt.Print("Enter your date of birth: ")
	sc.Scan()
	date := sc.Text()

	fmt.Print("Enter your height: ")
	sc.Scan()
	height := sc.Text()

	fmt.Print("Enter your weight: ")
	sc.Scan()
	weight := sc.Text()

	fmt.Print("Enter your favorite color: ")
	sc.Scan()
	color := sc.Text()

	fmt.Print("Enter your city: ")
	sc.Scan()
	city := sc.Text()

	fmt.Print("Enter your hobby: ")
	sc.Scan()
	hobby := sc.Text()

	year := date[len(date)-4:]

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Date of birth: %s\n", date)
	fmt.Printf("Height: %s\n", height)
	fmt.Printf("Weight: %s\n", weight)
	fmt.Printf("Favorite color: %s\n", color)
	fmt.Printf("City: %s\n", city)
	fmt.Printf("Hobby: %s\n", hobby)
	fmt.Printf("Year: %s\n", year)
}
