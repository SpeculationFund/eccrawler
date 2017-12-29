package crawlers

import (
    "time"
)


/*
    Define Ticker
*/
type Ticker struct {
    Bid float64
    Ask float64
    Fiat string
    Cryto string
    Time time.Time
    ExchangeCenter string
    Vwap float64
    Volume float64
}


