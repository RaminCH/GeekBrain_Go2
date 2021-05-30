package main

import (
	"os"
)


//getFile function 
func getFile(path string) {

	//----------------------------------Create output file-------
	fo, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

}


func main() {
	path := "go/src/Geekbrains/Go2/les1/"
	getFile(path)
}