package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"
	ticker := time.NewTicker(15 * time.Second) // set ticker untuk mengirim request setiap 15 detik
	timeout := time.NewTimer(5 * time.Minute)  // set timer untuk berhenti setelah 5 menit

	for {
		select {
		case <-ticker.C:

			data := map[string]int{
				"water": rand.Intn(100) + 1,
				"wind":  rand.Intn(100) + 1,
			}

			waterStatus := ""
			waterValue := data["water"]
			if waterValue < 5 {
				waterStatus = "Aman"
			} else if waterValue >= 5 && waterValue <= 8 {
				waterStatus = "Siaga"
			} else {
				waterStatus = "Bahaya"
			}

			windStatus := ""
			windValue := data["wind"]
			if windValue < 6 {
				windStatus = "Aman"
			} else if windValue >= 6 && windValue <= 15 {
				windStatus = "Siaga"
			} else {
				windStatus = "Bahaya"
			}

			payload, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshalling data:", err)
				continue
			}

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
			if err != nil {
				fmt.Println("Error creating request:", err)
				continue
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				continue
			}
			defer resp.Body.Close()

			fmt.Println("Response Status:", resp.Status)

			fmt.Printf("{ \"water\": %d,\n", waterValue)
			fmt.Printf("  \"wind\": %d\n}\n", windValue)
			fmt.Println("Status water:", waterStatus)
			fmt.Println("Status wind:", windStatus)

		case <-timeout.C:
			fmt.Println("Timeout. Stopping ticker...")
			ticker.Stop()
			return
		}
	}
}
