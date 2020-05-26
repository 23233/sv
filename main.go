package sv

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var (
	GlobalFailFunc = func(err error, ctx context.Context) {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(iris.Map{"detail": err.Error()})
		return
	}
	GlobalContextKey = "sv"
)

// 请求验证器
type ReqValid struct {
	Valid interface{}
	Mode  string
}

func Run(valid interface{}, mode ...string) context.Handler {
	b := new(ReqValid)
	b.Valid = valid
	if len(mode) >= 1 {
		b.Mode = mode[0]
	}
	return b.Serve
}

func (c *ReqValid) Serve(ctx context.Context) {
	ctx.Values().Set(GlobalContextKey, "")
	var err error
	switch c.Mode {
	case "query":
		err = ctx.ReadQuery(&c.Valid)
		break
	case "json":
		err = ctx.ReadJSON(&c.Valid)
		break
	case "xml":
		err = ctx.ReadXML(&c.Valid)
		break
	case "form":
		err = ctx.ReadForm(&c.Valid)
		break
	default:
		if ctx.Method() == "GET" {
			err = ctx.ReadQuery(&c.Valid)
		} else {
			err = ctx.ReadJSON(&c.Valid)
		}
		break
	}
	if err != nil {
		Warning.Printf("read valid data fail: %s", err.Error())
		GlobalFailFunc(err, ctx)
		return
	}

	if err := GlobalValidator.Check(c.Valid); err != nil {
		Warning.Printf("valid fields fail: %s", err.Error())
		GlobalFailFunc(err, ctx)
		return
	}
	// this is point struct
	ctx.Values().Set(GlobalContextKey, c.Valid)
	ctx.Next()
}
