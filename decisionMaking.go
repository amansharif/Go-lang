package main

import "fmt"

func main()  {
	
	score := 70

	if score >50 {
		fmt.Println("Congratulations!! You have passed.")
	} else {
		fmt.Println("I regret to inform you. But you have failed.")
	}

	switch score {
	case 100 : fmt.Println("Awesome !! Full marks")
	case 70 : fmt.Println("Congratulations!! You have passed.")
	default : fmt.Println("I regret to inform you. But you have failed.")
		
	}
}