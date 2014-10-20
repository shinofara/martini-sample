package main

import(
	    "github.com/codegangsta/martini-contrib/render"
)

type AboutViewRender struct {
	Title string
}

func AboutRender(r render.Render) {
	viewModel := AboutViewRender {
		"name",
	}
	r.HTML(200, "about", viewModel)
}