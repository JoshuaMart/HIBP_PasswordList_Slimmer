package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	hashFile, hashType, outputFile := flagCheck()

	// Starting the time counter
	startTimer := time.Now()

	// OPEN INPUT FILE
	file, _ := os.Open(hashFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// OPEN OUTPUT FILE
	output, _ := os.OpenFile(outputFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	defer file.Close()

	// HASH TYPE
	var hashLength int
	if hashType == "NTLM" {
		hashLength = 32
	} else if hashType == "SHA1" {
		hashLength = 40
	}

	for scanner.Scan() {
		hash := scanner.Text()[:hashLength]
		occurence := scanner.Text()[hashLength+1:]

		hash_byte, _ := hex.DecodeString(hash)
		output.Write(hash_byte)

		occurence_to_int, _ := strconv.Atoi(occurence)
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(occurence_to_int))
		output.Write(b)
	}

	log.Printf("It took : %s", time.Since(startTimer))
}

// Check if flag is provided
func flagCheck() (string, string, string) {
	flag.Parse()

	args := flag.Args()
	if len(args) < 3 {
		fmt.Println("Input and/or hash type and/or output file is missing.")
		os.Exit(1)
	}
	return args[0], args[1], args[2]
}
