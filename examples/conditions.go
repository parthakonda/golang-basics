package main

import "fmt"

func main() {
	var a = 20
	var b = 10

	var notificationCode = 0
	var notificationType string
	var _interface interface{}
	/**
	If statement
	*/
	if a < b {
		fmt.Printf("a < b is %t\n", a < b)
	}
	if a > b {
		fmt.Printf("a < b is %t\n", a > b)
	}

	/**
	If else statement
	*/
	if a < b {
		fmt.Printf("a < b is %t\n", a < b)
	} else {
		fmt.Printf("a < b is %t\n", a > b)
	}

	/**
	Nest if statement
	*/
	if a > b {
		if a == 20 {
			fmt.Printf("%d is Awesome Number", a)
		}
	}

	/**
	Expression Switch
	*/
	switch notificationCode {
	case 0:
		notificationType = "INFO"
	case 1:
		notificationType = "WARNING"
	case 2:
		notificationType = "ERROR"
	case 3:
		notificationType = "CRITICAL"
	default:
		notificationType = "UNKNOWN"
	}

	switch {
	case notificationType == "CRITICAL":
		fmt.Printf("Reporting the BUG")
	case notificationType == "WARNING":
		fmt.Printf("Print warning")
	case notificationType == "UNKNOWN":
		fmt.Printf("Report Admin")
	case notificationType == "INFO":
		fmt.Printf("Log request")
	}

	/**
	Type Switch
	*/

	switch _type := _interface.(type) {
	case nil:
		fmt.Printf("type of a is %T", _type)
	case int:
		fmt.Printf("type of a is int")
	}

}
