# Qiita SDK for Go
<span style="display: inline-block;">
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/ktsujichan/qiita-sdk-go.svg?branch=master)](https://travis-ci.org/ktsujichan/qiita-sdk-go)
</span>

Qiita API v2 client library written in Golang.

## Install
```
go get -u github.com/ktsujichan/qiita-sdk-go
```

## Library
```golang
package main

import (
	"context"
	"github.com/ktsujichan/qiita-sdk-go/qiita"
)

func main() {
	config := qiita.NewConfig()
	c, _ := qiita.NewClient("<qiita access token>", *config)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c.ListItems(ctx, 1, 10, "Golang")
	c.ListUsers(ctx, 1, 10)
	c.GetUser(ctx, "r7kamura")
}
```
