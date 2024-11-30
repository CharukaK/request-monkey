package ast

import "bytes"

type Url struct {
   Parts []ValuePart 
}

func (v *Url) Type() NodeType {
	return UrlType
}

func (v *Url) String() string {
    var buf bytes.Buffer

    for _, vp := range v.Parts {
        buf.WriteString(vp.String())
    }

    return buf.String()
}

