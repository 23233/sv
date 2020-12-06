package main

import (
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp()
	e := httptest.New(t, app)
	//e.GET("/111").WithQuery("desc", "123123").Expect().Status(httptest.StatusOK)
	bodyMap := map[string]interface{}{"name": "测试"}
	e.POST("/").WithJSON(bodyMap).Expect().Status(httptest.StatusOK)
	e.POST("/").WithForm(bodyMap).Expect().Status(httptest.StatusOK)
	e.GET("/").WithQuery("name", "测试").Expect().Status(httptest.StatusOK)
}
