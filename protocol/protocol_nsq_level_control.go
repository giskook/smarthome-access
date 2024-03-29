package protocol

import (
	"bytes"
	"github.com/giskook/smarthome-access/base"
	"github.com/giskook/smarthome-access/pb"
)

const CMD_LEVEL_CONTROL uint16 = 0x8011
const CMD_LEVEL_CONTROL_LEN uint16 = 0x0018
const CMD_LEVEL_PAUSE uint8 = 3

type Nsq_Level_Control_Packet struct {
	SerialNum       uint32
	DeviceID        uint64
	Endpoint        uint8
	CommandID       uint8 // constant value 00
	Level           uint8
	TransactionTime uint16 // constant value 0xffff
}

func (p *Nsq_Level_Control_Packet) Serialize() []byte {
	var writer bytes.Buffer
	writer.WriteByte(STARTFLAG)
	base.WriteWord(&writer, 0)
	base.WriteWord(&writer, CMD_LEVEL_CONTROL)
	base.WriteDWord(&writer, p.SerialNum)
	base.WriteQuaWord(&writer, p.DeviceID)
	writer.WriteByte(byte(p.Endpoint))
	writer.WriteByte(byte(p.CommandID))
	if p.CommandID != CMD_LEVEL_PAUSE {
		writer.WriteByte(byte(p.Level))
		base.WriteWord(&writer, p.TransactionTime)
	} else {
		writer.WriteByte(0x03)
	}
	base.WriteLength(&writer)
	writer.WriteByte(CheckSum(writer.Bytes(), uint16(writer.Len())))
	writer.WriteByte(ENDFLAG)

	return writer.Bytes()
}

func Parse_NSQ_Level_Control(serialnum uint32, paras []*Report.Command_Param) *Nsq_Level_Control_Packet {
	deviceid := paras[0].Npara
	endpint := uint8(paras[1].Npara)
	action := uint8(paras[2].Npara)
	level := uint8(paras[3].Npara)

	return &Nsq_Level_Control_Packet{
		SerialNum:       serialnum,
		DeviceID:        deviceid,
		Endpoint:        endpint,
		Level:           level,
		CommandID:       action,
		TransactionTime: 0xffff,
	}
}
