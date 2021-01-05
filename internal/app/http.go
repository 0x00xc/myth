/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 13:46
 */
package app

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"myth/conf"
	"myth/internal/auth"
	"myth/internal/im"
	"net/http"
)

type HTTPServer struct {
	srv  *http.Server
	chat *im.Manager
}

func (s *HTTPServer) init() {
	s.srv = new(http.Server)
	s.srv.Addr = fmt.Sprintf(":%d", conf.C.AppPort)
	s.chat = im.NewManager(nil, nil)

	s.chat.Auth = auth.WSAuth
	s.chat.OnConnected = func(id string, messenger im.Messenger) { logrus.Infoln(id, "connected") }
	s.chat.OnDisConnected = func(id string, messenger im.Messenger) { logrus.Infoln(id, "disconnected") }
}

func (s *HTTPServer) Serve() {
	s.init()
	s.route()
	//if len(conf.C.AppCert) == 2 {
	//	//TODO
	//}
	go s.chat.Serve() //TODO
	logrus.Infoln("http server listen on", s.srv.Addr)
	err := s.srv.ListenAndServe()
	if err != nil {
		logrus.Errorln(err)
	}
}

func (s *HTTPServer) Stop() {
	s.srv.Shutdown(context.Background())
	s.chat.Stop()
}
