package main

import(
	"fmt"
	"os"
	"time"
)

func main(){
	start := time.Now()
	var s, sep string
	for i := 1 ; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
		fmt.Println("\nÍndice:", i)
		fmt.Println("Valor:", os.Args[i])
	}
	fmt.Println("\n", os.Args[0], s)
	secs := time.Since(start).Seconds()
	fmt.Printf("Tempo: %.5fs\n", secs)
}