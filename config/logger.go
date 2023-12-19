package config

import (
	"log"
	"os"
)

var InfoLogger *log.Logger = log.New(os.Stderr, "[INFO]: ", log.Ldate|log.Ltime)
var WarningLogger *log.Logger = log.New(os.Stderr, "[WARN]: ", log.Ldate|log.Ltime)
var ErrorLogger *log.Logger = log.New(os.Stderr, "[ERROR]: ", log.Ldate|log.Ltime)
