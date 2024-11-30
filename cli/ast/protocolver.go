package ast

type ProtocolVer struct {
    Text string
}

func (pv *ProtocolVer) Type() NodeType {
    return ProtocolVType
}

func (pv *ProtocolVer) String() string {
    return pv.Text
}

