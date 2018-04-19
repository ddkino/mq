package do

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/user/mq/chassis"
)

func Flags(signature *chassis.Signature) {
	flag.StringVar(&signature.Env, "env", "development", "environnement")
	flag.StringVar(&signature.Topic, "topic", "", "=login, =byid etc ...")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		if args[0] == "h" || args[0] == "help" {
			fmt.Printf("Usage: run --topic=<topic> \n")
			fmt.Printf("\nsignature %+v\n", signature)
			os.Exit(3)
		}
	}
	if len(signature.Topic) == 0 {
		log.Fatalf("Usage: run --topic=<topic> \n")
		fmt.Println("topic missing")
		os.Exit(3)
	}
}
