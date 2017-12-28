package crawlers



import (
)

/*
    Define crawler factory interface
*/
type I_CrawlerFactory interface {
    CreateCPU() I_CrawlerProcessor
}


type C_CrawlerFactory struct {
}

/*
    Create CPU
        * Set exchange center crawlers
        * Set webapi
*/
func (factory *C_CrawlerFactory) CreateCPU() I_CrawlerProcessor {
    crawlers := make(map[string]I_MarketCrawler)
    webapi := &C_WebAPI{}
    cpu := &C_CrawlerProcessor{Crawlers: crawlers, WebAPI: webapi}
    cpu.AddCrawler(&Coinbase{})
    cpu.AddCrawler(&Bitstamp{})
    cpu.AddCrawler(&Coincheck{})
    return cpu
}

