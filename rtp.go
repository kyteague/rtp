// Package rtp provides RTP packetizer and depacketizer
package rtp

import (
	"io"
)

type ReadStream interface {
	io.ReadCloser
}

type WriteStream interface {
	io.Writer
	WriteRTP(header *Header, payload []byte) (int, error)
}

type Session interface {
	AcceptStream() (ReadStream, uint32, error)
	OpenReadStream(ssrc uint32) (ReadStream, error)
	OpenWriteStream() (WriteStream, error)
}
