package utils

import "log"

func CheckConf(inp interface{}) {
	if inp == "" {
		log.Fatal("Configurations are not compelete.")
	}
}
