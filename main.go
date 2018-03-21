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
	"html/template"
)

var page = `{{range $val := .}}{{$val}}
{{end}}`

type fund struct {
	Id string
}

func GetIdByISIN(isin string) string {
	url := fmt.Sprintf("https://lt.morningstar.com/api/rest.svc/9vehuxllxs/security_details/%s?viewId=investmentTypeLookup&idtype=isin&languageId=en-GB&currencyId=GBP&responseViewFormat=json", isin)

	httpClient := http.Client{
		Timeout: time.Second * 60,
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
		Timeout: time.Second * 30,
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
	http.HandleFunc("/", handler,)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func init(){
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func handler(w http.ResponseWriter, r *http.Request) {

	ids:= []string {"F00000O4Y5",
		"F00000PLW7",
		"F00000PLW9",
		"F00000P781",
		"F00000P7MI",
		"F0GBR0506U",
		"F000002NAB",
		"F0GBR04RMW",
		"F0GBR04RMU",
		"F00000OPX3",
		"F00000PVLK",
		"F00000OWM6",
		"F00000P0QE",
		"F00000MZDQ",
		"F00000MWJQ",
		"F00000J3S6",
		"F0GBR06I57",
		"F00000PLVU",
		"F00000PW2X",
		"F00000OWFG",
		"F00000OYEQ",
		"F00000OPVF",
		"F00000OPUT",
		"F00000OXIA"}
	var prices []float64

	for i := 0; i < len(ids); i++ {
		price, _, _ := GetNavPrice(ids[i])

		prices = append(prices, price)

		//fmt.Println(isin, price, currency, date, "\r\n")
		//fmt.Fprintf(w, "%f", price)
	}

	tmpl := template.New("page")
	tmpl, _ = tmpl.Parse(page)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	tmpl.Execute(w, prices)
}
