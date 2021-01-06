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
	"net/http"
)

type HTTPServer struct {
	srv *http.Server
}

func (s *HTTPServer) init() {
	s.srv = new(http.Server)
	s.srv.Addr = fmt.Sprintf(":%d", conf.C.AppPort)

	//opt := im.NewOptions()
	//opt.Auth = auth.WSAuth
	//opt.OnClientConnected = func(id string, messenger im.Messenger) { logrus.Infoln(id, "connected") }
	//opt.OnClientDisConnected = func(id string, messenger im.Messenger) { logrus.Infoln(id, "disconnected") }
	//opt.OnClientMessage = func(id string, data []byte, messenger im.Messenger) {
	//	logrus.Infoln(id, string(data))
	//	messenger.Send(id, []byte("hah"))
	//}
	//s.chat = im.NewManager(opt)
}

func (s *HTTPServer) Serve() {
	s.init()
	s.route()
	//if len(conf.C.AppCert) == 2 {
	//	//TODO
	//}
	logrus.Infoln("http server listen on", s.srv.Addr)
	err := s.srv.ListenAndServe()
	if err != nil {
		logrus.Errorln(err)
	}
}

func (s *HTTPServer) Stop() {
	s.srv.Shutdown(context.Background())
}
