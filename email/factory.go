package email

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

var (
	cwd          = os.Getenv("GOPATH")
	basePath     = "src/github.com/user/mq/email/templates"
	absolutePath = filepath.Join(cwd, basePath)
)

func (e *Email) FactoryEmail(templateName string, data map[string]string, ENV string) error {
	switch templateName {
	case "signup":
		templateData := struct {
			Name string
			URL  string
		}{
			Name: data["Name"],
			URL:  data["URL"],
		}
		var t *template.Template
		var err error

		if ENV == "development" || ENV == "" {
			templatePath := filepath.Join(absolutePath, "signup.html")
			t, err = template.ParseFiles(templatePath)
		} else {
			t, err = template.ParseFiles("/templates/signup.html")
		}
		if err != nil {
			return err
		}
		buf := new(bytes.Buffer)
		if err = t.Execute(buf, templateData); err != nil {
			return err
		}
		e.body = buf.String()
	default:
		panic("no template")
	}
	return nil
}
