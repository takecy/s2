## s2
[![Build Status](https://travis-ci.org/takecy/s2.svg?branch=master)](https://travis-ci.org/takecy/s2)  
![](https://img.shields.io/badge/golang-1.6.2-blue.svg?style=flat)
[![GoDoc](https://godoc.org/github.com/takecy/s2?status.svg)](https://godoc.org/github.com/takecy/s2)

s2 is convert to map from struct, and to struct from map for Go.  
I often seen `map[string]interface{}` type.   
ex) [mongo driver](https://godoc.org/labix.org/v2/mgo/bson#M).  
It is useful at the scene to be converted.

<br/>

## Usage
`go get github.com/takecy/s2`

[example](example/example.go)

### ToMap

```go
s := SomeStruct{ID:"id001"}
m, err := s2.ToMap("json", s)
```


### FromMap
```go
s := &SomeStruct{}
err := s2.FromMap(m, s)
```

<br/>

## Testing
```shell
$ go get github.com/smartystreets/goconvey/convey
$ go test
```
