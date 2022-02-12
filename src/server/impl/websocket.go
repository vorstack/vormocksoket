package impl

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"ru.vorstack/vormocksoket/parser"
	"ru.vorstack/vormocksoket/server"
	"strings"
)

const mockFileName = "message" + server.MOCK_FILE_EXTENSION

var upgrader = websocket.Upgrader{} // use default options
var messages map[string][]interface{}

type WebSocketServer struct{}

func (server WebSocketServer) Run(settings server.Settings) {
	rawByPath := getRawMessageByPath(settings.MockDir)
	messages = make(map[string][]interface{})
	for path, rawMessage := range rawByPath {
		message := parser.Parse(rawMessage)
		messages[path] = message
		http.HandleFunc(path, broadcast)
	}
	addr := fmt.Sprint("localhost:", settings.Port)
	fmt.Println("WebSocketServer is running by address: {}", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getRawMessageByPath(dirPath string) map[string][]byte {
	rawMesssages := make(map[string][]byte)

	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				panic(err)
			}
			if info.Name() == mockFileName {
				absPath, _ := filepath.Abs(path)
				content, err := ioutil.ReadFile(absPath)
				if err != nil {
					panic(err)
				}
				subWsPath := path[len(dirPath)-2 : len(path)-len(mockFileName)-1]
				wsPath := strings.ReplaceAll(subWsPath, "\\", "/") // -1 it is slash
				rawMesssages[wsPath] = content
			}
			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return rawMesssages
}

func broadcast(w http.ResponseWriter, r *http.Request) {
	log.Print("kek")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	path := r.RequestURI

	for _, message := range messages[path] {
		err = c.WriteMessage(websocket.TextMessage, []byte(message.(string)))
	}
}
