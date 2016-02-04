# Overview
[![Twitter](https://img.shields.io/badge/author-%40MachielMolenaar-blue.svg)](https://twitter.com/MachielMolenaar)
[![GoDoc](https://godoc.org/github.com/Machiel/gorf?status.svg)](https://godoc.org/github.com/Machiel/gorf)

Frog is Natural Language Processing software developed for Dutch. Frog provides
a server, which can be accessed through TCP. This library tries to simplify
working and communicating with this server.

Gorf provides a simple interface to communicate with a running
[frog](http://languagemachines.github.io/frog/) server.

Not everything for the Token struct is implemented, contributions are welcome :).

# License
Gorf is licensed under a MIT license.

# Installation
A simple `go get github.com/Machiel/gorf` should suffice.

# Usage

## Example

```go
package main

import (
    "fmt"

    "github.com/Machiel/gorf"
)

func main() {

    client, err := gorf.NewClient("192.168.99.100:9919")

    if err != nil {
        panic(err)
    }

    sentences, err := client.Parse("Hoi, ik ben Machiel. Hebben we verbinding?")

    if err != nil {
        panic(err)
    }

    for i, sentence := range sentences {
        fmt.Printf("Sentence %d\n", i)

        for _, token := range sentence {
            fmt.Println(token.Lemma)
        }
    }
}
```
