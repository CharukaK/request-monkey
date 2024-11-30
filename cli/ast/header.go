package ast

import "bytes"

type HeaderKey struct {
    Text string
}

func (hk *HeaderKey) Type() NodeType {
    return HeaderKeyType 
}

func (hk *HeaderKey) String() string {
    return hk.Text 
}

type Header struct {
    Key HeaderKey
    ValueParts []ValuePart
}

func (h *Header) Type() NodeType {
    return HeaderType 
}

func (h *Header) String() string {
    var buf bytes.Buffer

    buf.WriteString(h.Key.String())
    buf.WriteString(": ")

    for _, vp := range h.ValueParts {
        buf.WriteString(vp.String())
    }

    return buf.String()
}

