package chat_worker

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func Run(configPath string) {
	fmt.Println(`
 ================================================
 \\\   ######~~#####~~~##~~~~~~#####~~~#####   \\\
  \\\  ~~##~~~~##~~##~~##~~~~~~##~~~~~~##~~##   \\\
   ))) ~~##~~~~##~~##~~##~~~~~~####~~~~#####     )))
  ///  ~~##~~~~##~~##~~##~~~~~~##~~~~~~##~~##   ///
 ///   ######~~#####~~~######~~#####~~~##~~##  ///
 ================================================
	`)

	//cfg, err := config.Init(configPath)
	//if err != nil {
	//	logrus.Fatalf("error initializing configs: %s", err.Error())
	//}

	fmt.Println("Go WS")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))

	//logrus.Print("IDLER chat-worker application has started")
	//
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	//<-quit
	//
	//logrus.Print("IDLER facade application shutting down")
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
	http.HandleFunc("/ws", serveWS)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client has Successfully Connected")

	reader(ws)
}
