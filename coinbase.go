package crawlers


import (
    "time"
    "encoding/json"
    "strconv"
)


type Coinbase struct {
}

/*
    Define Coinbase Ticker structure, according to https://developers.coinbase.com/api/v2#prices
*/
type CoinbaseTicker struct {
    Amount string `json:amount`
    Base string `json:base`
    Currency string `json:currency`
}

/*
    Formalize data into Coinbase Ticker 
*/
func (tc *Coinbase) FormTicker(raw interface{}, ip chan *CoinbaseTicker) {
    meta := make(map[string]interface{})
    json.Unmarshal((raw).([]byte), &meta)
    price := &CoinbaseTicker{}
    data, _ := json.Marshal(meta["data"])
    json.Unmarshal(data, price)
    ip <- price
}

/*
    Get Coinbase ticker by url 
*/
func (tc *Coinbase) GetTickerByUrl(webapi I_WebAPI, ip chan *CoinbaseTicker, url string) {
    webchan := make(chan interface{})

    go webapi.Get(url, webchan)
    go tc.FormTicker(<-webchan, ip)
}


/*
    Get Coinbase Ticker by exchange pair
*/
func (tc *Coinbase) GetTickerByPair(webapi I_WebAPI, cryto string, fiat string, result chan interface{}) {
    url := "https://api.coinbase.com/v2/prices/"
    buy := make(chan *CoinbaseTicker)
    sell := make(chan *CoinbaseTicker)
    pair := cryto+"-"+fiat
    go tc.GetTickerByUrl(webapi, buy,  url + pair +"/buy")
    go tc.GetTickerByUrl(webapi, sell, url + pair + "/sell")
    bid, ask := <-buy, <-sell
    bidprice, _ := strconv.ParseFloat(bid.Amount, 64)
    askprice, _ := strconv.ParseFloat(ask.Amount, 64)
    r := &Ticker {
        Bid: bidprice,
        Ask: askprice,
        Fiat: fiat,
        Cryto: cryto,
        Time: time.Now(),
        ExchangeCenter: tc.GetExchangeCenterName(),
    }
    result <- r
}

/*
    Get Coinbase Tickers of all exchange currencies
*/
func (tc *Coinbase) GetTickers(webapi I_WebAPI, result chan interface{}) {
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
func (tc *Coinbase) GetExchangeCenterName() string {
    return "Coinbase"
}

/*
    Get crytocurrencies available in Coinbase
*/
func (tc *Coinbase) GetCrytoes() []string {
    return []string {
        "BTC",
        "ETH",
    }
}


/*
    Get fiatcurrencies available in Coinbase
*/
func (tc *Coinbase) GetFiats() []string {
    return []string {
        "USD",
    }
}
