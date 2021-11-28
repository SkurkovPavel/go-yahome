# Yandex IOT api client

yandex api documentation you can find [HERE](<https://yandex.ru/dev/dialogs/smart-home/doc/reference/resources.html>)


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