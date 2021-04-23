package log

import (
	sentryService "cliTest/services/sentry"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"time"
)

var file *os.File

func Info(desc string) {
	fmt.Println(desc)
}

func Error(err error) {

	fmt.Printf("[ERROR] %v", err)
	sentryService.SentryService.Error(nil, "", err)
	logToFile(err)
	//logrus.SetOutput(f)
	//logrus.WithFields(logrus.Fields{
	//	"desc": desc,
	//}).Trace(err)

}

func init() {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		panic(os.Mkdir("log", os.ModePerm))
	}
	file, _ = os.OpenFile(fmt.Sprintf("log/%s.log", time.Now().Month()),
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

func logToFile(err error) {

	//defer f.Close()
	_, _ = fmt.Fprintf(file, "[%s]\t \n%+v\n",
		time.Now().Format(time.RFC3339), errors.Wrap(err, "[ERROR] "))

}
