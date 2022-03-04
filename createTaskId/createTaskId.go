package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	id, err := createTaskID("6LflTU8UAAAAAMri7RQtZkELFzeglAoxSzRlGp_d")
	if err != nil {
		log.Printf("createTaskId err:\n%v", err)
	}
	fmt.Printf("taskid: %d\n", id)
}
func createTaskID(key string) (int, error) {
	log.Println("Creating anti-captcha task")
	log.Printf("Data-sitekey: %s", key)
	url := "https://api.anti-captcha.com/createTask"
	type task struct {
		TaskType   string `json:"type"`
		WebsiteURL string `json:"websiteURL"`
		WebsiteKey string `json:"websiteKey"`
	}
	type payload struct {
		ClientKey    string `json:"clientKey"`
		Task         task   `json:"task"`
		SoftID       int    `json:"softId"`
		LanguagePool string `json:"languagePool"`
	}
	var (
		jsonTask    task
		jsonPayload payload
	)

	jsonTask.TaskType = "NoCaptchaTaskProxyless"
	jsonTask.WebsiteURL = "http:\\/\\/minhaconta.riocardmais.com.br"
	jsonTask.WebsiteKey = key
	jsonPayload.Task = jsonTask
	jsonPayload.ClientKey = "cc23862b1385897e9151c342473bd3dd"
	jsonPayload.SoftID = 0
	jsonPayload.LanguagePool = "en"
	dataPayload, err := json.Marshal(jsonPayload)
	log.Println(string(dataPayload))
	if err != nil {
		return -1, fmt.Errorf("marshal err:\n%v", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(dataPayload)))
	if err != nil {
		return -1, fmt.Errorf("http.NewRequest err:\n%v", err)
	}

	req.Header.Add("expect", "")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, fmt.Errorf("http.DefaultClient.Do err:\n%v", err)
	}

	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("body response:\n%v", body)
			return -1, fmt.Errorf("ioutil.ReadAll err:\n%v", err)
		}
		var response struct { // anti-captcha api response
			ErrorID int `json:"errorId"`
			TaskID  int `json:"taskId"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Printf("Response body:\n%v", string(body))
			return -1, fmt.Errorf("Unmarshal err:\n%v", err)
		}
		return response.TaskID, nil
	}
	return -1, fmt.Errorf("Unhandled err creating task")
}
