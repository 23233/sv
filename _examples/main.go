package main

import (
	"github.com/23233/sv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
)

type req struct {
	Name string `json:"name" url:"name" comment:"name" validate:"required"`
}
type req2 struct {
	Desc string `json:"desc" url:"desc" comment:"desc"`
}

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	customLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		// Columns: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(customLogger)

	app.Any("/", sv.Run(new(req), "form"), func(ctx context.Context) {
		req := ctx.Values().Get("sv").(*req)
		_, _ = ctx.JSON(iris.Map{"name": req.Name})
	})
	app.Get("/111", sv.Run(new(req2)), func(ctx context.Context) {
		req := ctx.Values().Get("sv").(*req2)
		_, _ = ctx.JSON(iris.Map{"name": req.Desc})
	})

	_ = app.Listen(":8080")

}
