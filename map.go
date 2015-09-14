package main

import "fmt"

func main() {
	/* create a map*/
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New Delhi"}

	fmt.Println("Original map")

	/* print map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* delete an entry */
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("Updated map")

	/* print map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	capitals := map[string]string{"France":"Paris", "Italy":"Rome", "Japan":"Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

	days := map[string]string{"1":"Monday",
		"2":"Tuesday",
		"3":"Wednesday",
		"4":"Thursday",
		"5":"Friday",
		"6":"Saturday",
		"7":"Sunday"}

	var hasTuesday, hasHollyday bool
	for _, day := range days {
		fmt.Println(day)
		if day == "Tuesday" {
			hasTuesday = true
		}
		if day == "Hollyday" {
			hasHollyday = true
		}
	}
	if hasTuesday {
		fmt.Println("There is a Tuesday")
	} else {
		fmt.Println("There is not a Tuesday")
	}
	if hasHollyday {
		fmt.Println("There is a Hollyday")
	} else {
		fmt.Println("There is not a Hollyday")
	}

	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, isPresent = map1["Beijing"]
	if isPresent {
		fmt.Printf("The value of \"Beijin\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)
	fmt.Printf("Value is: %d\n", value)

	// delete an item:
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Washington")
	}
}
