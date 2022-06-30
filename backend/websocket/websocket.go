package websocket

import (
	"backend-task/backend/cors"
	"context"
	"net/http"

	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocket struct {
	ws *http.Server
}

func NewHandler(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	Clients[ws] = true
	return ws, nil
}

func (ws *WebSocket) Init(port string, handler http.Handler) {
	ws.ws = &http.Server{
		Addr:    port,
		Handler: cors.CORSEnabled(handler),
	}
}

func (ws *WebSocket) Run() error {

	if err := ws.ws.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
func (ws *WebSocket) Stop(ctx context.Context) error {

	err := ws.ws.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}
