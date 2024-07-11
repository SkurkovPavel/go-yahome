# go-yahome
&emsp;&emsp;<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" alt="drawing" height="75"/>&emsp;&emsp;&emsp;<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/d/d2/Alisa_Yandex.svg/1200px-Alisa_Yandex.svg.png" alt="drawing" height="75"/>

![lib is under develope](https://img.shields.io/badge/status-in%20develope-yellowgreen)
![test](https://github.com/SkurkovPavel/go-yahome/actions/workflows/go-yahome.yml/badge.svg)
![Coverage Status](https://coveralls.io/repos/github/SkurkovPavel/go-yahome/badge.svg?branch=main)
[![Go Report](https://goreportcard.com/badge/github.com/SkurkovPavel/go-yahome)](https://goreportcard.com/report/github.com/SkurkovPavel/go-yahome)
![License](https://shields.io/badge/license-Mit-green.svg)

# About
This lib helps you to use the basic smart home api from Yandex. You can run a scenario or check the status of devices.


- Lib is under development. If you have any suggestions, please let me know or create an MR. Thanks
- Created for development of Yandex Alice skills using Golang

# Yandex IOT api client

Yandex api documentation you can find [HERE](<https://yandex.ru/dev/dialogs/smart-home/doc/reference/resources.html>)


## Quick Start

Add this import line to the file you're working in:

```Go
import "github.com/SkurkovPavel/go-yahome/iot"
```

Create the client as shown in the example
```Go
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
