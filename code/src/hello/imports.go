package main

//分别导入多个包
import (
	"fmt"
	"os"
)
import "math"

func main() {

	fmt.Fprintf(os.Stdout, "%.3f", math.Sqrt(7))
}
