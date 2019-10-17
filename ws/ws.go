package ws

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// Send Send
func Send(path string, from string, to []string, title, contents string) error {
	if h, ok := m[path]; ok {
		h.broadcast <- &Broadcast{From: from, To: to, Title: title, Content: contents}
		return nil
	}
	return errors.New("无效的path")
}

var m = make(map[string]*Hub)

// Serve Serve
func Serve(e *gin.Engine, path string) error {
	if _, ok := m[path]; ok {
		return errors.New("重复的path")
	}
	hub := newHub()
	go hub.run()
	m[path] = hub
	e.GET(path, func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request, c.GetString("username"))
	})
	return nil
}
