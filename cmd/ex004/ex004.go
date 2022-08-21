package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	INFILE  string = "INPUT.TXT"
	OUTFILE string = "OUTPUT.TXT"
)

// Split three digit number int three individual
// numbers
func splitDecNumber(num int) (int, int, int) {
	cnt := num / 100
	dec := (num - cnt*100) / 10
	eds := (num - cnt*100 - dec*10)

	return cnt, dec, eds
}

func main() {
	inf, err := os.Open(INFILE)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer inf.Close()

	rd := bufio.NewScanner(inf)

	rd.Scan()
	K, _ := strconv.Atoi(rd.Text())

	outf, err := os.Create(OUTFILE)
	defer outf.Close()

	outf.WriteString(fmt.Sprintf("%v%v%v", K, 9, 9-K))
}
