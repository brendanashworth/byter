# byter
[![GoDoc](https://godoc.org/github.com/brendanashworth/byter?status.svg)](https://godoc.org/github.com/brendanashworth/byter)

> A simple CLI utility (with an internal API) for calculating the percentages of same-byte occurrences within a binary file. By default it will print all bytes with an occurrence rate of more than 1%.

### Usage
```bash
$ go build
$ ./byter <file>
```

Example output:
```
|--------------|
| 0x1  | 2%    |
|--------------|
| 0x65 | 1%    |
|--------------|
| 0x0  | 38%   |
|--------------|
| 0x2  | 2%    |
|--------------|
| 0x48 | 3%    |
|--------------|
| 0x24 | 2%    |
|--------------|
| 0x8B | 1%    |
|--------------|

Parsed 2091696 bytes.
```

### Example API
```go
package main

import (
	"github.com/brendanashworth/byter/api"
)

func main() {
	var data []byte
	// ..

	totalBytes, occurrenceMap := byter.CountOccurrences(data)
	// ..
}
```