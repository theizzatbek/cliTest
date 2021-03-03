package log

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"time"
)

var file *os.File

func Info(desc string) {
	fmt.Println(desc)
}

func Error(err error, args ...interface{}) {

	fmt.Sprintf("[ERROR] %v", err)
	//fmt.Println("ERR:", err.Error())
	logToFile(err, args...)
	//logrus.SetOutput(f)
	//logrus.WithFields(logrus.Fields{
	//	"desc": desc,
	//}).Trace(err)

}

func init() {
	file, _ = os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func logToFile(err error, args ...interface{}) {

	//defer f.Close()
	fmt.Fprintf(file, "[%s]\t \n%+v\n",
		time.Now().Format(time.RFC3339), errors.Wrap(err, "[ERROR] "))
}
