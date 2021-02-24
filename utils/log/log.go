package log

import (
	"fmt"
	"os"
	"time"
)

func Info(desc string) {
	fmt.Println(desc)
}

func Error(desc string, err error) {
	fmt.Println(desc)
	f, _ := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	fmt.Fprintf(f, "[%s]\t%s: %+v\n", time.Now().Format(time.RFC3339), desc, err)
	//logrus.SetOutput(f)
	//logrus.WithFields(logrus.Fields{
	//	"desc": desc,
	//}).Trace(err)

}
