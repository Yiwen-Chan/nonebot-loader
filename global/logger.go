package global

import (
	"log"
)

func INFO(v ...interface{}) {
	log.SetPrefix("[INFO NB-LOADER] ")
	log.Println(v...)
}

func WARN(v ...interface{}) {
	log.SetPrefix("[WARN NB-LOADER] ")
	log.Println(v...)
}

func ERROR(v ...interface{}) {
	log.SetPrefix("[ERROR NB-LOADER] ")
	log.Println(v...)
}
