/*The fastprint package provides fast string concat functionality in Go.

With large quantity of []string or [][]byte concat, you may use strings.Join(array, "") or bytes.Join(array, []byte{}), here we provide something else with less memory allocation and faster operations.

Benchmark

	Here's the benchmark result on Intel(R) Core(TM) i5-3210M CPU @ 2.50GHz MacBook Mini with 16 GB 1600 MHz DDR3.
	$ go test -bench . -benchmem

	BenchmarkSprintS-4      100000000               12.9 ns/op             4 B/op          0 allocs/op
	BenchmarkFmtSprint-4    10000000               204 ns/op              57 B/op          1 allocs/op
	BenchmarkStringsJoin-4  50000000                27.3 ns/op             8 B/op          0 allocs/op
	BenchmarkBprintS-4      100000000               10.9 ns/op             4 B/op          0 allocs/op
	BenchmarkSprintB-4      100000000               13.6 ns/op             4 B/op          0 allocs/op
	BenchmarkFmtSprintB-4    2000000               795 ns/op             178 B/op          5 allocs/op
	BenchmarkBprintB-4      100000000               12.6 ns/op             4 B/op          0 allocs/op
	BenchmarkBytesJoin-4    100000000               18.9 ns/op             4 B/op          0 allocs/op


Installation

	$ go get github.com/zenixls2/fastprint

Doc Generation

	$ godocdown > README.md

*/
package fastprint

import (
	"unsafe"
)

/*
Concat string params to a single string
*/
func SprintS(params ...string) string {
	if len(params) == 0 {
		return ""
	}
	if len(params) == 1 {
		return params[0]
	}
	n := 0
	for _, s := range params {
		n += len(s)
	}
	b := make([]byte, n)
	index := 0
	for _, s := range params {
		index += copy(b[index:], s)
	}
	return *(*string)(unsafe.Pointer(&b))
}

func BprintS(params ...string) []byte {
	if len(params) == 0 {
		return []byte{}
	}
	if len(params) == 1 {
		return []byte(params[0])
	}
	n := 0
	for i := 0; i < len(params); i++ {
		n += len(params[i])
	}
	b := make([]byte, n)
	index := 0
	for _, s := range params {
		index += copy(b[index:], s)
	}
	return b
}

func SprintB(params ...[]byte) (s string) {
	if len(params) == 0 {
		return ""
	}
	if len(params) == 1 {
		return string(params[0])
	}
	n := 0
	for i := 0; i < len(params); i++ {
		n += len(params[i])
	}
	b := make([]byte, n)
	index := 0
	for _, s := range params {
		index += copy(b[index:], s)
	}
	// hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	// hdr.Data = (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data
	// hdr.Len = len(b)
	return *(*string)(unsafe.Pointer(&b))
}

func BprintB(params ...[]byte) []byte {
	if len(params) == 0 {
		return []byte{}
	}
	if len(params) == 1 {
		return []byte(params[0])
	}
	n := 0
	for i := 0; i < len(params); i++ {
		n += len(params[i])
	}
	b := make([]byte, n)
	index := 0
	for _, s := range params {
		index += copy(b[index:], s)
	}
	return b
}
