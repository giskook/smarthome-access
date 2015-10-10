package sha

import (
	"github.com/giskook/smarthome-access/pb"
	"github.com/golang/protobuf/proto"
)

type NsqDeviceListPacket struct {
	GatewayID    uint64
	SerialNumber uint32
}

func (p *NsqDeviceListPacket) Serialize() []byte {
	g := NewGatewayHub().GetGateway(p.GatewayID)
	para := []*Report.Command_Param{
		&Report.Command_Param{
			Type:    Report.Command_Param_STRING,
			Strpara: g.Name,
		},

		&Report.Command_Param{
			Type:  Report.Command_Param_UINT16,
			Npara: uint64(g.Devicecount),
		},
	}

	if g.Devicecount > 0 {
		for i := uint16(0); i < g.Devicecount; i++ {
			para = append(para, &Report.Command_Param{
				Type:  Report.Command_Param_UINT64,
				Npara: g.Devicelist[i].Oid,
			})
			para = append(para, &Report.Command_Param{
				Type:  Report.Command_Param_UINT8,
				Npara: uint64(g.Devicelist[i].Type),
			})
			para = append(para, &Report.Command_Param{
				Type:  Report.Command_Param_UINT16,
				Npara: uint64(g.Devicelist[i].Company),
			})

			para = append(para, &Report.Command_Param{
				Type:  Report.Command_Param_UINT8,
				Npara: uint64(g.Devicelist[i].Status),
			})
			para = append(para, &Report.Command_Param{
				Type:    Report.Command_Param_STRING,
				Strpara: g.Devicelist[i].Name,
			})

		}
	}
	command := &Report.Command{
		Type:  Report.Command_CMT_REPDEVICELIST,
		Paras: para,
	}
	devicelist := &Report.ControlReport{
		Tid:          p.GatewayID,
		SerialNumber: p.SerialNumber,
		Command:      command,
	}

	data, _ := proto.Marshal(devicelist)

	return data
	//	var buf []byte
	//	buf = append(buf, 0xCE)
	//	buf = append(buf, 0x00)
	//	buf = append(buf, 0x0B)
	//	buf = append(buf, 0x80)
	//	buf = append(buf, 0x03)
	//	gatewayid := make([]byte, 8)
	//	binary.BigEndian.PutUint64(gatewayid, p.GatewayID)
	//	buf = append(buf, gatewayid[2:]...)
	//	buf = append(buf, CheckSum(buf, uint16(len(buf))))
	//	buf = append(buf, 0xCE)
	//
	//	return buf
}

func ParseNsqDeviceList(gatewayid uint64, serialnum uint32) *NsqDeviceListPacket {
	return &NsqDeviceListPacket{
		GatewayID:    gatewayid,
		SerialNumber: serialnum,
	}
}