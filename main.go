package main

import (
	"black-hat/shodan"
	"fmt"
	"log"
	"strconv"
)

func main() {
	apiInfo := shodan.New("<api-key>").APIInfo()

	log.Println("Api-info plan: " + apiInfo.Plan)
	log.Println("Api-info https: " + strconv.FormatBool(apiInfo.Https))
	log.Println("Api-info queryCredits: " + fmt.Sprint(apiInfo.QueryCredits))
	log.Println("Api-info scanCredits: " + fmt.Sprint(apiInfo.ScanCredits))
	log.Println("Api-info telnet: " + strconv.FormatBool(apiInfo.Telnet))
	log.Println("Api-info unlocked: " + strconv.FormatBool(apiInfo.Unlocked))

}
