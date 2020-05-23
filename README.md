# simple_valid
a simple paramater validator for iris , use middleware

## use see test 
```cassandraql
// Define sv
	sv := simple_valid.New("sv", func(err error, ctx context.Context) {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(iris.Map{"detail": err.Error()})
		return
	})
// Define req valid struct
	type req struct {
		Name string `json:"name" url:"name" comment:"name" validate:"required"`
	}
// Use middleware
	app.Get("/", sv.Run(new(req)), func(ctx context.Context) {
		req := ctx.Values().Get("sv").(*req) // <- this get req data 
		_, _ = ctx.JSON(iris.Map{"name": req.Name})
	})
```
