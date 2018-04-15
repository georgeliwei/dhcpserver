package dhcpmsg

import "net"

type tdhcpMsgOption struct {
	optionCode byte
	optionLen  uint8
	optionVal  []byte
}
type TdhcpMsg struct {
	msgBuf []byte
	msgLen uint32
}

func (msg TdhcpMsg) GetOp() byte       { return msg.msgBuf[0] }
func (msg TdhcpMsg) GetHtype() byte    { return msg.msgBuf[1] }
func (msg TdhcpMsg) GetHlen() byte     { return msg.msgBuf[2] }
func (msg TdhcpMsg) GetHops() byte     { return msg.msgBuf[3] }
func (msg TdhcpMsg) GetXid() []byte    { return msg.msgBuf[4:8] }
func (msg TdhcpMsg) GetSecs() []byte   { return msg.msgBuf[8:10] }
func (msg TdhcpMsg) GetFlags() []byte  { return msg.msgBuf[10:12] }
func (msg TdhcpMsg) GetCiaddr() net.IP { return net.IP(msg.msgBuf[12:16]) }
func (msg TdhcpMsg) GetYiaddr() net.IP { return net.IP(msg.msgBuf[16:20]) }
func (msg TdhcpMsg) GetSiaddr() net.IP { return net.IP(msg.msgBuf[20:24]) }
func (msg TdhcpMsg) GetGiaddr() net.IP { return net.IP(msg.msgBuf[24:28]) }
func (msg TdhcpMsg) GetChaddr() net.HardwareAddr {
	hlen := msg.GetHlen()
	if hlen > 16 {
		hlen = 16
	}
	return net.HardwareAddr(msg.msgBuf[28 : 28+hlen])
}
func (msg TdhcpMsg) GetSname() []byte  { return msg.msgBuf[44:108] }
func (msg TdhcpMsg) GetFile() []byte   { return msg.msgBuf[108:236] }
func (msg TdhcpMsg) GetCookie() []byte { return msg.msgBuf[236:240] }
func (msg TdhcpMsg) GetOptions() []byte {
	if msg.msgLen > 240 {
		return msg.msgBuf[240:]
	}
	return nil
}
