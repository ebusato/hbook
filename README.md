hbook
=====

[![Build Status](https://secure.travis-ci.org/go-hep/hbook.png)](http://travis-ci.org/go-hep/hbook)
[![GoDoc](https://godoc.org/github.com/go-hep/hbook?status.svg)](https://godoc.org/github.com/go-hep/hbook)

`hbook` is a set of data analysis tools for HEP (histograms (1D, 2D, 3D), profiles and ntuples).

`hbook` is a work in progress of a concurrent friendly histogram filling toolkit.
It is loosely based on `AIDA` interfaces and concepts as well as the "simplicity" of `HBOOK` and the previous work of `YODA`.

## Installation

```sh
$ go get github.com/go-hep/hbook
```

## Documentation

Documentation is available on godoc:

 http://godoc.org/github.com/go-hep/hbook

## Example

### H1D

```go
package main

import (
	   "math/rand"
	   "github.com/go-hep/hbook"
)

func main() {
	 h := hbook.NewH1D(100, 0, 100)
	 for i := 0; i < 100; i++ {
	 	 h.Fill(rand.Float64()*100, 1.)
	 }
}

```

### NTuple

#### Open an existing n-tuple

```go
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-hep/csvutil/csvdriver"
	"github.com/go-hep/hbook"
)

func main() {
	db, err := sql.Open("csv", "data.csv")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	nt, err := hbook.OpenNTuple(db, "csv")
	if err != nil {
		panic(err)
	}

	h1, err := nt.ScanH1D("px where pt>100", nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("h1: %v\n", h1)

	h2 := hbook.NewH1D(100, -10, 10)
	h2, err = nt.ScanH1D("px where pt>100 && pt < 1000", h2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("h2: %v\n", h2)

	h11 := hbook.NewH1D(100, -10, 10)
	h22 := hbook.NewH1D(100, -10, 10)
	err = nt.Scan("px, py where pt>100", func(px, py float64) error {
		h11.Fill(px, 1)
		h22.Fill(py, 1)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("h11: %v\n", h11)
	fmt.Printf("h22: %v\n", h22)
}
```
