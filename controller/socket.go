package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ConnectUser struct {
	Websocket *websocket.Conn
	ClientIP  string
}

func newConnectUser(ws *websocket.Conn, clientIP string) *ConnectUser {
	return &ConnectUser{
		Websocket: ws,
		ClientIP:  clientIP,
	}
}

type chat_message struct {
	Username string
	Message  string
	Role     string
}

var users = make(map[ConnectUser]int)

func (s *DataBase) WebsocketHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ws, _ := upgrader.Upgrade(w, r, nil)

	defer func() {
		if err := ws.Close(); err != nil {
			log.Println("Websocket could not be closed", err.Error())
		}
	}()

	log.Println("Client connected:", ws.RemoteAddr().String())
	var socketClient *ConnectUser = newConnectUser(ws, ws.RemoteAddr().String())
	users[*socketClient] = 0
	log.Println("Number client connected ...", len(users))
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Ws disconnect waiting", err.Error())
			delete(users, *socketClient)
			log.Println("Number of client still connected ...", len(users))
			return
		}

		session, _ := loggedUserSession.Get(r, "authenticated-user-session")
		userID, ok := session.Values["userID"]
		if !ok {
			continue
		}
		row := s.Data.QueryRow(`SELECT username, role FROM users_account WHERE id = ?`, userID)
		var username, role string
		err = row.Scan(&username, &role)
		_ = err
		msg := chat_message{Username: username, Message: string(message), Role: role}
		message, err = json.Marshal(&msg)

		if err != nil {
			fmt.Println("Err")
			continue
		}

		for client := range users {
			if err = client.Websocket.WriteMessage(messageType, message); err != nil {
				log.Println("Cloud not send Message to ", client.ClientIP, err.Error())
			}
		}

	}
}
