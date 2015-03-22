# dotenv [![Build Status](https://travis-ci.org/lazureykis/dotenv.svg?branch=master)](https://travis-ci.org/lazureykis/dotenv)

Port of [bkeepers/dotenv](https://github.com/bkeepers/dotenv) to Go language.

Loads environment variables from .env file. Useful for development configuration.

Get started:
```bash
go get github.com/lazureykis/dotenv
cd $GOPATH/src/github.com/lazureykis/dotenv/example
go run main.go
```

Usage:
```go
import "github.com/lazureykis/dotenv"


func main() {
  dotenv.Go()
  // or dotenv.GoWithPath('/full/path/to/.env_file')
}
```

