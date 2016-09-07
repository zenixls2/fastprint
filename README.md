# fastprint
--
    import "github.com/zenixls2/fastprint"

The fastprint package provides fast string concat functionality in Go.

With large quantity of []string or [][]byte concat, you may use
strings.Join(array, "") or bytes.Join(array, []byte{}), here we provide
something else with less memory allocation and faster operations.

### Benchmark

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

### Installation

    $ go get github.com/zenixls2/fastprint

### Doc Generation

    $ godocdown > README.md

## Usage

#### func  BprintB

```go
func BprintB(params ...[]byte) []byte
```

#### func  BprintS

```go
func BprintS(params ...string) []byte
```

#### func  SprintB

```go
func SprintB(params ...[]byte) (s string)
```

#### func  SprintS

```go
func SprintS(params ...string) string
```
Concat string params to a single string
