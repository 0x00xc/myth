/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 13:45
 */
package app

import (
	"github.com/sirupsen/logrus"
	"myth/internal/auth"
	"myth/internal/chat"
	"myth/internal/dao"
	"myth/internal/email"
	"myth/pkg/cache"
	"sync"
	"time"
)

func prepare() {
	err := dao.Init()
	if err != nil {
		logrus.Panic(err)
	}
	storage := cache.NewLocalStorage()
	go storage.Cleaning(time.Hour)
	auth.Init(storage)
	email.Init(storage)
	chat.Init()
}

var httpserver = new(HTTPServer)

func Start() {
	prepare()

	do(httpserver.Serve)
	do(chat.Serve)
}

func Stop() {
	httpserver.Stop()
	chat.Stop()
	wg.Wait()
	logrus.Infoln("stopped")
}

var wg = new(sync.WaitGroup)

func do(f func()) {
	wg.Add(1)
	go func() {
		f()
		wg.Done()
	}()
}
