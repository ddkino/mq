package main

import (
	"bytes"
	"flag"
	"path/filepath"
	"text/template"

	"fmt"
	"log"
	"os"
	// "github.com/mitchellh/mapstructure"
)

func Factory(templateName string, data map[string]interface{}) (*bytes.Buffer, error) {
	/**
	------------- TEMPLATING --------------
	*/
	var templateFilename string
	switch templateName {
	case "topic":
		templateFilename = "topic.go.text"
	}
	fmt.Println("templateName=", templateName)
	fmt.Println("data=", data)
	filePrefix, _ := filepath.Abs("./templates/")
	t := template.New(templateName)
	t, err := template.ParseFiles(filePrefix + "/" + templateFilename)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return buf, err
	}
	return buf, nil
}

func main() {
	var buf bytes.Buffer
	MODEL := flag.String("model", "", "select your model")
	namespace := flag.String("namespace", "", "ex: topic.user")
	topic := flag.String("topic", "", "ex: login")
	description := flag.String("description", "", "description")
	version := flag.String("version", "0.1", "ex: 0.1 or release3")
	flag.Parse()
	flag.Args()
	log.Println(*MODEL)
	if *MODEL == "" {
		log.Println("model missing")
		os.Exit(3)
	}
	switch *MODEL {
	case "topic":
		if *namespace == "" || *topic == "" {
			log.Println("need namespace & topic")
			os.Exit(3)
		}
		resultBuffer, err := Factory(*MODEL, map[string]interface{}{
			"Signature": map[string]interface{}{
				"Namespace":   *namespace,
				"Topic":       *topic,
				"Description": *description,
				"Version":     *version,
			},
		})
		buf = *resultBuffer
		if err != nil {
			panic(err)
		}
	default:
		fmt.Println("you must choose a MODEL : topic, ....")
		os.Exit(3)
	}

	// ----------- WRITE BUF in file main.go 
	f, err := os.Create("./out/main.go")
	if err != nil {
		panic(err)
	}
	log.Print("template created")
	f.Write(buf.Bytes())
	defer f.Close()
}
