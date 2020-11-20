package app

//Cmds contains registered flags
var Cmds map[string]strCmd = map[string]strCmd{
	"add":   strCmd{"add", ".", "the current or specified dir to be stored"},
	"alias": strCmd{"alias", "empty", "the given alias to be stored"},
	"ls":    strCmd{"ls", "all", "prints all stored aliases and paths"},
}
