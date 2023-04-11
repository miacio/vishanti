package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miacio/varietas/web"
	"github.com/miacio/vishanti/routers"
)

func init() {
	start()
}

func main() {
	g := web.New(gin.Default())

	// set static folder
	g.Static("/js", "../page/js")
	g.Static("/images", "../page/images")

	// load html folders
	g.LoadHTMLFolders([]string{"../page"}, ".html")

	routers.Register(g)

	g.Use(web.Limit(64))
	g.Run(":8080")
}
