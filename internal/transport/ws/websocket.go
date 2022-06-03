package ws

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/pkg/cache"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func SetupHandler(ctx context.Context, poolCache *cache.MemoryCache[string, Pool]) {
	http.HandleFunc("/ws/", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")

		pool, err := poolCache.Get(ctx, id)
		if err != nil {
			pool = NewPool(id, true)
			poolCache.Set(ctx, id, pool)
			go pool.Start()
		}

		serveWs(pool, w, r)
	})
}

func serveWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := getConnection(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := NewClient(conn, pool)

	pool.Register <- client
	go client.Read()
}

func getConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}
