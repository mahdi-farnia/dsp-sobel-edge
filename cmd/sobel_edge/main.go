/*
Apply Sobel Operator On Image For Edge Detection

usage:

	sobel_edge <img-path>.jpg <output-path>.jpg
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mahdi-farnia/dsp-sobel-edge/sobel"
)

const kFmtUsageString = "usage: %v <img-file>.jpg <output-path>.jpg"

func main() {
	if len(os.Args) != 3 {
		log.Fatalf(kFmtUsageString, os.Args[0])
	}

	//#region read input
	inFilename := os.Args[1]
	inFile, err := os.Open(inFilename)
	if err != nil {
		log.Fatalf("could not open file: %v", inFilename)
	}
	defer inFile.Close()
	//#endregion

	//#region apply sobel
	fmt.Println("> Applying sobel operator on image:", inFilename)

	writer, err := sobel.SobelOnJpeg(inFile)
	if err != nil {
		log.Fatal(err)
	}
	//#endregion

	//#region write output
	outFileName := os.Args[2]
	outFile, err := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatalf("could not open file: %v", outFileName)
	}
	defer outFile.Close()

	if err := writer(outFile); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully wrote to:", outFileName)
	//#endregion
}
