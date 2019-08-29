package main

import (
	"fmt"
	maths "go-learn/16_Maths"
	"os"
	"time"
)

func main() {
	t := time.Now()
	err := maths.SVGWriter(os.Stdout, t)

	if err != nil {
		fmt.Print("❌ error while generatign SVG")
	}
}
