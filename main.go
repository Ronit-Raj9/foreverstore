package main
import (
	"fmt"
	"log"
	"github.com/Ronit-Raj9/foreverstore/p2p"
)

func main(){

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":8005",
		Decoder: p2p.DefaultDecoder{},
		HandshakeFunc: p2p.NOPHandshakeFunc,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <- tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("We Gucci!!")

	select {}
}

