package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

//logTime function implementation series (3 stages)--shows the time of func run---
func runningtime() time.Time { //stage 1. begin
	return time.Now()
}
func track(startTime time.Time) {
	endTime := time.Now()
	log.Println(endTime.Sub(startTime)) //stage 2. end
}
func logTime() {
	defer track(runningtime()) //stage 3. combination

}

//-----------------------------------------------end of logTime series--------------

//getFile function -> transmitting data to the new file with os.Create
func getFile(path string) {

	//----------------------------Opening the file------------
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("f = %T\n", f)
	defer f.Close()

	//-----------------------------file size data--------------
	fi, err := f.Stat()
	if err != nil {
		// Could not obtain stat, handle error
	}
	fmt.Printf("The file is %d bytes long \n", fi.Size())
	//---------------------------------------------------------

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
	//----------------------------------------------------------

	//----------Writing data from Open file To Created file-----------
	row := 1
	s := bufio.NewScanner(f)
	for s.Scan() {
		nb, _ := fo.WriteString(fmt.Sprintf("%d:%s\n", row, s.Text()))
		_ = nb
		row++
	}
	//----------------------------------------------------------------

	defer rowsCount("out.txt")
	defer logTime()
}

//rowsCount function-------------------------------
func rowsCount(path string) {
	file, _ := os.Open(path)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("number of lines:", lineCount)
}
//--------------------------------------------------

func main() {
	path := "data/in.txt"
	getFile(path)

	fmt.Println("Main function complete!")
}

//Geekbrains/Go2/les1$ go run task1.go 
// The file is 5014 bytes long 
// 2021/05/23 21:37:36 328ns
// number of lines: 100
// Main function complete!