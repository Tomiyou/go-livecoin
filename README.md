go-livecoin
==========

go-livecoin is an implementation of the livecoin API (public and private) in Golang.

This version implement V1.1 livecoin API and the new HMAC authentification.

## Import
	import "github.com/bitbandi/go-livecoin"
	
## Usage

In order to use the client with go's default http client settings you can do:

~~~ go
package main

import (
	"fmt"
	"github.com/bitbandi/go-livecoin"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	// livecoin client
	livecoin := livecoin.New(API_KEY, API_SECRET)

	// Get balances
	balances, err := livecoin.GetBalances()
	fmt.Println(err, balances)
}
~~~

In order to use custom settings for the http client do:

~~~ go
package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/bitbandi/go-livecoin"
)

const (
	API_KEY    = "YOUR_API_KEY"
	API_SECRET = "YOUR_API_SECRET"
)

func main() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	// livecoin client
	bc := livecoin.NewWithCustomHttpClient(conf.livecoin.ApiKey, conf.livecoin.ApiSecret, httpClient)

	// Get balances
	balances, err := livecoin.GetBalances()
	fmt.Println(err, balances)
}
~~~

See ["Examples" folder for more... examples](https://github.com/bitbandi/go-livecoin/blob/master/examples/livecoin.go)
