package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Too few arguments.")
	}

	args = args[1:]

	for _, fname := range args {
		checkIn(fname)
	}
}

func checkIn(fname string) {
	file, err := os.Open(fname)
	if err != nil {
		log.Printf("Warning: Cannot open '%s'.", fname)
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		eventStr := sc.Text()
		fmt.Println(eventStr)
	}
}

type anyEvent struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}
