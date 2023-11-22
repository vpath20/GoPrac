package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	data := "This is the sampe data going into the file"

	file, err := os.Create("./myFile.txt")

	checkNilError(err)

	len, err2 := io.WriteString(file, data)

	checkNilError(err2)

	fmt.Println(len)

	defer file.Close()
	readFile("./myFile.txt")

}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println(string(dataByte))

}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
