package shodan

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const BaseUrl = "https://api.shodan.io"

type Shodan struct {
	apiKey string
}

func New(apiKey string) *Shodan {
	return &Shodan{
		apiKey: apiKey,
	}
}

func (s *Shodan) APIInfo() *APIInfo {

	url := fmt.Sprintf("%s/api-info?key=%s", BaseUrl, s.apiKey)

	log.Println(url)

	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	var apiInfo APIInfo

	err = json.NewDecoder(res.Body).Decode(&apiInfo)
	if err != nil {
		log.Fatalln(err)
	}

	return &apiInfo

}
