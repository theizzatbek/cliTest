package cmd

import (
	"fmt"
	"os"
	"time"
)

func print(desc string, err error) {
	fmt.Println(desc)

	if err != nil {
		f, _ := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer f.Close()
		fmt.Fprintf(f, "[%s]\t%s: %+v\n", time.Now().Format(time.RFC3339), desc, err)
		//logrus.SetOutput(f)
		//logrus.WithFields(logrus.Fields{
		//	"desc": desc,
		//}).Trace(err)

	}

}
