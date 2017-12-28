package crawlers


import (
)



/*
    Define the crawler process interface
*/
type I_CrawlerProcessor interface {
    GetTickers(result chan interface{})
    GetTickersByPair(cryto string, fiat string, result chan interface{})
    AddCrawler(crawler I_MarketCrawler) (error, string)
    ECNumber() int
}

/*
    Define the crawler processor structure
        * WebAPI:   handle web request
        * Crawlers: store the exchange center crawlers
*/
type C_CrawlerProcessor struct {
    WebAPI I_WebAPI
    Crawlers map[string]I_MarketCrawler
}

/*
    Get Ticker of all the exchange centers
*/
func (cpu *C_CrawlerProcessor) GetTickers(result chan interface{}) {
    for _, crawler := range cpu.Crawlers {
        go crawler.GetTickers(cpu.WebAPI, result)
    }
}

/*
    Add exchange center crawler
*/
func (cpu *C_CrawlerProcessor) AddCrawler(crawler I_MarketCrawler) (error, string) {
    tc := crawler.GetExchangeCenterName()
    cpu.Crawlers[tc] = crawler
    return nil, tc
}

/*
    Get Ticker price by exchange currency pair
*/
func (cpu *C_CrawlerProcessor) GetTickersByPair(cryto string, fiat string, result chan interface{}) {
    for _, crawler := range cpu.Crawlers {
        go crawler.GetTickerByPair(cpu.WebAPI, cryto, fiat, result)
    }
}

/*
    Get num of exchange center
*/
func (cpu *C_CrawlerProcessor) ECNumber() int {
    return len(cpu.Crawlers)
}
