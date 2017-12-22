package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	jevent "github.com/IgaguriMK/ed-journal/event"
	"github.com/pkg/errors"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Too few arguments.")
	}

	args = args[1:]

	total, success := 0, 0
	table := make(map[string]*RawEvevtSaver)
	for _, fname := range args {
		dt, ds := checkIn(table, fname)
		total += dt
		success += ds
	}

	for _, s := range table {
		s.closeFile()
	}

	fmt.Println("-----------------")
	fmt.Printf(" Total:   %6d\n", total)
	fmt.Printf(" Success: %6d\n", success)
	fmt.Printf(" Percent: %s\n", lFill(100.0*float64(success)/float64(total), 6))
	fmt.Println("-----------------")
}

func lFill(v float64, n int) string {
	s := fmt.Sprintf("%3.2f", v)
	for len(s) < n {
		s = " " + s
	}
	return s
}

func checkIn(table map[string]*RawEvevtSaver, fname string) (int, int) {
	file, err := os.Open(fname)
	if err != nil {
		log.Printf("Warning: Cannot open '%s'.", fname)
		return 0, 0
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	dt, ds := 0, 0
	for sc.Scan() {
		dt++
		eventStr := sc.Text()
		event, err := jevent.ParseEvent(eventStr)
		if err != nil {
			switch err := errors.Cause(err).(type) {
			case *jevent.UnknownEventType:
				saveFailRecord("0.unknown."+err.Type+".", ".json", err.Raw)
			default:
				log.Println(err)
				log.Fatal("    " + eventStr)
			}
			continue
		}

		getOrNewSaver(table, event).save(eventStr)

		if !checkLackOfField(eventStr, event) {
			continue
		}

		ds++
	}

	return dt, ds
}

func checkLackOfField(eventStr string, event jevent.Event) bool {
	bytes := []byte(eventStr)
	var v interface{}
	err := json.Unmarshal(bytes, &v)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err = json.MarshalIndent(v, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	eventStr = string(bytes)

	bytes, err = json.MarshalIndent(v, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	actualStr := string(bytes)

	wantLines := strings.Split(eventStr, "\n")
	getLines := strings.Split(actualStr, "\n")

	if len(wantLines) != len(getLines) {
		fmt.Println(eventStr)
		fmt.Println(actualStr)
		fmt.Println()
		saveFailRecord("1.mismatch."+event.GetEvent()+".", ".txt", eventStr+"\n"+actualStr)
		return false
	}

	return true
}

func saveFailRecord(prefix, suffix, content string) {
	hash := sha256.Sum256([]byte(content))
	hashStr := hex.EncodeToString(hash[:])

	name := prefix + hashStr[:10] + suffix
	file, err := os.Create("out_jcheck/" + name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Fprint(file, content)
}

type RawEvevtSaver struct {
	file    *os.File
	isFirst bool
}

func getOrNewSaver(table map[string]*RawEvevtSaver, event jevent.Event) *RawEvevtSaver {
	eventType := event.GetEvent()

	s, ok := table[eventType]
	if ok {
		return s
	}

	file, err := os.OpenFile("out_jcheck/types/"+eventType+".json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(file, "[")

	saver := &RawEvevtSaver{file: file, isFirst: true}
	table[eventType] = saver
	return saver
}

func (s *RawEvevtSaver) save(raw string) {
	if s.isFirst {
		s.isFirst = false
	} else {
		fmt.Fprint(s.file, ",")
	}

	fmt.Fprintf(s.file, "\n    %s", raw)
}

func (s *RawEvevtSaver) closeFile() {
	fmt.Fprintln(s.file, "\n]")
	s.file.Close()
}
