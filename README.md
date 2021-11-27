# go-yahome


![test](https://github.com/SkurkovPavel/go-yahome/actions/workflows/go-yahome.yml/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/SkurkovPavel/go-yahome/badge.svg?branch=main)](https://coveralls.io/github/SkurkovPavel/go-yahome?branch=main)
[![License](https://shields.io/badge/license-Apache-blue.svg)](http://copyfree.org)


Client for Yandex smart home api


## Installation

```bash
$ go get github.com/SkurkovPavel/go-yahome
```

## Quick Start

Add this import line to the file you're working in:

```Go
import "github.com/SkurkovPavel/go-yahome/iot"


func example(){
	
    config := iot.NewConfig()
    config.Token = "ACCESS TOKEN"

    cl := iot.NewIotClient(config)

    resp, err := cl.GetInfo()

    if err != nil {
        fmt.Print(err)
    }
	
    _ = resp
}
```