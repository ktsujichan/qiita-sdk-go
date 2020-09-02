你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Qiita SDK for Go

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/ktsujichan/qiita-sdk-go.svg?branch=master)](https://travis-ci.org/ktsujichan/qiita-sdk-go)
[![Code Climate](https://codeclimate.com/github/ktsujichan/qiita-sdk-go/badges/gpa.svg)](https://codeclimate.com/github/ktsujichan/qiita-sdk-go)
[![Issue Count](https://codeclimate.com/github/ktsujichan/qiita-sdk-go/badges/issue_count.svg)](https://codeclimate.com/github/ktsujichan/qiita-sdk-go)
[![Coverage Status](https://coveralls.io/repos/github/ktsujichan/qiita-sdk-go/badge.svg?branch=master)](https://coveralls.io/github/ktsujichan/qiita-sdk-go?branch=master)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/ktsujichan/qiita-sdk-go/issues)

Qiita API v2 client library written in Golang.

## Install
```
go get -u github.com/ktsujichan/qiita-sdk-go/qiita
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
