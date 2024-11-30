package ast

import "bytes"

type Value struct {
	Parts []ValuePart
}

func (v *Value) Type() NodeType {
	return ValueType
}

func (v *Value) String() string {
    var buf bytes.Buffer

    for _, vp := range v.Parts {
        buf.WriteString(vp.String())
    }

    return buf.String()
}

type ValuePart interface {
    Node
    valuePartType()
}

type LiteralValue struct {
    Text string
}

func (lv *LiteralValue) valuePartType() {}

func (lv *LiteralValue) Type() NodeType {
    return LiteralType
}

func (lv *LiteralValue) String() string {
    return lv.Text
}


type ReferenceValue struct {
    Reference Identifier 
}

func (rv *ReferenceValue) valuePartType() {}

func (rv *ReferenceValue) Type() NodeType {
    return RefereceType 
}

func (rv *ReferenceValue) String() string {
    var buf bytes.Buffer

    buf.WriteString("{{")
    buf.WriteString(rv.Reference.String())
    buf.WriteString("}}")

    return buf.String() 
}

