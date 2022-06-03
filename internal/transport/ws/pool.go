package ws

import (
	"fmt"
)

const (
	DialogSize = 2
)

type Pool struct {
	ID         string
	IsDialog   bool
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Send       chan Message
}

func NewPool(id string, isDialog bool) *Pool {
	return &Pool{
		ID:         id,
		IsDialog:   isDialog,
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Send:       make(chan Message),
	}
}

func (p *Pool) Start() {
	for {
		select {
		case client := <-p.Register:
			p.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(p.Clients))
			for client, _ := range p.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-p.Unregister:
			client.Delete(p, client)
			fmt.Println("Size of Connection Pool: ", len(p.Clients))

			if len(p.Clients) == 0 {
				p.Stop(p.ID)
			}

			for client, _ := range p.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-p.Send:
			fmt.Println("Sending message to clients in Pool")
			for client, _ := range p.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

func (p *Pool) Stop(id string) {
	//delete(Pools, id)
}
