// Copyright 2015 The go-hep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hbook

type Annotation map[string]interface{}

// Histogram is an n-dim histogram (with weighted entries)
type Histogram interface {
	// Annotation returns the annotations attached to the
	// histogram. (e.g. name, title, ...)
	Annotation() Annotation

	// Name returns the name of this histogram
	Name() string

	// Rank returns the number of dimensions of this histogram.
	Rank() int

	// Axis returns the axis of this histogram.
	Axis() Axis

	// Entries returns the number of entries of this histogram.
	Entries() int64
}

// EOF
