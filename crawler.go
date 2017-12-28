package crawlers



import (
)


/*
    Define crawler interface
*/
type I_MarketCrawler interface {
    // Get tickers of all possible exchange currency
    GetTickers(webapi I_WebAPI, result chan interface{})
    // Get Ticker of specified exchanged currencies
    GetTickerByPair(webapi I_WebAPI, cryto string, fiat string, result chan interface{})
    // Get Exchange center name
    GetExchangeCenterName() string
}


/*
type C_MarketCrawler struct {
}

func (tc *C_MarketCrawler) GetInsidePrice(webapi I_WebAPI, crytoes []string, fiats []string, result chan interface{}) {
    for _, cryto := range crytoes {
        for _, fiat := range fiats {
            go tc.GetInsidePriceByPair(webapi, cryto, fiat, result)
        }
    }
}

func (tc *C_MarketCrawler) GetInsidePriceByPair(webapi I_WebAPI, cryto string, fiat string, result chan interface{}) {
    result <- "No Implement"
}
*/
