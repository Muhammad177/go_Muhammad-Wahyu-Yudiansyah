// package main

// import (
//     "bufio"
//     "fmt"
//     "io"
//     "os"
//     "strings"
// )

// /*
//  * Complete the 'getAnagram' function below.
//  *
//  * The function is expected to return an INTEGER.
//  * The function accepts STRING s as parameter.
//  */

// func getAnagram(s string) int32 {
//  panjang := len(s)

//     if panjang % 2 != 0{
//         return -1
//     }

//     freq1 := make([]int32,10)

//     for i := 0; i < panjang/2; i++ {
//         freq1[int(s[i]) - '0']++
//     }

//     pergantian := 0
//     for i := panjang/2; i<panjang; i++ {
//         digit := int(s[i]) - '0'
//         freq1[digit]--
//         if freq1[digit]<0{
//             pergantian++
//         }
//     }
//     return int32(pergantian)
// }

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//     checkError(err)

//     defer stdout.Close()

//     writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

//     s := readLine(reader)

//     result := getAnagram(s)

//     fmt.Fprintf(writer, "%d\n", result)

//     writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'deleteProducts' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ids
 *  2. INTEGER m
 */

func deleteProducts(ids []int32, m int32) int32 {
	hitung := make(map[int32]int)
	for _, id := range ids {
		hitung[id]++
	}
	for i := 0; int32(i) < m; i++ {
		minFreq := len(ids) + 1
		minID := int32(0)
		for id, freq := range hitung {
			if freq > 0 && freq < minFreq {
				minFreq = freq
				minID = id
			}
		}
		if minFreq == len(ids)+1 {
			break
		}
		hitung[minID]--
	}

	count := 0
	for _, freq := range hitung {
		if freq > 0 {
			count++
		}
	}
	return int32(count)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	idsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var ids []int32

	for i := 0; i < int(idsCount); i++ {
		idsItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		idsItem := int32(idsItemTemp)
		ids = append(ids, idsItem)
	}

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := deleteProducts(ids, m)

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
