package crawler

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

// createFile creates the file for product in tmp directory
func createFile(p Product) {
	// marshal product details
	tmpData, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	// get the file name
	fileName := strings.Split(strings.Split(p.ProductUrl, "?")[0], "itm/")[1]

	// write the file to tmp directory
	err = os.WriteFile(tmpDir+"/"+fileName+".json", tmpData, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
