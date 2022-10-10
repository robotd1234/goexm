package main

import (
	"fmt"
	"os"
	"text/template"
)

type Content struct {
	MapContent     []string
	OutsideContent string
	Kind           bool
}

func handleTemplate(t string, v []Content) {

	// new template
	tmpl := template.Must(template.New("range").Parse(t))

	// execute template
	for _, c := range v {
		err := tmpl.Execute(os.Stdout, c)
		if err != nil {
			fmt.Printf("error is :%s\n", err)
		}
	}

}

func main() {

	// declare template string
	rangeTemplate := `
{{if .Kind}}
{{range $i, $v := .MapContent}}
{{.}}, {{$.OutsideContent}}
{{end}}
{{else}}
{{range .MapContent}}
{{.}} , {{$.OutsideContent}}
{{end}}
{{end}}
`
	// declare data
	str1 := []string{"第一次range", "使用了 index value"}
	str2 := []string{"第二次range", "未使用index value"}

	Contents := []Content{
		{str1, "outsidercontent1", true},
		{str2, "outsidercontent2", false},
	}

	handleTemplate(rangeTemplate, Contents)

}
