package generate

import (
	"api-gym/files"
	"api-gym/gym"
	"api-gym/util"
	"fmt"
	"os/exec"
	"strings"
)

func Run(g *gym.Gym) {
	base := "../" + g.Name
	exec.Command("rm", "-rf", base).CombinedOutput()
	files.Mkdir(base)
	files.Mkdir(base + "/network")
	http := files.ReadFile("network/http.go")
	files.SaveFile(base+"/network/http.go", http)

	files.SaveFile(base+"/main.go", files.ReadFile("generate/main.tmpl"))
	files.SaveFile(base+"/go.mod", files.ReadFile("generate/go.tmpl"))

	buff := []string{"package network"}
	for _, s := range g.Structs {
		name := util.ToCamelCase(s.Name)
		fmt.Println(name)
		buff = append(buff, fmt.Sprintf(`type %s struct {`, name))
		for _, f := range s.Fields {
			fname := util.ToCamelCase(f.Name)
			buff = append(buff, fmt.Sprintf("%s %s `json:\"%s\"`", fname, f.DataType(), f.Name))
		}
		buff = append(buff, fmt.Sprintf(`}`))
	}
	files.SaveFile(base+"/network/structs.go", strings.Join(buff, "\n"))
}
