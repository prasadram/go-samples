package main

import "fmt"

func main() {
	var x map[string]int
	x = make(map[string]int)
	x["Jan"] = 1
	x["Feb"] = 2
	x["March"] = 3
	fmt.Println(x)
	delete(x, "Feb")
	fmt.Println(x)
	if value, ok := x["March"]; ok {
		fmt.Println(value, ok)
	}

	months := map[string]int {
	 "Jan": 1,
	 "Feb": 2,
	 "March": 3,
	 "April": 4,
	 "May": 5,
	 "June": 6,
	 "July": 7,
	 "August": 8,
	 "September": 9,
	 "October": 10,
	 "November": 11,
	 "December": 12,
	}
	fmt.Println("Value of Novemnber:", months["November"])
	fmt.Println("value of March:", x["March"])
	studentage := make(map[string] int)
	studentage["John"] = 34
	fmt.Println(studentage["John"])
	fmt.Println("length of map:", len(studentage))
	studentage["Bob"] = 55
	fmt.Println("length of map:", len(studentage))
	fmt.Println("map:", studentage)
	fmt.Println("deleting bob from map" )
	delete(studentage, "Bob")
	fmt.Println("length of map:", len(studentage))

	superhero := map[string]map[string]string {
		"batman": map[string]string {
			"realname" : "Bruce Wayne",
			"city" : "Gotham City",
		},
		"superman": map[string]string {
			"realname": "Clark Kent",
			"city": "Metropolis",
		},
	}
	fmt.Println(superhero)
	if temp, hero := superhero["superman"]; hero {
		fmt.Println(temp["realname"], temp["city"])
	}
}

