package main

import (
	"fmt"

	"github.com/thangpham7793/qcd/dirs"
)

func main() {
	fmt.Println("Quick CD!")
	j := dirs.Dirs{}
	//Should create an empty json file
	j.Init(dirs.FILEPATH)
	j.AddOne("qcd", ".")
	fmt.Println(j)
}
