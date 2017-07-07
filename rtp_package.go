package RTP

import (
	"encoding/binary"
	"log"
)

/*
 0               1               2               3              4
 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7 0 1 2 3 4 5 6 7
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

|V=2|P|X| CC    |M|      PT     |       sequence number         |

+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

|                           timestamp                           |

+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

|                           SSRC                                |

+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

|                           CSRC                                |

|                           ....                                |

+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
*/
//RtpHeader the struct of the RTP header
type Header struct {
	Version     uint8
	Padding     uint8
	Extension   uint8
	Cc          uint8
	Marker      uint8
	PayloadType uint8
	Sequence    uint16
	TimeStamp   uint32
	SSRC        uint32
	CSRC        []uint32
}

//ParseRtpHeader parse the rtp header
func ParseRtpHeader(buf []byte) *Header {
	len := len(buf)
	if len < 12 {
		log.Println("parse rtp header failed")
		return nil
	}
	if len%4 != 0 {
		log.Printf("error,illegal header's len  %d", len)
		return nil
	}
	header := Header{}

	header.Version = (buf[0] & 0xc0) >> 6
	header.Padding = (buf[0] & 0x20) >> 5
	header.Extension = (buf[0] & 0x10) >> 4
	header.Cc = buf[0] & 0x0f
	header.Marker = (buf[1] & 0x8f) >> 7
	header.Padding = buf[1] & 0x7f
	header.Sequence = binary.BigEndian.Uint16(buf[2:])
	header.TimeStamp = binary.BigEndian.Uint32(buf[4:])
	header.SSRC = binary.BigEndian.Uint32(buf[8:])
	return &header
}
