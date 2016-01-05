package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func main() {
	content := web.New()
	goji.Handle("/content/*", content)
	content.Use(middleware.SubRouter)
	content.Use(SuperSecure)
	content.Get("/index", ContentIndex)
	content.Get("/new", ContentNew)
	content.Post("/new", ContentCreate)
	content.Get("/edit/:id", ContentEdit)
	content.Post("/update/:id", ContentUpdate)
	content.Get("/delete/:id", ContentDelete)
	goji.Serve()
}
