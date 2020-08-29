// custom/ in-house implementation of go's tab-writer
// because ls shows table in column major fation and according to screen size
package ctw

import (
	"bytes"
	"fmt"
	"math"
)

func widthsSum(w []int, p int) int {
	s := 0
	for _, v := range w {
		s += v + p
	}
	s -= p

	return s
}

func findMax(arr []int) int {
	max := 0
	for _, v := range arr {
		if v >= max {
			max = v
		}
	}

	return max
}

//Ctw has complexity of no. of columns x no. of files
// termW is width of the terminal
func Ctw(buf *bytes.Buffer, data []string, termW int) {
	dn := len(data) // length of data slice
	if dn == 0 {
		return
	}
	const pad = 2            // padding b/w columns
	var lens []int           // slice of lengths of each element in the data slice
	widths := make([]int, 0) // slice of widths of each column (don't use because it over run once)
	var prevWidths []int     // accurate widths value
	for _, v := range data {
		lens = append(lens, len(v))
	}

	prevj := 0 // prevj is previous jump value (row value if you may)

	// for loop runs for all possible nos. of column (1,2,3,4....)
	for {
		cols := len(widths) + 1
		widths = append(widths, 0)
		j := int(math.Ceil(float64(dn) / float64(cols))) // jump value corresponding to cols
		if prevj == j {                                  // removes redundant calculations
			continue
		}
		b := 0 // begining of column
		e := j // end of column
		// find optimal widths (width of each column and total no of columns)
		for i := 0; i < cols && e <= dn; i++ {
			widths[i] = findMax(lens[b:e])
			b, e = e, e+j
		}

		// for last column if last column is not complete
		if e-j < dn {
			widths[cols-1] = findMax(lens[e-j : dn])
		}

		prevj = j

		totW := widthsSum(widths, pad) //total width of the ls block
		if totW > termW {
			break
		} else if totW > termW/2 { // if total width of the ls block is more than half of terminal
			// copy widths to prevWidths
			prevWidths = make([]int, len(widths))
			for i := range widths {
				prevWidths[i] = widths[i]
			}
		}
		if cols == dn { // if content comes in one line of terminal
			// copy widths to prevWidths
			prevWidths = make([]int, len(widths))
			for i := range widths {
				prevWidths[i] = widths[i]
			}
			break
		}

	}

	// total no of rows
	rows := int(math.Ceil(float64(dn) / float64(len(prevWidths))))

	// loop to write entire ls block to buffer
	for i := 0; i < rows; i++ {
		for j := 0; j < len(prevWidths); j++ {
			if i+j*rows >= dn { // checks for last column if incomplete
				continue
			}
			fmt.Fprintf(buf, "%-*s", prevWidths[j]+pad, data[i+j*rows])
		}
		fmt.Fprintf(buf, "\n")
	}
}
