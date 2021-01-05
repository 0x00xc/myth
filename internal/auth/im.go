/**
 * singsen
 * SNQU
 * laixingchen@sndo.com
 * 2021/1/5 15:03
 */
package auth

import (
	"errors"
	"golang.org/x/net/websocket"
	"strings"
)

func WSAuth(conn *websocket.Conn) (string, error) {
	token := strings.TrimSpace(strings.TrimPrefix(conn.Request().Header.Get("Authorization"), "Bearer "))
	if token == "" {
		if c, _ := conn.Request().Cookie("token"); c != nil {
			token = c.Value
		}
	}
	if token != "" {
		if info, err := Verify(token); err == nil {
			return info.OpenID, err
		}
	}
	return "", errors.New("authorize failed")
}
