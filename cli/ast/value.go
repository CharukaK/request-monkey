package ast

import "bytes"

// Values
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

// Value segment
type ValuePart interface {
    Node
    valuePartType()
}

// literals
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

// references
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

