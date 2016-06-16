package protocol

import (
	"bytes"
	"github.com/giskook/smarthome-access/base"
	"github.com/giskook/smarthome-access/pb"
)

const CMD_QUERY_ATTR uint16 = 0x800b
const CMD_QUERY_ATTR_LEN uint16 = 0x14

type Nsq_Query_Attr_Packet struct {
	SerialNum uint32
	DeviceID  uint64
	Endpoint  uint8
}

func (p *Nsq_Query_Attr_Packet) Serialize() []byte {
	var writer bytes.Buffer
	writer.WriteByte(STARTFLAG)
	base.WriteWord(&writer, CMD_QUERY_ATTR_LEN)
	base.WriteWord(&writer, CMD_QUERY_ATTR)
	base.WriteDWord(&writer, p.SerialNum)
	base.WriteQuaWord(&writer, p.DeviceID)
	writer.WriteByte(p.Endpoint)
	writer.WriteByte(CheckSum(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(ENDFLAG)

	return writer.Bytes()
}

func Parse_NSQ_Query_Attr(serialnum uint32, paras []*Report.Command_Param) *Nsq_Query_Attr_Packet {
	deviceid := paras[0].Npara
	endpint := uint8(paras[1].Npara)

	return &Nsq_Query_Attr_Packet{
		SerialNum: serialnum,
		DeviceID:  deviceid,
		Endpoint:  endpint,
	}

}
