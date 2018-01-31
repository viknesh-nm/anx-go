anx-go
===========

anx-go is a package written in go for the REST api of the [anx](https://anxpro.com/) trading platform.


## Installation

```go
go get github.com/viknesh-nm/anx-go
```

## Usage 
```go
package main

import (
	"fmt"
	"log"

	"github.com/viknesh-nm/anx-go"
)

func main() {
	api := anx.New("KEY", "SECRET")

	ticker, err := api.GetTicker("BTCUSD")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ticker)
	fmt.Println(ticker.Result)
}
```
## Notes and Files

[ANX Documentation V2](http://docs.anxv2.apiary.io/)

[ANX Documentation V3](http://docs.anxv3.apiary.io/)
