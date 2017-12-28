# eccrawler
crawler api to retrieve realtime data from crytocurrency exchange center

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

How to build development environment and how to do release

### Download [release](https://github.com/SpeculationFund/ECCrawler/releases)


### Unzip

```
unzip crawler-yyyymmdd.zip
```

It is shown as following,

```
Archive:  crawler-20171213.zip
  Length      Date    Time    Name
---------  ---------- -----   ----
        0  2017-12-13 09:29   .Build/
  2294464  2017-12-13 09:29   .Build/crawler_amd64-20171213
  1853064  2017-12-13 09:29   .Build/crawler_386-20171213
  1875552  2017-12-13 09:29   .Build/crawler_arm-20171213
---------                     -------
  6023080                     4 files

```

### Execute 

```
cd .Build
./crawler_armd64-20171213
```
                                                                                                                                                                                          

# Getting Started (TODO)

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

What things you need to install the software and how to install them

```
Give examples
```

### Installing

How to install package


### Running the tests

Explain how to run the automated tests for this system

```
Give an example
```
### Documentation
Document list of the project

### Build

How to build


# Logistics

### Contributing (TODO)

Please read [CONTRIBUTING.md](TODO) for details on our code of conduct, and the process for submitting pull requests to us.

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


