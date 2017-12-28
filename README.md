# eccrawler
Go package of crawler api to retrieve realtime data from crytocurrency exchange center

### The data sources

* [coinbase](https://developers.coinbase.com/api/v2)
* [bittrex](https://bittrex.com/Home/Api)
* [localbitcoin](https://localbitcoins.com/api-docs/)
* [cex](https://cex.io/cex-api)
* [kraken](https://www.kraken.com/en-us/help/api)
* [bitstamp](https://www.bitstamp.net/api/)
* [coincheck](https://coincheck.com/api_settings)


### The target data

* Realtime price for each crytocurrency
* Market depth, see [reference-1](http://www.futuresmag.com/2014/04/30/trading-market-depth), [reference-2](https://hackernoon.com/depth-chart-and-its-significance-in-trading-bdbfbbd23d33)
* Exchange fee
* Successful deal number
* Cumulative depth



# Deployment 
```
package main


import (

    "github.com/SpeculationFund/eccrawler"
    "fmt"
)



func main() {

    // Init crawler processor
    crawlersfactory := &crawlers.C_CrawlerFactory{}
    cpu := crawlersfactory.CreateCPU()

    result := make(chan interface{})
    cpu.GetTickers(result)
    for {
        fmt.Println(<-result)
    }
}
```                                                                                                                                                                                          

# Getting Started (TO BE TESTED)

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

Install golang  `sudo apt install golang-go`


### Installing

```
git clone https://github.com/SpeculationFund/eccrawler.git
cd eccrawler
go build
go install
```

### Running the tests

```
go test
```

### Documentation
* [GoDoc of accrawler](https://godoc.org/github.com/SpeculationFund/eccrawler)
* [Development docs of eccrawler](https://github.com/SpeculationFund/eccrawler/wiki)


### Build 

```
cd github.com/SpeculationFund/eccrawler
go build
```

# Logistics

### Contributing

Please read [CONTRIBUTING.md](https://github.com/SpeculationFund/eccrawler/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

### Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the tags on this repository

### Authors

* **PhoenixAndMachine** - *Initial work* - [PhoenixAndMachine](https://github.com/PhoenixAndMachine)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

### Acknowledgments (TODO)

* Hat tip to anyone who's code was used
* Inspiration
* etc


### License (open source only)

This project is licensed under the MIT License - see the [LICENSE.md](https://gist.github.com/Brownyuan/0b754b6009b7a4257bde9d1a23586678) file for details


