// ctw is a custom tab writter package build to correctly display icons and colors for color-ls
package ctw

import (
	"bytes"
	"fmt"
	"math"
)

type CTW struct {
	d        [][]string // entire data passed to ctw
	nw       []int      // widths of each fileName
	gw       []int      // width of each git column
	sw       []int      //width of each size column
	ic       []string   // color of file icon of a row
	cols     int        // zero based no of cols (or total cols -1)
	showIcon bool
	termW    int
}

/* each file comprises of 4 columns
|<---->|<---->|<------------------>|<--------->|
| size | icon | name+ext+indicator | gitStatus |
*/

func New(termW int) *CTW {
	t := new(CTW)
	t.cols = 3
	t.termW = termW

	t.nw = make([]int, 0)
	t.gw = make([]int, 0)
	t.sw = make([]int, 0)
	t.d = make([][]string, 0)
	t.ic = make([]string, 0)

	return t
}

func (w *CTW) AddRow(args ...string) {
	// length checking for args
	if len(args) != w.cols+1 {
		return
	}

	w.sw = append(w.sw, len(args[0]))
	w.nw = append(w.nw, len(args[2]))
	w.gw = append(w.gw, len(args[3]))

	if w.showIcon == false {
		w.showIcon = len(args[1]) > 0
	}

	w.d = append(w.d, args)
}

func (w *CTW) IconColor(c string) {
	w.ic = append(w.ic, c)
}

func (w *CTW) Flush(buf *bytes.Buffer) {
	dn := len(w.d)
	if dn == 0 {
		return
	}
	pad := 2

	iw := make([][4]int, 0) // slice of widths of each column (don't use because it over run once)
	var widths [][4]int

	prevj := 0 // prevj is previous jump value (row value if you may)

	for {
		cols := len(iw) + 1
		iw = append(iw, [4]int{0, 0, 0, 0})
		j := int(math.Ceil(float64(dn) / float64(cols))) // jump value corresponding to cols
		if prevj == j {                                  // removes redundant calculations
			continue
		}
		b := 0 // begining of column
		e := j // end of column
		// find optimal widths (width of each column and total no of columns)
		for i := 0; i < cols && e <= dn; i++ {
			iw[i] = w.colW(b, e)
			b, e = e, e+j
		}

		// for last column if last column is not complete
		if e-j < dn {
			iw[cols-1] = w.colW(e-j, dn)
		}

		prevj = j

		totW := widthsSum(iw, pad) //total width of the ls block
		if totW > w.termW {
			// not even first iteration done print similar to logo-ls -1
			if len(widths) == 0 {
				widths = make([][4]int, len(iw))
				for i := range iw {
					widths[i] = iw[i]
				}
			}
			break
		} else if totW >= w.termW/2 { // if total width of the ls block is more than half of terminal
			// copy iw to widths
			widths = make([][4]int, len(iw))
			for i := range iw {
				widths[i] = iw[i]
			}
		}

		if cols == dn { // if content comes in one line of terminal
			// copy widths to prevWidths
			widths = make([][4]int, len(iw))
			for i := range iw {
				widths[i] = iw[i]
			}
			break
		}
	}

	// total no of rows
	rows := int(math.Ceil(float64(dn) / float64(len(widths))))

	// loop to write entire ls block to buffer
	for i := 0; i < rows; i++ {
		p := pad
		for j := 0; j < len(widths); j++ {
			if i+j*rows >= dn { // checks for last column if incomplete
				continue
			}
			if j == len(widths)-1 {
				p = 0
			}
			w.printCell(buf, i+j*rows, widths[j])
			fmt.Fprintf(buf, "%*s", p, "")
		}
		fmt.Fprintf(buf, "\n")
	}

}

func (w *CTW) colW(b, e int) [4]int {
	s, n, g := 0, 0, 0 // max od size column, name column, gitStatus column
	for i := b; i < e; i++ {
		if w.sw[i] > s {
			s = w.sw[i]
		}
		if w.nw[i] > n {
			n = w.nw[i]
		}
		if w.gw[i] > g {
			g = w.gw[i]
		}
	}

	ans := [4]int{0, 0, 0, 0}
	if s > 0 {
		ans[0] = s + 1
	}
	if w.showIcon {
		ans[1] = 2
	}
	ans[2] = n
	if g > 0 {
		ans[3] = 2
	}
	return ans
}

func (w *CTW) printCell(buf *bytes.Buffer, i int, cs [4]int) {
	if cs[0] > 0 {
		fmt.Fprintf(buf, "%-*s%s", cs[0]-1, w.d[i][0], brailEmpty)
	}
	if w.showIcon {
		fmt.Fprintf(buf, "%s%1s%s%s", w.ic[i], w.d[i][1], noColor, brailEmpty)
	}
	fmt.Fprintf(buf, "%s%-*s%s", getGitColor(w.d[i][3]), cs[2], w.d[i][2], noColor)

	if cs[3] > 0 {
		fmt.Fprintf(buf, "%s%s%1s%s", brailEmpty, getGitColor(w.d[i][3]), w.d[i][3], noColor)
	}
}
