package main

import (
	"encoding/base64"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zenazn/goji/web"
	"html/template"
	"models"
	"net/http"
	"strconv"
	"strings"
)

const Password string = "content:content"

var tpl *template.Template
var db gorm.DB

type FormData struct {
	Content models.Content
	Mess    string
}

func ContentIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	Contents := []models.Content{}
	db.Find(&Contents)
	tpl = template.Must(template.ParseFiles("view/content/index.html"))
	tpl.Execute(w, Contents)
}

func ContentNew(c web.C, w http.ResponseWriter, r *http.Request) {
	tpl = template.Must(template.ParseFiles("view/content/new.html"))
	tpl.Execute(w, FormData{models.Content{}, ""})
}

func ContentCreate(c web.C, w http.ResponseWriter, r *http.Request) {
	Content := models.Content{Content: r.FormValue("Content")}
	db.Create(&Content)
	http.Redirect(w, r, "/content/index", 301)
}

func ContentEdit(c web.C, w http.ResponseWriter, r *http.Request) {
	Content := models.Content{}
	Content.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Find(&Content)
	tpl = template.Must(template.ParseFiles("view/content/edit.html"))
	tpl.Execute(w, FormData{Content, ""})
}

func ContentUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	Content := models.Content{}
	Content.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Find(&Content)
	Content.Content = r.FormValue("Content")
	db.Save(&Content)
	http.Redirect(w, r, "/content/index", 301)
}

func ContentDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	Content := models.Content{}
	Content.Id, _ = strconv.ParseInt(c.URLParams["id"], 10, 64)
	db.Delete(&Content)
	http.Redirect(w, r, "/content/index", 301)
}

func SuperSecure(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Basic ") {
			pleaseAuth(w)
			return
		}

		password, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil || string(password) != Password {
			pleaseAuth(w)
			return
		}

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func pleaseAuth(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Gritter"`)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Go away!\n"))
}

func init() {
	db, _ = gorm.Open("mysql", "root@/gorm?charset=utf8&parseTime=True")
}
