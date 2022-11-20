package main

import (
	"fmt"
	"github.com/Telefonica/nfqueue"
)

var queueCfg = &nfqueue.QueueConfig{
	MaxPackets: 1000,
	BufferSize: 16 * 1024 * 1024,
	QueueFlags: []nfqueue.QueueFlag{nfqueue.FailOpen},
}

type handler struct {}
// Handle a nfqueue packet. It implements nfqueue.PacketHandler interface.
func (h *handler) Handle(p *nfqueue.Packet) {
	// Accept the packet
	fmt.Println(p)
	p.Accept()
}

func main() {
	hx := &handler {}
	q := nfqueue.NewQueue(0, hx, queueCfg)
	go func() {
		q.Start()
	}()

	select{}
}
	