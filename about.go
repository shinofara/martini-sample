package main

import (
  "martini/hoge"
	"github.com/codegangsta/martini-contrib/render"
)

type AboutViewRender struct {
	Title string
}

func AboutRender(r render.Render) {
	viewModel := AboutViewRender {
		hoge.Hoge(),
	}
	r.HTML(200, "about", viewModel)
}
