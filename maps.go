package main

import "fmt"

func main ()  {
	
	StudentAge := make(map[string] int)

	StudentAge["Aman"] = 25
	StudentAge["Sharif"] = 24
	StudentAge["Akhter"] = 28

	fmt.Println(StudentAge)
	fmt.Println("The length of the map is ", len(StudentAge))

	fmt.Println()
	fmt.Println()

	superhero := map[string]map[string]string{
		    "Superman" : map[string]string{
			"Real Name" : "Clark Kent",
			"City" : "Metropolis",
		    },
			"Batman" : map[string]string{
			"Real Name" : "Bruce Wayne",
			"City" : "Gotham",
			},
	}

	if temp,hero := superhero["Superman"]; hero {
		fmt.Println(temp["Real Name"], temp["City"])
	}
}