// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package protoapi is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	api.proto
*/
package protoapi

import (
	"context"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"
)
import google_protobuf1 "google.golang.org/protobuf/types/known/emptypb"

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathNewsPing = "/news.service.v1.News/Ping"
var PathNewsGetNewsInfoById = "/news/news-info"
var PathNewsGetNews = "/news/all"

// NewsBMServer is the server API for News service.
type NewsBMServer interface {
	Ping(ctx context.Context, req *google_protobuf1.Empty) (resp *google_protobuf1.Empty, err error)

	GetNewsInfoById(ctx context.Context, req *IdReq) (resp *NewsInfoResp, err error)

	GetNews(ctx context.Context, req *google_protobuf1.Empty) (resp *NewsResp, err error)
}

var NewsSvc NewsBMServer

func newsPing(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := NewsSvc.Ping(c, p)
	c.JSON(resp, err)
}

func newsGetNewsInfoById(c *bm.Context) {
	p := new(IdReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := NewsSvc.GetNewsInfoById(c, p)
	c.JSON(resp, err)
}

func newsGetNews(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := NewsSvc.GetNews(c, p)
	c.JSON(resp, err)
}

// RegisterNewsBMServer Register the blademaster route
func RegisterNewsBMServer(e *bm.Engine, server NewsBMServer) {
	NewsSvc = server
	e.GET("/news.service.v1.News/Ping", newsPing)
	e.GET("/news/news-info", newsGetNewsInfoById)
	e.GET("/news/all", newsGetNews)
}
