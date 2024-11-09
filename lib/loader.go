package lib

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseFile(p string) [][]int {
	// output
	o := make([][]int, 0)

	// open code file
	f, err := os.Open(p)

	if err != nil {
		log.Fatal(err)
	}

	// defer file close
	defer f.Close()

	// read the file line by line
	s := bufio.NewScanner(f)

	for s.Scan() {
		// split line by spaces
		lv := strings.Split(s.Text(), " ")

		// parse elements
		ro := make([]int, 0)

		for _, q := range lv {
			iv, _ := strconv.ParseInt(q, 10, 32)

			// store
			ro = append(ro, int(iv))
		}

		// store row
		o = append(o, ro)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// verify all rows are the same length
	l := -1

	for _, r := range o {
		if l == -1 {
			// store length
			l = len(r)
		}

		if len(r) != l {
			log.Fatal("All rows must be the same length.")
		}
	}

	// verify square matrix
	if len(o) != l {
		log.Fatal("Program must be a square matrix.")
	}

	// return result
	return o
}
