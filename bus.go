// Messages pertaining to the Bus class
package tsctl

import (
	"fmt"
)

type BusCommand uint8

const (
	Lock BusCommand = iota
	Unlock
	Prempt
	Peek8
	Poke8
	Peek16
	Poke16
	Peek32
	Poke32
	BitGet8
	BitAssign8
	BitSet8
	BitClear8
	BitGet16
	BitAssign16
	BitSet16
	BitClear16
	BitGet32
	BitAssign32
	BitSet32
	BitClear32
	PeekStream
	PokeStream
	Refresh
	Commit
	BitToggle8
	BitToggle16
	BitToggle32
	Assign8X
	Assign16X
	Assign32X
	BitsGet8
	BitsGet16
	BitsGet32
)

// Command to lock the Bus
func BusLockMsg(l LockType) ([]byte, error) {
	return PackMsg(TsReq{Bus, 0, uint8(Lock)}, uint32(0), l)
}

// Command to unlock the Bus
func BusUnlockMsg(l LockType) ([]byte, error) {
	return PackMsg(TsReq{Bus, 0, uint8(Unlock)}, uint32(0), l)
}

// Command to read a memory address
func PeekMsg(address uint32, size int) ([]byte, error) {
	var cmd BusCommand

	switch size {
	case 8:
		cmd = Peek8
	case 16:
		cmd = Peek16
	case 32:
		cmd = Peek32
	default:
		return nil, fmt.Errorf("Invalid size: %v", size)
	}

	return PackMsg(TsReq{Bus, 0, uint8(cmd)}, address)
}

// Command to write to a memory address
func PokeMsg(address uint32, size int, val uint32) ([]byte, error) {
	switch size {
	case 8:
		return PackMsg(TsReq{Bus, 0, uint8(Poke8)}, address, uint8(val))
	case 16:
		return PackMsg(TsReq{Bus, 0, uint8(Poke16)}, address, uint16(val))
	case 32:
		return PackMsg(TsReq{Bus, 0, uint8(Poke32)}, address, val)
	}

	return nil, fmt.Errorf("Invalid size: %v", size)
}

// Command to set a bit at a memory address
func BitSetMsg(address uint32, size int, bitnum uint32) ([]byte, error) {
	var cmd BusCommand

	switch size {
	case 8:
		if bitnum >= 0 && bitnum < 8 {
			return nil, fmt.Errorf("Invalid bit number %d", bitnum)
		}
		cmd = BitSet8
	case 16:
		if bitnum >= 0 && bitnum < 16 {
			return nil, fmt.Errorf("Invalid bit number %d", bitnum)
		}
		cmd = BitSet16
	case 32:
		if bitnum >= 0 && bitnum < 32 {
			return nil, fmt.Errorf("Invalid bit number %d", bitnum)
		}
		cmd = BitSet32
	default:
		return nil, fmt.Errorf("Invalid size: %v", size)
	}

	return PackMsg(TsReq{Bus, 0, uint8(cmd)}, address, bitnum)
}

// Command to clear a bit at a memory address
func BitClearMsg(address uint32, size int, bitnum uint32) ([]byte, error) {
	var cmd BusCommand

	switch size {
	case 8:
		if bitnum >= 0 && bitnum < 8 {
			return nil, fmt.Errorf("Invalid bit number %d", bitnum)
		}
		cmd = BitClear8
	case 16:
		if bitnum >= 0 && bitnum < 16 {
			return nil, fmt.Errorf("Invalid bit number %d", bitnum)
		}
		cmd = BitClear16
	case 32:
		if bitnum >= 0 && bitnum < 32 {
			return nil, fmt.Errorf("Invalid bit number %d", bitnum)
		}
		cmd = BitClear32
	default:
		return nil, fmt.Errorf("Invalid size: %v", size)
	}

	return PackMsg(TsReq{Bus, 0, uint8(cmd)}, address, bitnum)
}

// Command to toggle a bit at a memory address
func BitToggleMsg(address uint32, size int, bitnum uint32) ([]byte, error) {
	var cmd BusCommand

	switch size {
	case 8:
		if bitnum >= 0 && bitnum < 8 {
			return nil, fmt.Errorf("Invalid bit number %d", bitnum)
		}
		cmd = BitToggle8
	case 16:
		if bitnum >= 0 && bitnum < 16 {
			return nil, fmt.Errorf("Invalid bit number %d", bitnum)
		}
		cmd = BitToggle16
	case 32:
		if bitnum >= 0 && bitnum < 32 {
			return nil, fmt.Errorf("Invalid bit number %d", bitnum)
		}
		cmd = BitToggle32
	default:
		return nil, fmt.Errorf("Invalid size: %v", size)
	}

	return PackMsg(TsReq{Bus, 0, uint8(cmd)}, address, bitnum)
}
