/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2021/1/6 11:38
 */
package chat

import (
	"github.com/sirupsen/logrus"
	"myth/internal/auth"
	"myth/internal/im"
	"net/http"
)

type Server struct {
	chat *im.Manager
}

func (s *Server) Init() {
	opt := im.NewOptions()
	opt.Auth = auth.WSAuth
	opt.OnClientConnected = func(id string, messenger im.Messenger) { logrus.Infoln(id, "connected") }
	opt.OnClientDisConnected = func(id string, messenger im.Messenger) { logrus.Infoln(id, "disconnected") }
	opt.OnClientMessage = s.onClientMessage
	s.chat = im.NewManager(opt)
}

func (s *Server) Handler() http.Handler {
	return s.chat.Handler()
}

func (s *Server) Serve() {
	s.chat.Serve()
}

func (s *Server) Stop() {
	s.chat.Stop()
}
