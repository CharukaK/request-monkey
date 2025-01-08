package ast

type Method struct {
    Text string
}

func (m *Method) Type() NodeType {
    return MethodType 
}

func (m *Method) String() string {
    return m.Text
}

