package main

import(
	    "github.com/codegangsta/martini-contrib/render"
)

type IndexViewRender struct {
	Title string
}

func IndexRender(r render.Render) {
	viewModel := IndexViewRender {
		"name",
	}
	r.HTML(200, "index", viewModel)
}
