package sv

import (
	"github.com/kataras/iris/v12"
	"reflect"
	"strings"
)

var (
	GlobalFailFunc = func(err error, ctx iris.Context) {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(iris.Map{"detail": err.Error()})
		return
	}
	GlobalContextKey = "sv"
)

func Run(valid interface{}, mode ...string) iris.Handler {
	var m string
	if len(mode) >= 1 {
		m = mode[0]
	}

	return func(ctx iris.Context) {
		// 回复到初始状态
		v := reflect.ValueOf(valid).Elem()
		v.Set(reflect.Zero(v.Type()))
		var err error
		switch m {
		case "query":
			err = ctx.ReadQuery(valid)
			break
		case "json":
			err = ctx.ReadJSON(valid)
			break
		case "xml":
			err = ctx.ReadXML(valid)
			break
		case "form":
			err = ctx.ReadForm(valid)
			break
		default:
			if ctx.Method() == "GET" {
				err = ctx.ReadQuery(valid)
			} else {
				contentType := ctx.GetContentTypeRequested()
				if contentType == "application/x-www-form-urlencoded" || strings.HasPrefix(contentType, "multipart/form-data") {
					err = ctx.ReadForm(valid)
				} else if contentType == "application/xml" {
					err = ctx.ReadXML(valid)
				} else {
					err = ctx.ReadJSON(valid)
				}
			}
			break
		}
		if err != nil {
			Warning.Printf("read valid data fail: %s", err.Error())
			GlobalFailFunc(err, ctx)
			return
		}

		if err := GlobalValidator.Check(valid); err != nil {
			Warning.Printf("valid fields fail: %s", err.Error())
			GlobalFailFunc(err, ctx)
			return
		}
		// this is point struct
		ctx.Values().Set(GlobalContextKey, valid)
		ctx.Next()
	}
}
