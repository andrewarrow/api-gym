package generate

import (
	"api-gym/files"
	"api-gym/gym"
	"api-gym/util"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"text/template"
)

type Menu struct {
	Menu string
}

func Run(g *gym.Gym) {
	base := "../" + g.Name
	exec.Command("rm", "-rf", base).CombinedOutput()
	files.Mkdir(base)
	files.Mkdir(base + "/network")
	http := files.ReadFile("network/http.go")
	files.SaveFile(base+"/network/http.go", http)

	//files.SaveFile(base+"/main.go", files.ReadFile("generate/main.tmpl"))
	tmpl := template.New("mainMenu")
	tmpl.Parse(files.ReadFile("generate/main.tmpl"))
	mainMenu := Menu{}
	buff := []string{""}
	for _, s := range g.Structs {
		name := util.Plural(s.Name)
		buff = append(buff, fmt.Sprintf(`} else if command == "%s" {`, name))
		buff = append(buff, fmt.Sprintf(`var %s network.Homes`, name))
		buff = append(buff, fmt.Sprintf(`jsonString := network.List%s()`, util.ToCamelCase(name)))
		buff = append(buff, fmt.Sprintf(`json.Unmarshal([]byte(jsonString), &%s)`, name))
		buff = append(buff, fmt.Sprintf(`fmt.Println(%s)`, name))
	}
	mainMenu.Menu = strings.Join(buff, "\n")
	b := new(bytes.Buffer)
	tmpl.Execute(b, mainMenu)
	files.SaveFile(base+"/main.go", b.String())
	files.SaveFile(base+"/go.mod", files.ReadFile("generate/go.tmpl"))

	buff = []string{"package network"}
	for _, s := range g.Structs {
		name := util.ToCamelCase(s.Name)
		fmt.Println(name)
		buff = append(buff, fmt.Sprintf(`type %s struct {`, name))
		for _, f := range s.Fields {
			fname := util.ToCamelCase(f.Name)
			buff = append(buff, fmt.Sprintf("%s %s `json:\"%s\"`", fname, f.DataType(), f.Name))
		}
		buff = append(buff, fmt.Sprintf(`}`))
		buff = append(buff, fmt.Sprintf(`type %s struct {`, util.Plural(name)))
		buff = append(buff, fmt.Sprintf("%s []%s `json:\"%s\"`", util.Plural(name), name, util.Plural(s.Name)))
		buff = append(buff, fmt.Sprintf(`}`))
	}

	files.SaveFile(base+"/network/structs.go", strings.Join(buff, "\n"))

	buff = []string{"package network"}
	//buff = append(buff, `import "fmt"`)
	for _, s := range g.Structs {
		name := util.ToCamelCase(s.Name)
		fmt.Println(name)
		buff = append(buff, fmt.Sprintf(`func List%s() string {`, util.Plural(name)))
		buff = append(buff, fmt.Sprintf(`  jsonString := DoGet("/api/v1/%s")`, util.Plural(s.Name)))
		buff = append(buff, fmt.Sprintf(`  return jsonString`))
		buff = append(buff, fmt.Sprintf(`}`))
	}
	files.SaveFile(base+"/network/calls.go", strings.Join(buff, "\n"))
}
