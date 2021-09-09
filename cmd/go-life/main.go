package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	life "github.com/svetlana-rezvaya/go-life"
)

func main() {
	outDelay :=
		flag.Duration("outDelay", 100*time.Millisecond, "delay between frames")
	flag.Parse()

	fieldBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("unable to read the field: ", err)
	}
	fieldBytes = bytes.TrimSpace(fieldBytes)

	field, err := life.ParseField(string(fieldBytes))
	if err != nil {
		log.Fatal("unable to unmarshal the field: ", err)
	}

	for {
		fmt.Print(field)
		time.Sleep(*outDelay)

		field = field.NextField()
	}
}
