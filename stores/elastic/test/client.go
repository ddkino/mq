package main

import (
	"context"
	"fmt"

	"gopkg.in/olivere/elastic.v5"
)

type dbParams struct {
	serveur string
	port    string
}

var server = "127.0.0.1"
var port = "9200"

func main() {

	ctx := context.Background()
	fmt.Println("--- clientElastic ----")
	clientElastic, err := elastic.NewClient(elastic.SetURL("http://" + server + ":" + port))
	fmt.Println(clientElastic)
	fmt.Println(err)
	if err != nil {
		fmt.Println("connection: Error")
	}

	fmt.Println("connection: ok")

	mapping := `
	{
	"settings":{
		"number_of_shards":10,
		"number_of_replicas":3
	},
	"mappings":{
		"tweet":{
			"properties":{
				"tags":{
					"type":"text"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

	clientElastic.Index()
	createIndex, err := clientElastic.CreateIndex("twitter2").BodyString(mapping).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
	}

}
