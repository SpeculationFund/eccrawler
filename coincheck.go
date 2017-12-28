package crawlers


import (
    "time"
    "encoding/json"
    "strings"
//    "fmt"
)

/*
    Define Coincheck
*/
type Coincheck struct {
}

/*
    Coincheck ticker structure, according to https://coincheck.com/api_settings 
*/
type CoincheckTicker struct {
    High float64 `json:high`
    Low float64 `json:low`
    Last float64 `json:last`
    Timestamp float64 `json:timestamp`
    Bid float64 `json:bid`
    Ask float64 `json:ask`
    Volume float64 `json:volume`
}

/*
    Formalize data into Coincheck Ticker
*/
func (tc *Coincheck) FormTicker(raw interface{}, ip chan *CoincheckTicker) {
//    fmt.Println(string(raw.([]byte)))
    price := &CoincheckTicker{}
    json.Unmarshal(raw.([]byte), price)
    ip <- price
}

/*
    Get Coincheck Ticker by url
*/
func (tc *Coincheck) GetTickerByUrl(webapi I_WebAPI, ip chan *CoincheckTicker, url string) {
    webchan := make(chan interface{})

    go webapi.Get(url, webchan)
    go tc.FormTicker(<-webchan, ip)
}

/*
    Get Coincheck Ticker by exchange pair
*/
func (tc *Coincheck) GetTickerByPair(webapi I_WebAPI, cryto string, fiat string, result chan interface{}) {
    ticker := make(chan *CoincheckTicker)

    url := "https://coincheck.com/api/ticker?pair=" + strings.ToLower(cryto) + "_" + strings.ToLower(fiat)
    go tc.GetTickerByUrl(webapi, ticker, url)
    t := <-ticker
    r := &Ticker {
        Bid: t.Bid,
        Ask: t.Ask,
        Time: time.Now(),
        ExchangeCenter: tc.GetExchangeCenterName(),
        Cryto: cryto,
        Fiat: fiat,
        Volume: t.Volume,
    }
    result <- r
}

/*
    Get Coincheck Tickers of all exchange curriencies
*/
func (tc *Coincheck) GetTickers(webapi I_WebAPI, result chan interface{}) {
    crytoes := tc.GetCrytoes()
    fiats := tc.GetFiats()
    for _, cryto := range crytoes {
        for _, fiat := range fiats {
            go tc.GetTickerByPair(webapi, cryto, fiat, result)
        }
    }
}

/*
    Get exchange center name
*/
func (tc *Coincheck) GetExchangeCenterName() string {
    return "Coincheck"
}


/*
    Get crytocurrencies available in Coincheck
*/
func (tc *Coincheck) GetCrytoes() []string {
    return []string {
        "BTC",
        "ETH",
    }
}

/*
    Get fiatcurrencies available in Coincheck
*/
func (tc *Coincheck) GetFiats() []string {
    return []string {
        "JPY",
    }
}

