package main

import (
    "time"
    "math/rand"
)

func main() {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str1 := "huangmenji"
	str2 := "malatang"

    switch(r.Intn(2)){
		case 0:
			print(str1)
		case 1:
			print(str2)
	}
    
}