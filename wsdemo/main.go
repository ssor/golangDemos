// wsdemo project main.go
package main

import (
	"../go.net/websocket"
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const (
	listenAddr = "192.168.1.103:4000" // server address
)

var (
	pwd, _        = os.Getwd()
	RootTemp      = template.Must(template.ParseFiles(pwd + "/chat.html"))
	JSON          = websocket.JSON           // codec for JSON
	Message       = websocket.Message        // codec for string, []byte
	ActiveClients = make(map[ClientConn]int) // map containing clients
)

// Initialize handlers and websocket handlers
func init() {
	http.HandleFunc("/", RootHandler)
	http.Handle("/sock", websocket.Handler(SockServer))
}

// Client connection consists of the websocket and the client ip
type ClientConn struct {
	websocket *websocket.Conn
	clientIP  string
}

// WebSocket server to handle chat between clients
func SockServer(ws *websocket.Conn) {
	var err error
	var clientMessage string
	// use []byte if websocket binary type is blob or arraybuffer
	// var clientMessage []byte

	// cleanup on server side
	defer func() {
		if err = ws.Close(); err != nil {
			log.Println("Websocket could not be closed", err.Error())
		}
	}()

	client := ws.Request().RemoteAddr
	log.Println("Client connected:", client)
	sockCli := ClientConn{ws, client}
	ActiveClients[sockCli] = 0
	log.Println("Number of clients connected ...", len(ActiveClients))

	// for loop so the websocket stays open otherwise
	// it'll close after one Receieve and Send
	for {
		if err = Message.Receive(ws, &clientMessage); err != nil {
			// If we cannot Read then the connection is closed
			log.Println("Websocket Disconnected waiting", err.Error())
			// remove the ws client conn from our active clients
			delete(ActiveClients, sockCli)
			log.Println("Number of clients still connected ...", len(ActiveClients))
			return
		}

		clientMessage = sockCli.clientIP + " Said: " + clientMessage
		for cs, _ := range ActiveClients {
			if err = Message.Send(cs.websocket, clientMessage); err != nil {
				// we could not send the message to a peer
				log.Println("Could not send message to ", cs.clientIP, err.Error())
			}
		}
	}
}

// RootHandler renders the template for the root page
func RootHandler(w http.ResponseWriter, req *http.Request) {
	err := RootTemp.Execute(w, listenAddr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

//func ChatWith(ws *websocket.Conn) {
//	var err error

//	for {
//		var reply string

//		if err = websocket.Message.Receive(ws, &reply); err != nil {
//			fmt.Println("Can't receive")
//			break
//		}

//		fmt.Println("Received back from client: " + reply)

//		//msg := "Received from " + ws.Request().Host + "  " + reply
//		msg := "welcome to websocket do by pp"
//		fmt.Println("Sending to client: " + msg)

//		if err = websocket.Message.Send(ws, msg); err != nil {
//			fmt.Println("Can't send")
//			break
//		}
//	}
//}

//func main() {
//	//
//	http.Handle("/", websocket.Handler(ChatWith))
//	//http.HandleFunc("/chat", Client)

//	fmt.Println("listen on port 8001")
//	//fmt.Println("visit http://127.0.0.1:8001/chat with web browser(recommend: chrome)")

//	if err := http.ListenAndServe(":8001", nil); err != nil {
//		log.Fatal("ListenAndServe:", err)
//	}
//}

//func main() {
//	http.Handle("/", http.FileServer(http.Dir("."))) // <-- note this line
//	http.Handle("/socket", websocket.Handler(Echo))

//	if err := http.ListenAndServe(":1234", nil); err != nil {
//		log.Fatal("ListenAndServe:", err)
//	}
//}
//func Echo(ws *websocket.Conn) {
//	var err error

//	for {
//		var reply string

//		if err = websocket.Message.Receive(ws, &reply); err != nil {
//			fmt.Println("Can't receive")
//			break
//		}

//		fmt.Println("Received back from client: " + reply)

//		msg := "Received:  " + reply
//		fmt.Println("Sending to client: " + msg)

//		if err = websocket.Message.Send(ws, msg); err != nil {
//			fmt.Println("Can't send")
//			break
//		}
//	}
//}
