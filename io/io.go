package main

import "os"
import "fmt"
import "bufio"
import "strconv"
import "io"
import "flag"
import "strings"

func ioread(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open ", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line")
			return
		}

		str := string(line)
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			continue
		}

		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func iowrite(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("create file failed ", outfile, err)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

var infile = flag.String("i", "infile", "input file")
var outfile = flag.String("o", "outfile.dat", "output file")

func main() {
	flag.Parse()
	if infile != nil {
		values, err := ioread(*infile)
		if err != nil {
			fmt.Println("ioread failed, ", err)
			return
		}
		for i, v := range values {
			fmt.Println(i, v)
		}
		if outfile != nil {
			fmt.Println("writing...")
			err = iowrite(values, *outfile)
			if err != nil {
				fmt.Println("write failed ", err)
				return
			}
			fmt.Println("write done")
		}
	}
}
