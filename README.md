# IPInfo Library for Golang

Simple library to convert IP address to geolocation informations using API from http://ipinfo.io and calculate distance
between the coordinates points.


### Install

```bash
$ go get github.com/paulopinda/go-ipinfo
```

### Example

```golang
package main

import (
    "fmt"
    "github.com/paulopinda/go-ipinfo"
)

func main() {
    // My IP address.
    my := ipinfo.MyIP()
    fmt.Println(my.Hostname)

    // Other IP address.
    other := ipinfo.OtherIP("8.8.8.8")
    fmt.Println(other.Country)

    // Distance between two locals.
    distance := ipinfo.Distance(my, other)

    fmt.Println(distance)
}
```
