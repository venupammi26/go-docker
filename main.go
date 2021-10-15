// This is a test of go logging using the lumberjack module
// to handle file rotation.
package main

import (
    "log"
    "github.com/natefinch/lumberjack"
    "time"
	"encoding/json"
	//"fmt"
)

type Fields struct {
    Name     string
    NickName string
    Age      int
    Addr     string
    Mobile   int
    Blood_G  string
}

func main() {

	var f = Fields{
        Name:     "Abc",
        NickName: "abc",
        Age:      19,
        Addr :    "Hyd",
        Mobile :  5788989,
        Blood_G :  "B+",
    }

	jsonLogger :=&lumberjack.Logger{
			Filename:   "student.json",
			MaxSize:    1, // megabytes
			MaxBackups: 3,
			MaxAge:     7, //days
			Compress:   true, // disabled by default
		}
	log.SetOutput(jsonLogger)
    log.Println("Entering main")
   
	
	enc := json.NewEncoder(jsonLogger)
	for i := 0; i < 300000; i++ {
        time.Sleep(1 * time.Second)
        log.Println(enc.Encode(f), i)
    }
}

