package main

import (
	"text/template"
	"os"
)

type Content struct {
	MapContent     []string
	OutsideContent string
	Kind           bool
}

func handleTemplate(t string, v []Content) {

	// new template
	tmpl = template.Must(template.New("range").Parse(t))
    
   // execute template
    for _, c range := v {
       tmpl.Execute(os.Stdout, c)
	}

}

func main() {

	// declare template string
	rangeTemplate := `
{{if .Kind}}
{{range $i, $v := .MapContent}}
{{$i}} => {{$v}} , {{$.OutsideContent}}
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
