package engine

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允許所有來源（僅供測試）
	},
}

type ResolveFunc func(b []byte) string

type Engine struct {
	RouterGroup
	route          *Router
	resolver       ResolveFunc
	channelManager *channelManager
}

func New() *Engine {
	e := &Engine{
		RouterGroup: RouterGroup{
			Handlers: make([]HandlerFunc, 0),
			basePath: "",
		},
		route: NewRouter(),
		channelManager: &channelManager{
			channels: make(map[string]*channel),
		},
	}
	e.RouterGroup.engine = e
	return e
}

func (e *Engine) Use(middleware HandlerFunc) {
	e.RouterGroup.Use(middleware)
}

func (e *Engine) addRoute(path string, handlers []HandlerFunc) {
	e.route.Add(path, handlers)
}

func (e *Engine) getRoute(path string) HandlerFunc {
	return nil
}

func (e *Engine) PathResolver(resolver ResolveFunc) {
	e.resolver = resolver
}

func (e *Engine) Run(addr string) error {
	// 每次訪問路徑時，都會創建一個 ServeWs 並且不會結束(for 死循環)
	http.HandleFunc("/game", e.ServeWs)
	// 啟動伺服器
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}
	return nil
}

func (e *Engine) ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 註冊用戶
	client := &Client{conn: conn, engine: e, keys: make(map[string]interface{}), send: make(chan []byte)}
	go client.run()
	e.channelManager.Add("default", client)
}
