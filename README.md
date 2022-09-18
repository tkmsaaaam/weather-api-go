# weather-api-go
This is the [weather.tsukumijima.net](https://weather.tsukumijima.net/) library for Go.
## Installation
```shell
go get github.com/tkmsaaaam/weather-api-go
```
## Usage
```go
package main

import (
	"fmt"
	"github.com/tkmsaaaam/weather-api-go"
)

const tokyo = "130010"

func main() {
	client := weather.New()
	body, err := client.Get(tokyo)
	if err != nil {
		fmt.Println("Error Request API")
	}
	fmt.Println(body.Title)
}
```
## License
MIT
## Author
tkmsaaaam
## Server
- https://weather.tsukumijima.net/
- https://github.com/tsukumijima/weather-api
