package main

import (
	"fmt"
	"os"
	"github.com/Telefonica/nfqueue"
)

var how string

var queueCfg = &nfqueue.QueueConfig{
	MaxPackets: 1000,
	BufferSize: 16 * 1024 * 1024,
	QueueFlags: []nfqueue.QueueFlag{nfqueue.FailOpen},
}

type handler struct {}
func (h *handler) Handle(p *nfqueue.Packet) {
	// TODO: write packet handler code here
	if how == "accept" {
		p.Accept()
	}
	if how == "drop" {
		p.Drop()
	}
	fmt.Println(p)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Must indicate if accept or drop")
		os.Exit(1)
	}

	switch (os.Args[1]) {
	case "accept":
		fmt.Println("Filter mode: ACCEPT")
		how = "accept"
	case "drop":
		fmt.Println("Filter mode: DROP")
		how = "drop"
	default:
		return
	}

	hx := &handler {}
	q := nfqueue.NewQueue(0, hx, queueCfg)
	go func() {
		q.Start()
	}()

	select{}
}