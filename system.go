// Messages pertaining to the System class
package tsctl

type SystemCommand uint8

const (
	ClassCount SystemCommand = iota
	InstanceCount
	ApiCount
	LockCount
	LockHolderInfo
	ConnWaitInfo
	CANBusGet
	BuildTime
	ModelId
	BaseBoardId
	MapLength
	MapGet
	MapLookup
	MapLookupPartial
	MapAdd
	MapDelete
	Note
	Version
	UptimeServer
	UptimeHost
	FPGARevision
	EchoNumber
)

// Return the number of classes supported by the server
func ClassCountMsg() ([]byte, error) {
	return PackMsg(TsReq{System, 0, uint8(ClassCount)})
}

// Return the number of instances for a given class
func InstanceCountMsg(classnum uint32) ([]byte, error) {
	return PackMsg(TsReq{System, 0, uint8(InstanceCount)}, classnum)
}

// Return the model-id of the CPU board
func ModelIdMsg() ([]byte, error) {
	return PackMsg(TsReq{System, 0, uint8(ModelId)})
}

// Return the model-id of the base-board
func BaseBoardIdMsg() ([]byte, error) {
	return PackMsg(TsReq{System, 0, uint8(BaseBoardId)})
}

// Lookup a DIO number by its name
func MapLookupMsg(name string) ([]byte, error) {
	n := uint32(len(name))
	return PackMsg(TsReq{System, 0, uint8(MapLookup)}, n, []byte(name))
}
