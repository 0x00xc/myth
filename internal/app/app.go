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
	"myth/internal/dao"
	"myth/internal/email"
	"myth/pkg/cache"
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
}

var httpserver = new(HTTPServer)

func Start() {
	prepare()

	go httpserver.Serve()

}

func Stop() {
	httpserver.Stop()
}
