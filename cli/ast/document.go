package ast

import "bytes"

type Document struct {
    Statements []Statement
}

func (d *Document) Type() NodeType {
    return DocumentType 
}

func (d *Document) String() string {
    var buf bytes.Buffer

    for _, st := range d.Statements {
        buf.WriteString(st.String())
        buf.WriteString("\n")
    }

    return buf.String()
}

