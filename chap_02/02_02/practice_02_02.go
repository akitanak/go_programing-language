package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			v, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "practice 2.2: %v\n", err)
				os.Exit(1)
			}
			convertAndPrint(v)
		}
	} else {
		for _, arg := range args {
			v, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "practice 2.2: %v\n", err)
				os.Exit(1)
			}
			convertAndPrint(v)
		}
	}
}

func convertAndPrint(v float64) {
	f := Feet(v)
	m := Meter(v)
	fmt.Printf("%s = %s, %s = %s\n", f, FtToM(f), m, MToFt(m))
}

type Feet float64
type Meter float64

const FeetToMeterRatio = 3.2808

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// フィートからメートルへ変換する
func FtToM(f Feet) Meter { return Meter(f / FeetToMeterRatio) }

// メートルからフィートへ変換する
func MToFt(m Meter) Feet { return Feet(m * FeetToMeterRatio) }
