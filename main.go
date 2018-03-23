package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var page = `{{range $y, $x := .}}{{$x.Price}} {{$x.Date}}
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

var c = cache.New(10*time.Minute, 10*time.Minute)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type results struct {
	Price float64
	Date  string
}

func handler(w http.ResponseWriter, r *http.Request) {

	ids := []string{"F00000O4Y5",
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
	var viewmodel []results

	item, found := c.Get("viewmodel")
	if found {
		viewmodel = item.([]results)
	} else {
		for i := 0; i < len(ids); i++ {
			price, _, date := GetNavPrice(ids[i])

			fmt.Print(date)
			fmt.Println(price)

			viewmodel = append(viewmodel, results{price, date})
		}

		c.Set("viewmodel", viewmodel, cache.DefaultExpiration)
	}

	fmt.Println(viewmodel)

	tmpl := template.New("page")
	tmpl, _ = tmpl.Parse(page)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	tmpl.Execute(w, viewmodel)
}
