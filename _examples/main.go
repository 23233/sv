package main

import (
	"github.com/23233/sv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

type req struct {
	Name string `json:"name" form:"name" url:"name" xml:"name" comment:"name" validate:"required"`
}
type req2 struct {
	Desc string `json:"desc" url:"desc" comment:"desc"`
}

func NewApp() *iris.Application {
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Any("/", sv.Run(new(req)), func(ctx context.Context) {
		req := ctx.Values().Get("sv").(*req)
		_, _ = ctx.JSON(iris.Map{"name": req.Name})
	})
	app.Get("/111", sv.Run(new(req2)), func(ctx context.Context) {
		req := ctx.Values().Get("sv").(*req2)
		_, _ = ctx.JSON(iris.Map{"name": req.Desc})
	})
	return app
}

func main() {
	app := NewApp()
	_ = app.Listen(":8080", iris.WithOptimizations)

}
