package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	setImei()
	// newImei := newImei()
	// fmt.Printf("imei: %v size: %v\n", newImei, len(newImei))
}

func setImei() {
	url := "https://35.188.98.162/api/v1/hardware/imei"
	user := "genymotion"
	pwd := "imF8ntMXB849HJp8"
	newImei := newImei()
	postData := map[string]string{"value": newImei}
	fmt.Printf("newimei: %v\n", postData["value"])
	// newImei := strings.SplitAfterN(strconv.FormatInt(rand.New(rand.NewSource(time.Now().UnixNano())).Int63(), 10), "", 5)[4]
	// json.Encoder(&postData)
	encodedData, err := json.Marshal(postData)
	if err != nil {
		fmt.Printf("marshal error: %v\n", err)
	}
	// fmt.Printf("encodedData: %v\n", encodedData)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(encodedData)))
	if err != nil {
		fmt.Printf("newrequest error: %v", err)
		return
	}
	req.SetBasicAuth(user, pwd)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	// resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("newrequest error: %v\n", err)
		return
	}
	if resp.StatusCode == 200 {
		fmt.Printf("Succesfully POST'ed new IMEI\n")
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("readbody error: %v\n", err)
			return
		}
		fmt.Printf("error body: %v\n", string(body))
	}

}
func newRandNumber() int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10)
}

func newImei() string {
	var imei []string
	var newImei int
	for i := 0; i < 13; i++ {
		newImei = newRandNumber()
		imei = append(imei, strconv.Itoa(newImei))
	}
	index14 := newRandNumber()
	index15 := newRandNumber()
	lastChar := lastChar(index14, index15)
	imei = append(imei, strconv.Itoa(index14))
	imei = append(imei, strconv.Itoa(index15))
	imei = append(imei, strconv.Itoa(lastChar))
	// fmt.Printf("newimei: %v size: %v\n", strings.Join(imei, ""), len(imei))
	return strings.Join(imei, "")
}

func lastChar(i13, i14 int) int {
	odd := i13 + i14
	even := 2 * i13 % 10
	return (odd + even) % 10
}
