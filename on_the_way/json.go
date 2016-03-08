package main
import (
	"fmt"
	"encoding/json"
	"github.com/bitly/go-simplejson" // for json get
)


type MyData struct {
	Name  string    `json:"item"`
	Other float32   `json:"amount"`
}

func main() {
	var detail MyData

	detail.Name  = "1"
	detail.Other = 2

	body, err := json.Marshal(detail)
	if err != nil {
		panic(err.Error())
	}

	js, err := simplejson.NewJson(body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(js)
}



