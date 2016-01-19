package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func RootRouter(m *web.Mux) {
	m.Get("/", Root)
}

func ContentRouter(m *web.Mux) {
	m.Get("/content/", ContentIndex)
	m.Get("/content/index", ContentIndex)
	m.Get("/content/new", ContentNew)
	m.Post("/content/new", ContentCreate)
	m.Get("/content/edit/:id", ContentEdit)
	m.Post("/content/update/:id", ContentUpdate)
	m.Get("/content/delete/:id", ContentDelete)
}

func main() {
	RootRouter(goji.DefaultMux)
	ContentRouter(goji.DefaultMux)
	goji.Serve()
}
