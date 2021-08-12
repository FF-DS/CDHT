package Util

import (
	"fmt"
)

type Logger struct{
    debug bool
}

func (logger *Logger) ToggleDebug(){
    logger.debug = !logger.debug
}

func (logger *Logger) GetToggleDebug() (status bool){
    return logger.debug 
}

func (logger Logger) Log(message string){
    if logger.debug {
        fmt.Println(message)
    }
}