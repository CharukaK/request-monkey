package ast

type Payload struct {
    Text string
}

func (p *Payload) Type() NodeType {
    return PayloadType
}

func (p *Payload) String() string {
    return p.Text
}

