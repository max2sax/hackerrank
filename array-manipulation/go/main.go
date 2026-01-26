package array_manipulation

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read input from stdin and create data
	// copied from hackerrank stub
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInut := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	nTemp, err := strconv.ParseInt(firstMultipleInut[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mtemp, err := strconv.ParseInt(firstMultipleInut[0], 10, 64)
	checkError(err)
	m := int32(mtemp)

	queries := make([][]int32, m)
	for r := range m {
		queriesRowTemp := strings.Split(strings.Trim(readLine(reader), " \t\r\n"), " ")
		if len(queriesRowTemp) != 3 {
			panic("Bad input")
		}

		queriesRow := make([]int32, 3)
		for i, qeuriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(qeuriesRowItem, 10, 64)
			checkError(err)
			queriesRow[i] = int32(queriesItemTemp)
		}
		queries[r] = queriesRow
	}
	result := arrayManipulation(n, queries)
	fmt.Fprintf(writer, "%d\n", result)
	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
