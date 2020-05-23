package simple_valid

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	type req struct {
		Name string `json:"name" url:"name" comment:"name" validate:"required"`
	}
	sv := New("sv", func(err error, ctx context.Context) {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(iris.Map{"detail": err.Error()})
		return
	})
	app := iris.New()
	app.Get("/", sv.Run(new(req)), func(ctx context.Context) {
		req := ctx.Values().Get("sv").(*req)
		_, _ = ctx.JSON(iris.Map{"name": req.Name})
	})

	e := httptest.New(t, app)
	e.GET("/").Expect().Status(httptest.StatusBadRequest)

}
