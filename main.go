package main

import (
	"errors"
	"flag"
	"log"
)

var (
	operation = flag.String("op", "none", "Operation to perform on numbers 1 & 2.")
	num1      = flag.Int("num1", 0, "First Number (integer)")
	num2      = flag.Int("num2", 0, "Second Number (integer)")

	operations = map[string]fnOperation{
		"none": NoneOperator,
		"add":  AddOperator,
		"sub":  SubOperator,
		"mul":  MultiOperator,
		"div":  DivOperator,
	}
)

type fnOperation func(num1 int, num2 int) (float64, string, error)

func main() {
	flag.Parse()
	log.Println("FlagFunc - How to use flags with functions and interfaces.")
	if operations[*operation] == nil {
		log.Printf("Valid Values are:\n")
		for key, _ := range operations {
			log.Println(key)
		}
		log.Fatalf("Operation %v does not exist.", *operation)
	} else {
		fnOperator := operations[*operation]
		if result, sign, err := fnOperator(*num1, *num2); err != nil {
			log.Println("Error:", err)
		} else {
			log.Printf("%v %v %v = %v", *num1, sign, *num2, result)
		}
	}

}

func NoneOperator(num1 int, num2 int) (float64, string, error) {
	return 0.0, "", errors.New("No operation was specified!")
}

func AddOperator(num1 int, num2 int) (float64, string, error) {
	return float64(num1 + num2), "+", nil
}

func SubOperator(num1 int, num2 int) (float64, string, error) {
	return float64(num1 - num2), "-", nil
}

func MultiOperator(num1 int, num2 int) (float64, string, error) {
	return float64(num1 * num2), "*", nil
}

func DivOperator(num1 int, num2 int) (float64, string, error) {
	if num2 == 0 {
		return 0.0, "", errors.New("For division (div), num2 cannot equal 0.")
	}
	return float64(num1) / float64(num2), "/", nil
}
