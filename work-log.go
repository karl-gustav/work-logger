package main

import (
	"fmt"
	"os"
	"time"
)

const (
	tf = "2006-01-02 15:04:05"
	pa = "C:\\temp\\Work.log"
)

func main() {
	write("Started machine: " + time.Now().Format(tf))
	for i := 0; i < 3; i++ {
		ts := time.Now()
		time.Sleep(20 * time.Second)
		if time.Since(ts).Seconds() > 2 {
			write(fmt.Sprintf("off: %s\non:  %s\n", ts.Format(tf), time.Now().Format(tf)))
		} else {
			write("Less than 2 sec")
		}
	}
}

func write(text string) {
	f, err := os.OpenFile(pa, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(text); err != nil {
		panic(err)
	}
}
