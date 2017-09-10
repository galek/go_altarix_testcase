/*DEBUG HELPER*/
package main

import (
	"fmt"
	"log"
	"runtime"
)

var err error

func file_line() string {
    _, fileName, fileLine, ok := runtime.Caller(1)
    var s string
    if ok {
        s = fmt.Sprintf("%s:%d", fileName, fileLine)
    } else {
        s = ""
    }
    return s
}

func printError(_fileName string) {
	if err != nil {
		log.Fatalln("Fatal error. File: ", _fileName, " Desc: ", err.Error())
	}
}