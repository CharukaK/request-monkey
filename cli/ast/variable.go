package ast

import "bytes"

type Variable struct {
    Name Identifier
    Value Value
}

func (v *Variable) StatementNode() {}

func (v *Variable) Type() NodeType {
    return VarType
}

func (v *Variable) String() string {
    var buf bytes.Buffer

    buf.WriteString("@")
    buf.WriteString(v.Name.String())
    buf.WriteString("=")
    buf.WriteString(v.Value.String())

    return buf.String()
}

