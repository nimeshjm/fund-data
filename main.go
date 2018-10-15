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
	"github.com/gorilla/mux"
)

var pricesTemplate = `{{range $y, $x := .}}{{$x.Price}}
{{end}}`

var datesTemplate = `{{range $y, $x := .}}{{$x.Date}}
{{end}}`

type fund struct {
	Id string
}

func GetIdByISIN(isin string) string {
	url := fmt.Sprintf("https://lt.morningstar.com/api/rest.svc/9vehuxllxs/security_details/%s?viewId=investmentTypeLookup&idtype=isin&languageId=en-GB&currencyId=GBP&responseViewFormat=json", isin)

	httpClient := http.Client{
		Timeout: time.Second * 600,
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

	item, found := c.Get(url)
	if found {
		return item.(EtfSnapshot).LastPrice.Value, item.(EtfSnapshot).LastPrice.Currency.ID, item.(EtfSnapshot).LastPrice.Date
	} else {

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
			return 0, "", ""
		}

		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
			return 0, "", ""
		}

		perf := make([]EtfSnapshot, 0)
		jsonErr := json.Unmarshal(body, &perf)
		if jsonErr != nil {
			log.Fatal(jsonErr)
			return 0, "", ""
		}

		c.Set(url, perf[0], cache.DefaultExpiration)
		return perf[0].LastPrice.Value, perf[0].LastPrice.Currency.ID, perf[0].LastPrice.Date

	}
}

var c = cache.New(10*time.Minute, 10*time.Minute)

func main() {
	//getIds()
	router := mux.NewRouter()
	router.HandleFunc("/{id}/prices", handlerPrices).Methods("GET")
	router.HandleFunc("/{id}/dates", handlerDates).Methods("GET")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

type results struct {
	Price float64
	Date  string
}

func handlerPrices(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["id"]
	ids := getAccountIds(accountId)
	viewmodel := buildViewModel(ids)

	tmpl := template.New("pricesTemplate")
	tmpl, _ = tmpl.Parse(pricesTemplate)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	tmpl.Execute(w, viewmodel)
}

func handlerDates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["id"]
	ids := getAccountIds(accountId)
	viewmodel := buildViewModel(ids)

	tmpl := template.New("datesTemplate")
	tmpl, _ = tmpl.Parse(datesTemplate)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	tmpl.Execute(w, viewmodel)
}

func buildViewModel(ids []string) []results {
	var viewmodel []results
	item, found := c.Get("viewmodel")
	if found {
		return item.([]results)
	} else {
		for i := 0; i < len(ids); i++ {
			price, _, date := GetNavPrice(ids[i])

			if price == 0 {
				continue
			}

			fmt.Print(date)
			fmt.Println(price)

			viewmodel = append(viewmodel, results{price, date})
		}

		c.Set("viewmodel", viewmodel, cache.DefaultExpiration)
	}

	fmt.Println(viewmodel)
	return viewmodel
}

func getAccountIds(accountId string) []string {
	if accountId == "1" {
		return []string{
			"F00000O4Y5",
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
			"F00000OXIA",
			"F0GBR05BIW",
			"F00000NBK6",
			"F00000PPN1",
			"F00000OTTT",
			"F00000O7XF",
			"F00000MWKB",
			"F00000OWFJ",
			"F00000OPTP",
			"F0GBR0506N",
			"F00000OTU8",
			"F00000LK2Q",
			"F00000NRHT",
			"F00000OU8O",
			"F00000PLVX",
			"F00000NUVE",
			"F00000OIGY",
			"F00000OBRC",
			"F00000UMET",
			"F00000OPTZ",
			"F00000271D",
			"F00000OSWK",
			"F00000MJPU",
			"F0GBR050DM"}
	} else if accountId == "2" {
		return []string{"F00000O4Y5",
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
	}
	return nil
}
