package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/kataras/iris"
)

func main() {

	redisInit()

	app := iris.New()

	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/healthcheck", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"status": "OK"})
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	type UriParams struct {
		Uri string
	}
	app.Post("/shorten", func(ctx iris.Context) {
		var u UriParams
		err := ctx.ReadJSON(&u)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": err})
			return
		}

		fmt.Println(u)

		code := GenerateCode(6)
		if dbCreate(code, u.Uri) {
			ctx.StatusCode(iris.StatusAccepted)
			ctx.JSON(iris.Map{"status": "accepted", "url": ctx.Host() + "/" + code})
		} else {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": "database write failed"})
		}

	})

	app.Get("/{code:string}", func(ctx iris.Context) {
		code := ctx.Params().Get("code")
		match, _ := regexp.MatchString("^[a-zA-Z0-9]{6}$", code)
		if match {
			if url := dbGet(code); len(url) > 0 {
				ctx.Redirect(url, iris.StatusMovedPermanently)
			} else {
				ctx.StatusCode(iris.StatusNotFound)
				ctx.JSON(iris.Map{"error": "not found"})
			}
		} else {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": "bad request"})
		}
	})

	app.Run(iris.Addr(":"+os.Getenv("PORT")), iris.WithoutServerError(iris.ErrServerClosed))
}
