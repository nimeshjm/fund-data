package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
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

func GetNavPrice(id string) (float64, string, string) {
	log.Print(id)
	if id == "0" {
		return 0, "", ""
	}

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

		if perf[0].LastPrice.Currency.ID == "GBX" {
			perf[0].LastPrice.Value = perf[0].LastPrice.Value / 100.0
		}

		c.Set(url, perf[0], cache.DefaultExpiration)
		return perf[0].LastPrice.Value, perf[0].LastPrice.Currency.ID, perf[0].LastPrice.Date

	}
}

var c = cache.New(60*time.Minute, 60*time.Minute)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{id}/prices", handlerPrices).Methods("GET")
	router.HandleFunc("/{id}/dates", handlerDates).Methods("GET")
	router.HandleFunc("/{id}/pricessold", handlerPricesSold).Methods("GET")
	router.HandleFunc("/{id}/datessold", handlerDatesSold).Methods("GET")
	log.Fatal(http.ListenAndServe("0.0.0.0:80", router))
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
	viewmodel := buildViewModel(accountId, ids)

	tmpl := template.New("pricesTemplate")
	tmpl, _ = tmpl.Parse(pricesTemplate)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	tmpl.Execute(w, viewmodel)
}

func handlerPricesSold(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["id"]
	ids := getSoldAccountIds(accountId)
	viewmodel := buildViewModel(accountId, ids)

	tmpl := template.New("pricesTemplate")
	tmpl, _ = tmpl.Parse(pricesTemplate)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	tmpl.Execute(w, viewmodel)
}

func handlerDates(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["id"]
	ids := getAccountIds(accountId)
	viewmodel := buildViewModel(accountId, ids)

	tmpl := template.New("datesTemplate")
	tmpl, _ = tmpl.Parse(datesTemplate)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	tmpl.Execute(w, viewmodel)
}

func handlerDatesSold(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["id"]
	ids := getSoldAccountIds(accountId)
	viewmodel := buildViewModel(accountId, ids)

	tmpl := template.New("datesTemplate")
	tmpl, _ = tmpl.Parse(datesTemplate)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	tmpl.Execute(w, viewmodel)
}

func buildViewModel(accountID string, ids []string) []results {
	var viewmodel []results
	cacheKey := accountID + ids[0]
	item, found := c.Get(cacheKey)
	if found {
		return item.([]results)
	} else {
		for i := 0; i < len(ids); i++ {
			price, _, date := GetNavPrice(ids[i])

			fmt.Print(date)
			fmt.Println(price)

			viewmodel = append(viewmodel, results{price, date})
		}

		c.Set(cacheKey, viewmodel, cache.DefaultExpiration)
	}

	fmt.Println(viewmodel)
	return viewmodel
}

func getSoldAccountIds(accountId string) []string {
	if accountId == "1" {
		return []string{
			"F00000PLW9",
			"F000002NAB",
			"F0GBR04RMW",
			"F0GBR04RMU",
			"F00000MZDQ",
			"F00000PLVU",
			"F00000OWFG",
			"F00000OPUT",
			"F00000NBK6",
			"F00000MJPU",
		}
	} else if accountId == "3" {
		return []string{
			"F00000O4Y5",
			"F00000PLW7",
			"F00000P781",
			"F00000P7MI",
			"F0GBR0506U",
			"F00000OPX3",
			"F00000PVLK",
			"F00000OWM6",
			"F00000P0QE",
			"F00000MWJQ",
			"F00000J3S6",
			"F0GBR06I57",
			"F00000PW2X",
			"F00000OYEQ",
			"F00000OPVF",
			"F00000OXIA",
			"F0GBR05BIW",
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
			"F0GBR050DM",
			"F00000VZBX",
			"F00000O7XV",
			"F0GBR04H80",
			"F00000OTU3",
			"F000003WHT",
			"F00000PBNF",
			"F00000OWFL",
			"F0GBR06IAV",
			"F00000PW4I",
			"F00000NQ9X",
			"F000003W5M",
			"F00000MEXK",
			"F00000OTU0",
			"F000002Y9G",
			"F00000P8DN",
			"F00000OXGE",
			"F00000OPUR",
			"F00000MWJV",
		}
	}
	return nil
}

func getAccountIds(accountId string) []string {
	if accountId == "1" {
		return []string{
			"F00000O4Y5",
			"F00000PLW7",
			"F00000P781",
			"F00000P7MI",
			"F0GBR0506U",
			"F00000OPX3",
			"F00000PVLK",
			"F00000OWM6",
			"F00000P0QE",
			"F00000MWJQ",
			"F00000J3S6",
			"F0GBR06I57",
			"F00000PW2X",
			"F00000OYEQ",
			"F00000OPVF",
			"F00000OXIA",
			"F0GBR05BIW",
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
			"F0GBR050DM",
			"F00000VZBX",
			"F00000O7XV",
			"F0GBR04H80",
			"F00000OTU3",
			"F000003WHT",
			"F00000PBNF",
			"F00000OWFL",
			"F0GBR06IAV",
			"F00000PW4I",
			"F00000NQ9X",
			"F000003W5M",
			"F00000MEXK",
			"F00000OTU0",
			"F000002Y9G",
			"F00000P8DN",
			"F00000OXGE",
			"F00000OPUR",
			"0", //"F0GBR04H80",
			"F00000MWJV",
			"0", 
			"0", 
			"0", 
			"0", 
			"F00000OU94",
			"F00000P77J",
			"F0GBR06IK3",
			"0", 
			"0", 
			"0", 
			"0", 
			"0", 
			"0", 
			"0", 
			"F00000OPUR",
			"F000000INI",
			"F00000PW4I",
			"F00000QF2T",
			"F0GBR06IAV",
			"0", 
			"0", 
			"0", 
			"0", 
			"0", 
			"0", 
			"F00000NUVE",
			"F00000OIGY",
			"F00000OBRC",
			"F00000LK2Q",
			"F00000VZBX",
			"F000001ZY2",
			"F0GBR05BIW",
			"F00000OWFJ",
			"F00000OTU8",
			"F00000O7XV",
			"F00000OPX8",
			"F00000OTTT",	
		}
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
	} else if accountId == "3" {
		return []string{"F00000OU94", "F00000P77J", "F0GBR06IK3", "F0GBR0506U", "F00000P7MI", "F0GBR06I57", "F0GBR0506N", "F00000NUVE", "F00000OIGY", "F00000OBRC", "F00000LK2Q", "F00000VZBX", "F00000OPUR", "F000001ZY2", "F000000INI", "F0GBR05BIW", "F00000OWFJ", "F00000OTU8", "F00000O7XV", "F00000PW4I", "F00000QF2T", "F0GBR06IAV", "F00000OPX8"}
	}
	return nil
}
