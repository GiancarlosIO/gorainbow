[![Go Reference](https://pkg.go.dev/badge/github.com/giancarlosio/gorainbow.svg)](https://pkg.go.dev/github.com/giancarlosio/gorainbow)

`gorainbow` lets you use a rainbow version of a string ready to print it to the terminal.

![plot](./images/result.png)

## Install
```bash
go get github.com/giancarlosio/gorainbow
```

## Example
```go
  import "gorainbow"

  str := Rainbow("Hello world")

  fmt.Println(str)
```
