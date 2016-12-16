package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//1. Skapa ett konto hos api servern och spara nyckeln
	openweatherAccountID := "ffe0d100aea24a36ee53bafb5712ef19"
	stad := "London,uk"

	//2. Ta reda på API anropets endpoint
	apiEndPoint := "http://api.openweathermap.org"

	//3. API anrop för att hämta vädret i en stad
	// api.openweathermap.org/data/2.5/weather?q= London,uk &APPID= ffe0d100aea24a36ee53bafb5712ef19
	openweatherURLStr := apiEndPoint + "/data/2.5/weather?q=" + stad + "&APPID=" + openweatherAccountID

	//4. Gör API anropet
	openweatherResp, err := http.Get(openweatherURLStr)

	// 5. Kolla om det gick bra
	if err != nil {
		panic(err)
	}

	// 6. Om inget gick fel - spara resultet som servern skickade
	// Read the body of the response and store (will be in bytes)
	openweatherRespBytes, err := ioutil.ReadAll(openweatherResp.Body)
	if err != nil {
		fmt.Println("Detta gick fel: ", err)
	}

	// 7. Unmarshall alla bytes från body in i en map
	var resultatData map[string]interface{}
	if err := json.Unmarshal(openweatherRespBytes, &resultatData); err != nil {
		panic(err)
	}

	// 8. Returnera resultatet
	// {"coord":{"lon":-0.13,"lat":51.51},"weather":[{"id":721,"main":"Haze","description":"haze","icon":"50d"},{"id":701,"main":"Mist","description":"mist","icon":"50d"}],"base":"stations","main":{"temp":282.07,"pressure":1025,"humidity":93,"temp_min":281.15,"temp_max":283.15},"visibility":5000,"wind":{"speed":2.1,"deg":90},"clouds":{"all":88},"dt":1481889000,"sys":{"type":1,"id":5091,"message":2.1385,"country":"GB","sunrise":1481875272,"sunset":1481903517},"id":2643743,"name":"London","cod":200}{"speed":2.1,"deg":90},"clouds":{"all":88},"dt":1481889000,"sys":{"type":1,"id":5091,"message":2.1385,"country":"GB","sunrise":1481875272,"sunset":1481903517},"id":2643743,"name":"London","cod":200}

	fmt.Printf("%+v", resultatData)
}
