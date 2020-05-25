package simple_valid

import (
	"github.com/kataras/iris/v12/context"
	"reflect"
)

// 请求验证器
type ReqValid struct {
	Valid      interface{}
	Mode       string
	FailFunc   func(err error, ctx context.Context)
	ContextKey string
}

func (c *ReqValid) Run(valid interface{}) context.Handler {
	b := ReqValid{
		Valid:      valid,
		Mode:       c.Mode,
		FailFunc:   c.FailFunc,
		ContextKey: c.ContextKey,
	}
	return b.Serve
}

func New(cKey string, fail func(err error, ctx context.Context), more ...string) ReqValid {
	var c ReqValid
	if len(more) >= 1 {
		c.Mode = more[0]
	}
	c.FailFunc = fail
	c.ContextKey = cKey
	return c
}

func (c *ReqValid) Serve(ctx context.Context) {
	ctx.Values().Set(c.ContextKey, "")
	valid := reflect.New(reflect.TypeOf(c.Valid).Elem()).Interface()
	var err error
	switch c.Mode {
	case "query":
		err = ctx.ReadQuery(&valid)
		break
	case "json":
		err = ctx.ReadJSON(&valid)
		break
	case "xml":
		err = ctx.ReadXML(&valid)
		break
	case "form":
		err = ctx.ReadForm(&valid)
		break
	default:
		if ctx.Method() == "GET" {
			err = ctx.ReadQuery(&valid)
		} else {
			err = ctx.ReadJSON(&valid)
		}
		break
	}
	if err != nil {
		Warning.Printf("read valid data fail %s", err.Error())
		c.FailFunc(err, ctx)
		return
	}

	if err := GlobalValidator.Check(valid); err != nil {
		Warning.Printf("valid fields fail %s", err.Error())
		c.FailFunc(err, ctx)
		return
	}
	// this is point struct
	ctx.Values().Set(c.ContextKey, valid)
	ctx.Next()
}
