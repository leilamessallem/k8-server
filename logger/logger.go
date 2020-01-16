package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main()  {
	for t := range time.Tick(10*time.Minute) {
		dumpFileContent(t)
	}
}

func dumpFileContent(t time.Time) {
	fmt.Println(t)

	file, err := os.Open("/data/data")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
