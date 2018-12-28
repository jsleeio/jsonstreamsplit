// lifted and modified from https://blog.golang.org/json-and-go
// thanks!

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	outputPrefix := flag.String("output-prefix", "output", "base filename to which a sequence ID and .json are appended")
	flag.Parse()
	dec := json.NewDecoder(os.Stdin)
  for i := 0; true; i++ {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
    outputFilename := fmt.Sprintf("%s-%05d.json",*outputPrefix,i)
    outputFile,err := os.Create(outputFilename)
    if err != nil {
      log.Fatalf("can't open output file %q, giving up: %v", outputFilename, err)
    }
    enc := json.NewEncoder(outputFile)
    enc.SetIndent("","    ")
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
    outputFile.Close()
    log.Printf("wrote %q", outputFilename)
	}
}
