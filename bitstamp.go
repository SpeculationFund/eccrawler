package crawlers


import (
    "time"
    "encoding/json"
    "strconv"
)


type Bitstamp struct {
}

/*
    Bitstamp ticker structure, according to https://www.bitstamp.net/api/
*/
type BitstampTicker struct {
    High string `json:high`
    Last string `json:last`
    Timestamp string `json:timestamp`
    Bid string `json:bid`
    Ask string `json:ask`
    Open string `json:open`
    Vwap string `json:vwap`
    Volume string `json:volume`
}


/*
    Formalize data into Bitstamp Ticker
*/
func (tc *Bitstamp) FormTicker(raw interface{}, ip chan *BitstampTicker) {
    price := &BitstampTicker{}
    json.Unmarshal(raw.([]byte), price)
    ip <- price
}

/*
    Get Bitstamp Ticker by url
*/
func (tc *Bitstamp) GetTickerByUrl(webapi I_WebAPI, ip chan *BitstampTicker, url string) {
    webchan := make(chan interface{})

    go webapi.Get(url, webchan)
    go tc.FormTicker(<-webchan, ip)
}

/*
    Get Bitstamp Ticker by exchange pair
*/
func (tc *Bitstamp) GetTickerByPair(webapi I_WebAPI, cryto string, fiat string, result chan interface{}) {
    url := "https://www.bitstamp.net/api/v2/ticker/"
    ticker := make(chan *BitstampTicker)
    pair := cryto+fiat
    go tc.GetTickerByUrl(webapi, ticker, url + pair + "/")
    t := <-ticker
    bidprice, _ := strconv.ParseFloat(t.Bid, 64)
    askprice, _ := strconv.ParseFloat(t.Ask, 64)
    vwap, _ := strconv.ParseFloat(t.Vwap, 64)
    volume, _ := strconv.ParseFloat(t.Volume, 64)
    r := &Ticker {
        Bid: bidprice,
        Ask: askprice,
        Time: time.Now(),
        ExchangeCenter: tc.GetExchangeCenterName(),
        Cryto: cryto,
        Fiat: fiat,
        Vwap: vwap,
        Volume: volume,
    }
    result <- r
}

/*
    Get Bitstamp Tickers of all exchange currencies 
*/
func (tc *Bitstamp) GetTickers(webapi I_WebAPI, result chan interface{}) {
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
func (tc *Bitstamp) GetExchangeCenterName() string {
    return "Bitstamp"
}

/*
    Get crytocurrencies available in Bitstamp
*/
func (tc *Bitstamp) GetCrytoes() []string {
    return []string {
        "BTC",
        "ETH",
    }
}

/*
    Get fiatcurrencies available in Bitstamp
*/
func (tc *Bitstamp) GetFiats() []string {
    return []string {
        "USD",
    }
}
