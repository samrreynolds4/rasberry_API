package utils

import "log"

type Logger struct {
	Source string
	pkg    string
}

func (l *Logger) Log(msg string) {
	log.Printf("%s/%s: [INFO]: %s", l.pkg, l.Source, msg)
}

func (l *Logger) LogObj(msg string, obj interface{}) {
	log.Printf("%s/%s: [INFO]: %s \n %v", l.pkg, l.Source, msg, obj)
}

func (l *Logger) LogError(msg string) {
	log.Fatalf("%s/%s: [ERROR]: %s", l.pkg, l.Source, msg)
}

func (l *Logger) LogErrorObj(msg string, obj interface{}) {
	log.Fatalf("%s/%s: [ERROR]: %s \n %v", l.pkg, l.Source, msg, obj)
}
