/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2021/1/6 11:40
 */
package chat

import (
	"github.com/sirupsen/logrus"
	"myth/internal/im"
)

func (s *Server) onClientMessage(id string, data []byte, messenger im.Messenger) {
	logrus.Infoln(id, string(data))
	messenger.Send(id, []byte("hah"))
}
