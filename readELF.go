package main

import (
	"fmt"
	"os"
)

func main() {
	parseELF("main", "main.out")
}

func parseELF(in string, out string) {
	// open input file
	fin, err := os.Open(in)
	defer fin.Close()
	if err != nil {
		fmt.Println(in, err)
		return
	}

	// create output file
	fout, err := os.Create(out)
	defer fout.Close()
	if err != nil {
		fmt.Println(out, err)
	}

	// read input file content
	buf := make([]byte, 16)
  // line number
  i := 0
	for {
    // line number
    i++
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		// write into outfile
    outputLine(fout, i, buf[:n])
	}
}

// output fomate one line
func outputLine(f *os.File, row int, data []byte){
		// line scope
    line := fmt.Sprintf("%04X:%04X", row, row*16)
		f.WriteString(line)

    // split
    f.WriteString(" | ")

		// 16 hex content
    length := len(data)
		content := fmt.Sprintf("%X", data[:length])
		f.WriteString(content)

    // split
    f.WriteString(" | ")

		// string content
    f.Write(data[:length])

    f.WriteString("\n")
}
