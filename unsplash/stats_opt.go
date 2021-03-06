// Copyright (c) 2017 Hardik Bagdi <hbagdi1@binghamton.edu>
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package unsplash

import (
	"bytes"
	"strconv"
)

// StatsOpt is used for custom statistics range
type StatsOpt struct {
	Resolution string `url:"resolution"`
	Quantity   int    `url:"quantity"`
}

var defaultStatsOpt = &StatsOpt{
	Resolution: "days",
	Quantity:   30,
}

// Valid validates the values in a StatsOpt
func (opt *StatsOpt) Valid() bool {
	if opt.Quantity < 0 || opt.Quantity > 30 {
		return false
	}
	if opt.Quantity == 0 {
		opt.Quantity = 30
	}
	if opt.Resolution == "" {
		opt.Resolution = "days"
	}
	if opt.Resolution != "days" {
		return false
	}
	return true
}

func (opt *StatsOpt) String() string {
	var buf bytes.Buffer
	if !opt.Valid() {
		buf.WriteString("StatsOpt is invalid")
	}
	buf.WriteString("StatsOpt: Resolution[")
	buf.WriteString(opt.Resolution)
	buf.WriteString("], Quantity[")
	buf.WriteString(strconv.Itoa(opt.Quantity))
	buf.WriteString("]")
	return buf.String()
}
