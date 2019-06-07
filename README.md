Go env
======
> GoLang package to load environment variables from `.env`.

[![Go Report Card][goreportcard-image]][goreportcard-image] [![codecov][scrutinizer-image]][scrutinizer-link] [![License][license-image]][license-link] [![Build Status][travis-image]][travis-link] [![codecov][codecov-image]][codecov-link] 
## Installation
Native
```bash
$ go get github.com/zhooravell/go-env
```
dep
```bash
$ dep ensure -add github.com/zhooravell/go-env
```
## Usage

```go
package main

import (
	"github.com/zhooravell/go-env"
	"log"
)

func init()  {
	if err := env.Load(); err != nil {
		log.Fatal(err)
	}
}

func main()  {
	s := env.GetVar("ENV_STRING", "")
	b, _ := env.GetBoolVar("ENV_BOOL", false)
	i, _ := env.GetIntVar("ENV_INT", 0)
	i8, _ := env.GetInt8Var("ENV_INT8", 0)
	i16, _ := env.GetInt16Var("ENV_INT16", 0)
	i32, _ := env.GetInt32Var("ENV_INT32", 0)
	i64, _ := env.GetInt32Var("ENV_INT64", 0)
	f32, _ := env.GetFloat32Var("ENV_FLOAT32", 0)
	f64, _ := env.GetFloat64Var("ENV_FLOAT64", 0)

	log.Println("string: ", s)
	log.Println("bool: ", b)
	log.Println("int: ", i)
	log.Println("int8: ", i8)
	log.Println("int16: ", i16)
	log.Println("int32: ", i32)
	log.Println("int64: ", i64)
	log.Println("float32: ", f32)
	log.Println("float64: ", f64)
}
```

## Source(s)

* [environment variable wiki](https://en.wikipedia.org/wiki/Environment_variable)
* [os](https://golang.org/pkg/os/)
* [strconv](https://golang.org/pkg/strconv/)
* [strings](https://golang.org/pkg/strings/)
* [regexp](https://golang.org/pkg/regexp/)

[license-link]: https://github.com/zhooravell/go-env/blob/master/LICENSE
[license-image]: https://img.shields.io/dub/l/vibe-d.svg

[travis-link]: https://travis-ci.com/zhooravell/go-env
[travis-image]: https://travis-ci.com/zhooravell/go-env.svg?branch=master

[codecov-link]: https://codecov.io/gh/zhooravell/go-env
[codecov-image]: https://codecov.io/gh/zhooravell/go-env/branch/master/graph/badge.svg

[scrutinizer-link]: https://scrutinizer-ci.com/g/zhooravell/go-env/?branch=master
[scrutinizer-image]: https://scrutinizer-ci.com/g/zhooravell/go-env/badges/quality-score.png?b=master

[goreportcard-link]: https://goreportcard.com/report/github.com/zhooravell/go-env
[goreportcard-image]: https://goreportcard.com/badge/github.com/zhooravell/go-env