package main

import (
	"fmt"
)

const (
    DEBUG = iota // 0
    INFO         // 1
    WARNING      // 2
    ERROR        // 3
)

func getLogLevelName(level int) string {
    switch level {
    case DEBUG:
        return "DEBUG"
    case INFO:
        return "INFO"
    case WARNING:
        return "WARNING"
    case ERROR:
        return "ERROR"
    default:
        return "UNKNOWN"
    }
}

const (
	queque01 = iota
	queque02
	queque03
	queque04
	queque05
	queque06
	queque07
	queque08
	queque09
	queque10
)

func main() {
    // Cetak semua level log
    fmt.Println("Log Levels:")
    fmt.Printf("%d - %s\n", DEBUG, getLogLevelName(DEBUG))
    fmt.Printf("%d - %s\n", INFO, getLogLevelName(INFO))
    fmt.Printf("%d - %s\n", WARNING, getLogLevelName(WARNING))
    fmt.Printf("%d - %s\n", ERROR, getLogLevelName(ERROR))

	fmt.Println("Queue02:", queque10)
}
