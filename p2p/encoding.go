package p2p

import (
	"encoding/gob"
	"io"
) 

type Decoder interface {	
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct {

}

func (dec GOBDecoder) Decode(r io.Reader, msg *RPC) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, msg *RPC) error {
	buf := make([]byte, 2000)
	n, err := r.Read(buf)

	if err != nil {
		return err
	}

	msg.Payload = buf[:n]
	return nil
}