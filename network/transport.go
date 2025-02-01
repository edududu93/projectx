package network

//import "fmt"

type NetAddr string

type RPC struct {
	From NetAddr
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC // YT: Consume <-chan RPC ### CHAT: Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error // YT: SendMessage(NetAddr, payload []byte) error ### CHAT: SendMessage(to NetAddr, []byte) error
	Addr() NetAddr
}
