package chat_service

import (
	"context"
	"fmt"
	"github.com/eugene-krivtsov/idler/internal/transport/ws"
	"github.com/eugene-krivtsov/idler/pkg/cache"
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
	poolCache := cache.NewMemoryCache[string, ws.Pool]()
	ws.SetupHandler(context.Background(), poolCache)
	log.Fatal(http.ListenAndServe(":8081", nil))

	//logrus.Print("IDLER chat-worker application has started")
	//
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	//<-quit
	//
	//logrus.Print("IDLER facade application shutting down")
}
