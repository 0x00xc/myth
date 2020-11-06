/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2020/10/30 13:46
 */
package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
)

type HTTPServer struct {
	srv *http.Server
}

func (s *HTTPServer) Serve() {
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
