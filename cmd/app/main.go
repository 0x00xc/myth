/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 13:39
 */
package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"myth/internal/app"
	"os"
	"os/signal"
	"path"
	"syscall"
)

type mFormatter struct{}

func (f *mFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var rid string
	if entry.Context != nil {
		v := entry.Context.Value("request_id")
		if v != nil {
			rid, _ = v.(string)
		}
	}

	t := entry.Time.Format("2006/01/02 15:04:05")
	msg := entry.Message
	var file string
	var line int
	level := entry.Level.String()
	if entry.Caller != nil {
		file = path.Base(entry.Caller.File)
		line = entry.Caller.Line
	}
	s := fmt.Sprintf("[%s]%s %s %s:%d %s\n", level, rid, t, file, line, msg)
	return []byte(s), nil
}

func init() {
	logrus.SetFormatter(&mFormatter{})
	logrus.SetReportCaller(true)
}

func main() {
	app.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	app.Stop()
}
