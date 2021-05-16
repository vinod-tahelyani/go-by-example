package main

import (
	"fmt"
	"time"
)


func Switch()  {
	i := 1

	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("i is some other")
	} 

	switch time.Now().Day() {
	case int(time.Saturday), int(time.Sunday):
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	// switch time.Now().Weekday() {
    // case time.Saturday, time.Sunday:
    //     fmt.Println("It's the weekend")
    // default:
    //     fmt.Println("It's a weekday")
    // }

    // t := time.Now()
    // switch {
    // case t.Hour() < 12:
    //     fmt.Println("It's before noon")
    // default:
    //     fmt.Println("It's after noon")
    // }


	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatIAm := func (t interface{}) {
		switch t.(type) {
		case int:
			fmt.Println("I'm int")
		case bool:
			fmt.Println("I'm bool")
		case float32:
			fmt.Println("I'm float32")
		case float64:
			fmt.Println("I'm float64")
		case string:
			fmt.Println("I'm string")
		default:
			fmt.Println("Can't figure out")
		}
	}

	whatIAm(10)
	whatIAm(true)
	whatIAm(34.333)
		
}