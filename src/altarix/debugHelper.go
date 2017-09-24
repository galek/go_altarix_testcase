/*DEBUG HELPER*/
package main

import (
	"fmt"
	"log"
	"runtime"
)

var err error

/**/
var ISDebug bool = false
var ISShowSendGetReq bool = true

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
		if ISDebug {
			log.Fatalln("Fatal error. File: ", _fileName, " Desc: ", err.Error())
		} else {
			if errlog != nil {
				errlog.Fatalln("Fatal error. File: ", _fileName, " Desc: ", err.Error())
			} else {
				log.Fatalln("Fatal error. File: ", _fileName, " Desc: ", err.Error())
			}
		}
	}
}
