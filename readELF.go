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
	buf := make([]byte, 1024)
	for {
		n, _ := fin.Read(buf)
		if 0 == n {
			break
		}
		// write into outfile
		fout.Write(buf[0:n])
	}
}
