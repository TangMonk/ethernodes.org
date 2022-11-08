package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/oschwald/geoip2-golang"

	"github.com/go-resty/resty/v2"
)

func main() {
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	counter := 0
	encodes := ""
	for {
		url := `https://ethernodes.org/data?draw=5&columns[0][data]=id&columns[0][name]=&columns[0][searchable]=true&columns[0][orderable]=true&columns[0][search][value]=&columns[0][search][regex]=false&columns[1][data]=host&columns[1][name]=&columns[1][searchable]=true&columns[1][orderable]=true&columns[1][search][value]=&columns[1][search][regex]=false&columns[2][data]=isp&columns[2][name]=&columns[2][searchable]=true&columns[2][orderable]=true&columns[2][search][value]=&columns[2][search][regex]=false&columns[3][data]=country&columns[3][name]=&columns[3][searchable]=true&columns[3][orderable]=true&columns[3][search][value]=&columns[3][search][regex]=false&columns[4][data]=client&columns[4][name]=&columns[4][searchable]=true&columns[4][orderable]=true&columns[4][search][value]=&columns[4][search][regex]=false&columns[5][data]=clientVersion&columns[5][name]=&columns[5][searchable]=true&columns[5][orderable]=true&columns[5][search][value]=&columns[5][search][regex]=false&columns[6][data]=os&columns[6][name]=&columns[6][searchable]=true&columns[6][orderable]=true&columns[6][search][value]=&columns[6][search][regex]=false&columns[7][data]=lastUpdate&columns[7][name]=&columns[7][searchable]=true&columns[7][orderable]=true&columns[7][search][value]=&columns[7][search][regex]=false&columns[8][data]=inSync&columns[8][name]=&columns[8][searchable]=true&columns[8][orderable]=true&columns[8][search][value]=&columns[8][search][regex]=false&order[0][column]=3&order[0][dir]=asc&start=%d&length=100&search[value]=&search[regex]=false&_=1667833128069`
		url = fmt.Sprintf(url, counter)

		counter = counter + 100

		client := resty.New()
		resp, err := client.R().
			Get(url)

		if err != nil {
			panic(err)
		}

		response := Response{}
		json.Unmarshal(resp.Body(), &response)

		if len(response.Data) == 0 {
			break
		}

		isChina := false
		for _, v := range response.Data {
			if v.Country == "China" || v.Country == "Hong Kong" {
				encode := fmt.Sprintf(`"enode://%s@%s:%d",`, v.ID, v.Host, v.Port)

				ip := net.ParseIP(v.Host)
				record, err := db.City(ip)
				if err != nil {
					log.Fatal(err)
				}

				if record.City.Names["en"] == "Chengdu" {
					encodes = encode + encodes
				} else {
					encodes = encodes + encode
				}

				isChina = true
			} else {
				if isChina == true {
					break
				}
			}
		}
	}
	encodes = "StaticNodes = [" + encodes + "]"
	d1 := []byte(encodes)

	err = os.WriteFile("data.txt", d1, 0644)
	if err != nil {
		panic(err)
	}

}
