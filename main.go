package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
)

type Shape [][]bool

type Tetronimo struct {
	shape  Shape
	repr   string // letter representing the tetronimo
	rowLen int    //  rowLen and colLen to assertain if tetronimo is valid of size 4*4
	colLen int
}

type Field struct {
	Tetronimos []Tetronimo
	Dimension  int
}

var Tetris Field

func main() {
	file := os.Args[1]
	reader, err := fs.ReadFile(os.DirFS("./testExamples"), file)
	if err != nil {
		fmt.Println("error reading file ", err)
	}
	Tetris.fieldConstructor(reader)
	fmt.Println(Tetris.Tetronimos)
	fmt.Println()
	Tetris.printer()
}

func constructShape(t []byte) (Shape, int, int) {
	s := Shape{}
	var rows int
	var cols int
	for _, line := range bytes.Split(t, []byte(" ")) {
		if len(line) == 0 {
			continue
		}
		var row []bool
		for _, point := range line {
			if string(point) == "." {
				row = append(row, false)
				rows++
			} else if string(point) == "#" {
				row = append(row, true)
				rows++

			}
		}
		cols++
		s = append(s, row)
	}

	return s, rows / cols, cols
}
func constructTetronimos(t []byte) []Tetronimo {
	var tetronimos []Tetronimo
	var tetronimo Tetronimo
	newT := refiner(bytes.Split(t, []byte("\n")))
	for _, tet := range newT {
		tetronimo.shape, tetronimo.rowLen, tetronimo.colLen = constructShape(tet)
		tetronimo.repr = string('A' + len(tetronimos))
		tetronimos = append(tetronimos, tetronimo)
		tetronimo.shape, tetronimo.rowLen, tetronimo.colLen = nil, 0, 0
	}
	return tetronimos
}

func (f *Field) fieldConstructor(t []byte) {
	f.Tetronimos = constructTetronimos(t)
}

func refiner(t [][]byte) [][]byte {
	var tetronimosRefined [][]byte
	var tetronimo []byte
	for _, line := range t {
		if len(line) == 0 {
			if len(tetronimo) > 0 {
				tetronimosRefined = append(tetronimosRefined, tetronimo)
				tetronimo = nil
			}
			continue
		}
		tetronimo = append(tetronimo, line...)
		tetronimo = append(tetronimo, byte(' '))
	}
	return tetronimosRefined
}
func (f *Field) printer() {
	for _, tetronimo := range f.Tetronimos {
		for _, line := range tetronimo.shape {
			for _, point := range line {
				if !point {
					fmt.Printf(".")
				} else if point {
					fmt.Printf(tetronimo.repr)
				}
			}
			fmt.Println()

		}
		fmt.Println()
	}
}
