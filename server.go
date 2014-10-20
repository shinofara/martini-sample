package main

import (
	"flag"
	"github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func main() {
	
	m := martini.Classic()
	
	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout: "layout",
		Extensions: []string{".html", ".tmpl"},
		Charset: "utf-8",
	}))
	
	//redisから値をとってみるサンプル
	m.Get("/", IndexRender)
	m.Get("/about", AboutRender)
	m.Run()
}

