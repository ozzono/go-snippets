package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	taskResponse(1296675737, 0)
}

func taskResponse(taskID, retryCount int) (string, error) {
	log.Printf("Fetching response for taskID %d", taskID)
	if retryCount > 10 {
		return "", fmt.Errorf("Max retry count reached; Aborting proccess")
	}
	if retryCount > 1 {
		sleeper(15000 / 100)
	}
	log.Println("Fetching anti-captcha response")
	retryCount++
	type Solution struct {
		GRecaptchaResponse string `json:"gRecaptchaResponse"`
	}
	url := "https://api.anti-captcha.com/getTaskResult"

	payload := struct {
		ClientKey string `json:"clientKey"`
		TaskID    int    `json:"taskId"`
	}{
		ClientKey: "e6f3e3bf049bd40978b23ed144db6189", // key only used while testing
		TaskID:    taskID,
	}
	bytePayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("marshal err:\n - %v", err)
		sleeper(15000 / 100)
		return taskResponse(taskID, retryCount)
	}
	log.Printf("string(bytePayload): %v", string(bytePayload))
	req, err := http.NewRequest("POST", url, strings.NewReader(string(bytePayload)))
	if err != nil {
		log.Printf("http.NewRequest err:\n - %v", err)
		sleeper(15000 / 100)
		return taskResponse(taskID, retryCount)
	}

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("http.DefaultClient.do err:\n - %v", err)
		sleeper(15000 / 100)
		return taskResponse(taskID, retryCount)
	}

	acResponse := struct {
		ErrorID    int      `json:"errorId"`
		Status     string   `json:"status"`
		Solution   Solution `json:"solution"`
		Cost       string   `json:"cost"`
		IP         string   `json:"ip"`
		CreateTime int      `json:"createTime"`
		EndTime    int      `json:"endTime"`
		SolveCount int      `json:"solveCount"`
	}{}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll err:\n - %v", err)
		sleeper(15000 / 100)
		return taskResponse(taskID, retryCount)
	}
	err = json.Unmarshal(body, &acResponse)
	if err != nil {
		log.Printf("Unmarshal err:\n - %v", err)
		sleeper(15000 / 100)
		return taskResponse(taskID, retryCount)
	}
	if res.StatusCode == 200 {
		fmt.Printf("\n")
		log.Println("AntiCaptcha returned with success (200)")
		fmt.Printf("acResponse.ErrorID: %v\n", acResponse.ErrorID)
		fmt.Printf("acResponse.Status: %v\n", acResponse.Status)
		fmt.Printf("acResponse.Solution.GRecaptchaResponse: %v\n", acResponse.Solution.GRecaptchaResponse)
		fmt.Printf("acResponse.Cost: %v\n", acResponse.Cost)
		fmt.Printf("acResponse.IP: %v\n", acResponse.IP)
		fmt.Printf("acResponse.CreateTime: %v\n", acResponse.CreateTime)
		fmt.Printf("acResponse.EndTime: %v\n", acResponse.EndTime)
		fmt.Printf("acResponse.SolveCount: %v\n", acResponse.SolveCount)

		if acErr := acErrList(acResponse.ErrorID); acErr != nil {
			log.Printf("acErr[\"errID\"] ------------> %d", acErr["errID"])
			log.Printf("acErr[\"errorCode\"] --------> %s", acErr["errorCode"])
			log.Printf("acErr[\"errorDescription\"] -> %s", acErr["errorDescription"])
			fmt.Printf("\n")
			sleeper(15000 / 100)
			return taskResponse(taskID, retryCount)
		}
		if acResponse.Status == "ready" {
			return acResponse.Solution.GRecaptchaResponse, nil
		} else if acResponse.Status == "processing" {
			time.Sleep(time.Duration(30 * time.Second))
			return "", nil
		}
	}
	time.Sleep(time.Duration(30 * time.Second))
	log.Printf("Unhandled err at the end of taskResponse")
	sleeper(15000 / 100)
	return taskResponse(taskID, retryCount)
}

func sleeper(delay int) {
	if delay > 15 {
		log.Printf("Waiting a longer sleep: %vs", delay*100/1000)
	}
	for i := 0; i < delay; i++ {
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}

func acErrList(errID int) map[string]interface{} {
	switch errID {
	case 12:
		return map[string]interface{}{
			"errID":            errID,
			"errorCode":        "ERROR_CAPTCHA_UNSOLVABLE",
			"errorDescription": "Captcha could not be solved by 5 different workers"}
	case 16:
		return map[string]interface{}{
			"errID":            errID,
			"errorCode":        "ERROR_NO_SUCH_CAPCHA_ID",
			"errorDescription": "Task you are requesting does not exist in your current task list or has been expired. Tasks not found in active tasks"}
	default:
		return nil
	}
}
