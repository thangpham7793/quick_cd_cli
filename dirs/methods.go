package dirs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/thangpham7793/qcd/utils"
)

//Dirs a map of aliases and their abs paths
type Dirs map[string]string

//FILEPATH path to json file
const FILEPATH = "dir.json"

// getPath helper get method
func (d *Dirs) getPath(alias string) (path string, ok bool) {
	if !d.hasAlias(alias) {
		log.Fatalf(`unknown alias "%s"`, alias)
	}
	return path, true
}

func (d *Dirs) hasAlias(alias string) bool {
	_, ok := (*d)[alias]
	if !ok {
		return false
	}
	return true
}

//processPath checks if a path is valid and returns the abs path
func processPath(path string) (abs string, err error) {
	if isValidPath(path) {
		abs, err := filepath.Abs(path)
		utils.HandleErr(err, "get the absolute path")
		return abs, nil
	}
	return "", fmt.Errorf(`invalid path: %s`, path)
}

func isValidPath(path string) (valid bool) {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

//Init fills Dirs with existing aliases and abs as k-v pairs
func (d *Dirs) Init(path string) {
	pairs, err := ioutil.ReadFile(FILEPATH)

	if err != nil {
		*d = make(map[string]string)
		err = ioutil.WriteFile(FILEPATH, []byte("{}"), 0611)
		utils.HandleErr(err, "create a new empty json file")
		return
	}

	err = json.Unmarshal(pairs, &d)
	utils.HandleErr(err, "parse json")
}

//AddCurrent the current path and its alias to dirs
func (d *Dirs) AddCurrent(alias string) {
	cwd, err := os.Getwd()
	utils.HandleErr(err, "get current working dir")

	d.AddOne(alias, cwd)
	d.Save()
}

//List lists all saved paths and aliases
func (d *Dirs) List() {
	for a, p := range *d {
		fmt.Printf("%s: %s\n", a, p)
	}
}

//ListOne lists a pair for an exact matched alias
func (d *Dirs) ListOne(alias string) {
	if d.hasAlias(alias) {
		dir, _ := d.getPath(alias)
		fmt.Printf("%s: %s", alias, dir)
	}
}

//Save saves all changes
func (d *Dirs) Save() {
	j, err := json.Marshal(*d)
	utils.HandleErr(err, "marshal json")
	err = ioutil.WriteFile(FILEPATH, j, 0611)
	utils.HandleErr(err, "write to file")
}

//AddOne adds a new pair of alias and path
func (d *Dirs) AddOne(alias, path string) {
	if clean, err := processPath(path); err == nil && !d.hasAlias(alias) {
		(*d)[alias] = clean
		d.Save()
	} else {
		utils.HandleErr(err, "add with the given path")
	}

}

//UpdateOne updates an existing path with the given new path
func (d *Dirs) UpdateOne(alias, path string) {
	if clean, err := processPath(path); err == nil && d.hasAlias(alias) {
		(*d)[alias] = clean
		d.Save()
	} else {
		utils.HandleErr(err, "update with the given path")
	}
}

//DeleteOne an entry if the alias key exists
func (d *Dirs) DeleteOne(alias string) {
	if d.hasAlias(alias) {
		delete(*d, alias)
		d.Save()
	} else {
		fmt.Printf("no such alias %s", alias)
	}
}

//Clean removes aliases with invalid paths
func (d *Dirs) Clean() {
	for a, p := range *d {
		_, err := os.Stat(p)
		if err != nil {
			fmt.Printf("Deleting entry %s: %s", a, p)
			d.DeleteOne(a)
		}
	}
}
