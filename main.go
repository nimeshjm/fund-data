package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"os"
	"bufio"
	"strings"
)

type fund struct {
	Id string
}

func GetIdByISIN(isin string) string {
	url := fmt.Sprintf("https://lt.morningstar.com/api/rest.svc/9vehuxllxs/security_details/%s?viewId=investmentTypeLookup&idtype=isin&languageId=en-GB&currencyId=GBP&responseViewFormat=json", isin)

	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fund := make([]fund, 0)
	jsonErr := json.Unmarshal(body, &fund)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return fund[0].Id
}

func getIds() {

	file, err := os.Open("isin.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	f, err := os.OpenFile("ids.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()

	for scanner.Scan() {
		line := scanner.Text()
		id := GetIdByISIN(line)
		f.WriteString(fmt.Sprintf("%s,%s\r\n", line, id))
	}
}

func GetNavPrice(id string) (float64, string, string) {
	url := fmt.Sprintf("http://tools.morningstar.co.uk/api/rest.svc/9vehuxllxs/security_details/%s?viewId=ETFsnapshot&idtype=msid&responseViewFormat=json", id)

	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	perf := make([]EtfSnapshot, 0)
	jsonErr := json.Unmarshal(body, &perf)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return perf[0].LastPrice.Value, perf[0].LastPrice.Currency.ID, perf[0].LastPrice.Date
}

func main() {

	// read ids from file
	file, err := os.Open("ids.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		row := scanner.Text()
		//isin := strings.Split(row, ",")[0]
		id := strings.Split(row, ",")[1]

		price, _, _ := GetNavPrice(id)

		//fmt.Println(isin, price, currency, date, "\r\n")
		fmt.Println(price)
	}
}
