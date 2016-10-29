package main

import "fmt"

// It's nice to see that the filtering has detected the impostor.
// But Here is my greeting was definitely not supposed to be printed.

// The problem is that msg, err := greet(name) does declare a new variable msg,
// and also accidentally shadows the result parameter err. This is tricky because
// using the short declaration operator := allows mixing new and existing variables,
// and it works well on a very similar code. It depends whether the scope of the
// existing variable is exactly the same as the scope of the assignment, or not.

func main() {
	greeting, err := filterGreet("BeagleBoy")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Here is my greeting:", greeting)
}

func filterGreet(name string) (greeting string, err error) {
	if name == "" {
		greeting, err = "", fmt.Errorf("Not supposed to greet the void...")
	} else {
		msg, err := greet(name)
		if name == "Huey" || name == "Dewey" || name == "Louie" {
			greeting = msg
		} else {
			err = fmt.Errorf("not a legit nephew : %q", name)
			greeting = fmt.Sprintf("No greeting because of error [%v]", err.Error())
		}
	}
	return
}

func greet(name string) (greeting string, err error) {
	return fmt.Sprint("Hello ", name), nil
}