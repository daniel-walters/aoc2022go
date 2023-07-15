package generator

import (
	"fmt"
	"os"
	"path"
	"text/template"
)

type data struct {
	Day string
}

func Generate(day int) {
	files := []string{
		"generator/template_main.tmpl",
		"generator/template_test.tmpl",
	}

	toGen := fmt.Sprintf("day%02d", day)

	err := os.Mkdir(toGen, 0777)
	if err != nil {
		panic(err)
	}

	data := data{toGen}

	mainFileName := fmt.Sprintf("%s/%s.go", toGen, toGen)
	testFileName := fmt.Sprintf("%s/%s_test.go", toGen, toGen)

	createFromTemplate(mainFileName, files[0], data)
	createFromTemplate(testFileName, files[1], data)

}

func createFromTemplate(fileName, tmplFile string, data data) {
	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	t := template.New(path.Base(tmplFile))
	t.ParseFiles(tmplFile)

	err = t.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
