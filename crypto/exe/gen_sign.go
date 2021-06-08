package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "\n%s [flags] method api_key args\n\n", filepath.Base(os.Args[0]))
		flag.Usage()
		os.Exit(1)
	}

	total := ""
	for _, arg := range flag.Args() {
		total += arg
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(timestamp)
}
