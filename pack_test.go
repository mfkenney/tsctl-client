// Test packing/unpacking functions
package tsctl

import (
	"bytes"
	"testing"
)

func TestPack(t *testing.T) {
	var err error
	var cc_serial []byte
	var ic_serial []byte
	var ml_serial []byte

	cc_wire := []byte{0x00, 0x00, 0x00, 0x00}
	ic_wire := []byte{0x00, 0x00, 0x00, 0x01, 0x2a, 0x00, 0x00, 0x00}
	ml_wire := []byte{0x00, 0x00, 0x00, 0x0c, 0x03, 0x00, 0x00, 0x00, 'f', 'o', 'o'}

	cc_serial, err = ClassCountMsg()
	if err != nil {
		t.Errorf("Cannot serialize message: %v", err)
	}

	if bytes.Compare(cc_wire, cc_serial) != 0 {
		t.Errorf("Error packing CC message: %v", cc_serial)
	}

	ic_serial, err = InstanceCountMsg(42)
	if err != nil {
		t.Errorf("Cannot serialize message: %v", err)
	}

	if bytes.Compare(ic_wire, ic_serial) != 0 {
		t.Errorf("Error packing IC message: %v", ic_serial)
	}

	ml_serial, err = MapLookupMsg("foo")
	if err != nil {
		t.Errorf("Cannot serialize message: %v", err)
	}

	if bytes.Compare(ml_wire, ml_serial) != 0 {
		t.Errorf("Error packing ML message: %v", ml_serial)
	}
}

func TestUnpack(t *testing.T) {
	var err error

	reply := StringReply{}
	p1_wire := []byte("\x00\x00\x11\x50\x04\x00\x00\x001.42\x80")
	buf := bytes.NewBuffer(p1_wire)
	err = UnpackReply(buf, &reply)
	if err != nil {
		t.Errorf("Cannot unpack buffer: %v", err)
	}

	if reply.Value != "1.42" {
		t.Errorf("Packet decoding error: %v", reply)
	}
}
