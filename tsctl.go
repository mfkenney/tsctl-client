// Package implements a subset of the tsctl TCP protocol
// developed by Technologic Systems.
//
// See http://wiki.embeddedarm.com/wiki/Tsctl#tsctl_TCP_protocol
package tsctl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type TsClass uint16

const (
	System TsClass = iota
	Bus
	Time
	Pin
	Dioraw
	Dio
	Twi
	Can
	Spi
	Aio
	Edio
)

// The most basic message type. All fields must be exported
// so the structure can be deserialized with UnpackMsg
type TsReq struct {
	Class    TsClass
	Instance uint8
	Command  uint8
}

type TsReply struct {
	Class    TsClass
	Instance uint8
	Tag      uint8
}

const (
	tag_strlen    = 0x50
	tag_int       = 0x13
	tag_end       = 0x80
	tag_int_array = 0x53
	tag_byte      = 0x00
	tag_word      = 0x01
)

type LockType uint32

const (
	NonBlocking LockType = 1
	Shared      LockType = 2
	NoUnlock    LockType = 4
)

type ScalarReply struct {
	TsReply
	Value  int32
	Endtag uint8
}

type StringReply struct {
	TsReply
	Strlen uint32
	Value  []byte
	Endtag uint8
}

func assert_true(test bool, msg string) {
	if test != true {
		panic(msg)
	}
}

// Convert a message to its wire format
func PackMsg(hdr TsReq, params ...interface{}) ([]byte, error) {
	var err error
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, hdr)
	if err == nil {
		for _, p := range params {
			err = binary.Write(buf, binary.LittleEndian, p)
			if err != nil {
				break
			}
		}
	}
	return buf.Bytes(), err
}

func read_multi(buf io.Reader, vals ...interface{}) error {
	var err error
	for _, v := range vals {
		err = binary.Read(buf, binary.LittleEndian, v)
		if err != nil {
			break
		}
	}
	return err
}

func read_string(buf io.Reader) ([]byte, error) {
	var (
		strval []byte
		n      uint32
	)
	err := binary.Read(buf, binary.LittleEndian, &n)
	if err == nil {
		strval = make([]byte, n)
		err = binary.Read(buf, binary.LittleEndian, strval)
	}
	return strval, err
}

// Convert a reply from its wire format
func UnpackReply(buf io.Reader, reply interface{}) error {
	var err error
	var ival uint32
	var sval uint16
	var bval uint8

	hdr := TsReply{}
	switch t := reply.(type) {
	case *TsReply:
		return binary.Read(buf, binary.LittleEndian, reply)
	case *ScalarReply:
		err = binary.Read(buf, binary.LittleEndian, &hdr)
		switch hdr.Tag {
		case tag_byte:
			err = binary.Read(buf, binary.LittleEndian, &bval)
			t.Value = int32(bval)
		case tag_word:
			err = binary.Read(buf, binary.LittleEndian, &sval)
			t.Value = int32(sval)
		case tag_int, 0xc4, 0xc0, 0x03:
			err = binary.Read(buf, binary.LittleEndian, &ival)
			t.Value = int32(ival)
		default:
			return fmt.Errorf("invalid packet header: %v", hdr)
		}

		if err == nil {
			err = binary.Read(buf, binary.LittleEndian, &bval)
			t.Endtag = bval
		}
	case *StringReply:
		err = binary.Read(buf, binary.LittleEndian, &hdr)
		if hdr.Tag != tag_strlen {
			return fmt.Errorf("invalid packet header: %v", hdr)
		}
		t.TsReply = hdr
		t.Value, err = read_string(buf)
		if err == nil {
			t.Strlen = uint32(len(t.Value))
			err = binary.Read(buf, binary.LittleEndian, &bval)
			t.Endtag = bval
		}
	default:
		return fmt.Errorf("invalid reply type: %v", t)
	}
	return err
}
