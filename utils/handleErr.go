package utils

import "log"

//HandleErr logs failed task, err and exits
func HandleErr(err error, task string) {
	if err != nil {
		log.Fatalf("could not %s: %v", task, err)
	}
}
