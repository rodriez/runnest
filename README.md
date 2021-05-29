# runnest
A simple module for build BDD test cases

## Installation

```bash
go get github.com/rodriez/runnest
```

## Usage

```go
import "github.com/rodriez/runnest"

func Test(t *testing.T) {
    testCases := []runnest.TestCase{
        {
            Name: "Given ping When while service is active Then return pong",
            Given: func() interface{} {
                return true
            },
            When: func(req interface{}) (interface{}, error) {
                active := req.(bool)
                serv := &fakeService{active}

                return serv.ping(), nil
            },
            Then: func(t *testing.T, resp interface{}, e error) {
                if e != nil {
                    t.Errorf("Error: %s", e.Error())
                }

                if pong := resp.(string); pong != "pong" {
                    t.Errorf("Error: expected pong Received %s", pong)
                }
            },
        },
    }

    runnest.NewRunest(t).Run(testCases)
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

[Hey 👋 buy me a beer! ](https://www.buymeacoffee.com/rodriez)