package main
import (
	"github.com/bitly/go-simplejson"
	"fmt"
	"net/http"
	"io/ioutil"
)


//const url = "http://192.168.1.223:5000/v1/" + "pickup_stations?division=" + "320200"
const url = "http://192.168.1.223:5000/v1/" + "pickup_stations"

func main() {
	client := http.Client{}
	fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	js, _ := simplejson.NewJson(b)
	stations := js.Get("pickup_stations")
	fmt.Println(stations)

	for i := 0;; i++ {
		station := stations.GetIndex(i)
		id, _ := station.Get("id").Int()
		if id != 0 {
			name, _ := station.Get("name").String()
			fmt.Println(name)
			address, _ := station.Get("address").String()
			fmt.Println(address)
		} else {
			break
		}
	}
}
