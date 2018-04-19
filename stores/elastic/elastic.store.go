package elastic

import (
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

func (c *Server) Connect() (*elastic.Client, error) {

	clientElastic, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:%d", c.Host, c.Port)))
	log.Printf("Try to connect to Elastic %s:%d", c.Host, c.Port)
	return clientElastic, err
}
