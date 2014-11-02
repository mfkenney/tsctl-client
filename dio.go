// Messages pertaining to the DIO class
package tsctl

type DioCommand uint8

const (
	DioLock DioCommand = iota
	DioUnlock
	DioPrempt
	DioRefresh
	DioCommit
	DioSet
	DioGet
	DioSetAsync
	DioGetAsync
	DioWait
	DioCount
	DioCapabilities
	DioGetMulti
)

type DioState int32

const (
	InputLow  DioState = -3
	InputHigh DioState = -2
	Input     DioState = -1
	Low       DioState = 0
	High      DioState = 1
)

// Command to lock access to the DIO class
func DioLockMsg(l LockType) ([]byte, error) {
	return PackMsg(TsReq{Dio, 0, uint8(Lock)}, uint32(0), l)
}

// Command to unlock access to the DIO class
func DioUnlockMsg(l LockType) ([]byte, error) {
	return PackMsg(TsReq{Dio, 0, uint8(Unlock)}, uint32(0), l)
}

// Command to get the current state of a DIO line
func DioGetAsyncMsg(num uint32) ([]byte, error) {
	return PackMsg(TsReq{Dio, 0, uint8(DioGetAsync)}, num)
}

// Command to set the state of a DIO line
func DioSetAsyncMsg(num uint32, state DioState) ([]byte, error) {
	return PackMsg(TsReq{Dio, 0, uint8(DioSetAsync)}, num, int32(state))
}
