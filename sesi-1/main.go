package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	duration := 15 * time.Second
	doEvery(duration)

}

func getStatus(water, wind int) (waterStatus, windStatus string) {
	if water < 5 {
		waterStatus = "aman"
	} else if water > 5 && water <= 8 {
		waterStatus = "siaga"
	} else if water > 8 {
		waterStatus = "bahaya"
	}

	if wind < 6 {
		windStatus = "aman"
	} else if wind >= 7 && wind <= 15 {
		windStatus = "siaga"
	} else if wind > 15 {
		windStatus = "bahaya"
	}

	return waterStatus, windStatus
}

func doEvery(d time.Duration) {
	for range time.Tick(d) {
		water := rand.Intn(100)
		wind := rand.Intn(100)
		waterStatus, windStatus := getStatus(water, wind)

		data := map[string]int{
			"water": water,
			"wind":  wind,
		}

		requestJson, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts",
			bytes.NewBuffer(requestJson))
		req.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body))
		log.Println("status water : ", waterStatus)
		log.Println("status wind : ", windStatus)

	}
}
