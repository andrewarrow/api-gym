package main

import (
	"embed"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/api-gym/app"
	"github.com/andrewarrow/feedback/common"
	"github.com/andrewarrow/feedback/router"
)

//go:embed app/feedback.json
var embeddedFile []byte

//go:embed views/*.html
var embeddedTemplates embed.FS

//go:embed assets/**/*.*
var embeddedAssets embed.FS

var buildTag string

func main() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		//PrintHelp()
		return
	}

	arg := os.Args[1]

	if arg == "import" {
	} else if arg == "email" {
	} else if arg == "render" {
		router.RenderMarkup()
	} else if arg == "run" {
		router.BuildTag = buildTag
		router.EmbeddedTemplates = embeddedTemplates
		router.EmbeddedAssets = embeddedAssets
		tf := common.TemplateFunctions()
		router.CustomFuncMap = &tf
		r := router.NewRouter("DATABASE_URL", embeddedFile)
		r.Paths["/"] = app.HandleWelcome
		r.Paths["gym"] = app.HandleGym
		r.BucketPath = "/Users/aa/bucket"
		r.ListenAndServe(":" + os.Args[2])
	} else if arg == "help" {
	}

}
