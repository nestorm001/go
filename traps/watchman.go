package main

import "fmt"

func main() {
	greeting, err := secureGreet("Dewey")
	if err != nil {
		fmt.Printf("error calling secureGreet: %v", err)
		return
	}
	fmt.Println(greeting)
}
//An interface value is nil only if the inner value and type are both unset, (nil, nil).
// In particular, a nil interface will always hold a nil type. If we store a nil pointer
// of type *int inside an interface value, the inner type will be *int regardless of the
// value of the pointer: (*int, nil). Such an interface value will therefore be non-nil
// even when the pointer inside is nil.
func secureGreet(nephew string) (string, error) {
	if nephew != "Huey" && nephew != "Dewey" && nephew != "Louie" {
		return "", &NephewError{nephew}
	} else {
		return fmt.Sprint("Hello ", nephew), nil
	}
}

//func secureGreet(nephew string) (string, error) {
//	var greeting string
//	var err *NephewError
//	if nephew != "Huey" && nephew != "Dewey" && nephew != "Louie" {
//		greeting, err = "", &NephewError{nephew}
//	} else {
//		greeting, err = fmt.Sprint("Hello ", nephew), nil
//	}
//	return greeting, err
//}

type NephewError struct {
	impostor string
}

func (e *NephewError) Error() string {
	return fmt.Sprint("I don't recall having a nephew named ", e.impostor)
}