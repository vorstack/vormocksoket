package main

import (
	"ru.vorstack/vormocksoket/server"
	serverImpl "ru.vorstack/vormocksoket/server/impl"
)

func main() {
	settings := server.Settings{"webSocket", true, 8080, "../ws"}
	serverImpl.WebSocketServer{}.Run(settings)
}
