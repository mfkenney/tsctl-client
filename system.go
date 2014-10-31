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

func ClassCountMsg() ([]byte, error) {
	return PackMsg(TsReq{System, 0, uint8(ClassCount)})
}

func InstanceCountMsg(classnum uint32) ([]byte, error) {
	return PackMsg(TsReq{System, 0, uint8(InstanceCount)}, classnum)
}

func ModelIdMsg() ([]byte, error) {
	return PackMsg(TsReq{System, 0, uint8(ModelId)})
}

func BaseBoardIdMsg() ([]byte, error) {
	return PackMsg(TsReq{System, 0, uint8(BaseBoardId)})
}

func MapLookupMsg(name string) ([]byte, error) {
	n := uint32(len(name))
	return PackMsg(TsReq{System, 0, uint8(MapLookup)}, n, []byte(name))
}
