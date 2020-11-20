package main

import (
	"fmt"

	"github.com/thangpham7793/qcd/app"
	"github.com/thangpham7793/qcd/dirs"
)

func main() {
	d := dirs.Dirs{}
	d.Init("dir.json")
	d.AddCurrent("cli")
	a := app.New(app.Cmds, &d)
	a.ExecList()
	fmt.Println("Did not stop!")
}
