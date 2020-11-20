package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/thangpham7793/qcd/dirs"
)

type strCmd struct {
	name  string
	value string
	desc  string
}

//App contains all commands and added methods
type App struct {
	add   string
	ls    string
	alias string
	d     dirs.Dirs
}

//New initialises app, implementing the command pattern, delegating actual work to dirs package
func New(cmds map[string]strCmd, d *dirs.Dirs) (NewApp *App) {
	a := App{}
	add := flag.String(cmds["add"].name, cmds["add"].value, cmds["add"].desc)
	ls := flag.String(cmds["ls"].name, cmds["ls"].value, cmds["ls"].desc)
	alias := flag.String(cmds["alias"].name, cmds["alias"].value, cmds["alias"].desc)
	flag.Parse()
	a.add = *add
	a.ls = *ls
	a.d = *d
	a.alias = *alias
	return &a
}

//PrintArgs print received args
func (a *App) PrintArgs() {
	fmt.Println((*a).add)
	fmt.Println((*a).ls)
}

//ExecList lists stored all aliases and paths
func (a *App) ExecList() {
	switch a.ls {
	case "all":
		a.d.List()
	default:
		a.d.ListOne(a.ls)
	}
	os.Exit(0)
}

//ExectAdd adds either the current or given path and alias
// func (a *App) ExectAdd() {
// 	switch a.add {
// 	case ".":
// 		a.d.AddOne()
// 	}
// }

//ExecDeleteOne

//ExecUpdateOne

//ExecClean
