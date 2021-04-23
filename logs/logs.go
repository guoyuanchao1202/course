package logs

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	logFileName = flag.String("log", "course.log", "Log file name")
)

func InitLog()  {
	flag.Parse()
	//set logfile Stdout
	logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("init log success")
}