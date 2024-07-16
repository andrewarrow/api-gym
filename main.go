package main

import (
	"context"
	"embed"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andrewarrow/api-gym/app"
	"github.com/andrewarrow/feedback/common"
	"github.com/andrewarrow/feedback/router"
	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v2"
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
	} else if arg == "parse" {
		filePath := "/Users/aa/pura/dc/openapi/openapi.yaml"
		file, _ := os.ReadFile(filePath)
		var doc openapi3.T
		yaml.Unmarshal(file, &doc)
		ctx := context.Background()
		doc.Validate(ctx)
		//doc.Validate(openapi3.NewLoader())
		fmt.Printf("OpenAPI Version: %s\n", doc.OpenAPI)
		fmt.Printf("Title: %s\n", doc.Info.Title)
		fmt.Printf("Description: %s\n", doc.Info.Description)
		fmt.Printf("Version: %s\n", doc.Info.Version)
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
		r.Paths["markup"] = app.Markup
		r.BucketPath = "/Users/aa/bucket"
		r.ListenAndServe(":" + os.Args[2])
	} else if arg == "help" {
	}

}
