package ast

import "bytes"

type Request struct {
	Method      Method
	Url         Url
	ProtocolVer *ProtocolVer
	Headers     []Header
	Payload     *Payload
}

func (r *Request) StatementNode() {}

func (r *Request) Type() NodeType {
	return RequestType
}

func (r *Request) String() string {
	var buf bytes.Buffer

	buf.WriteString(r.Method.String())
	buf.WriteString(" ")
	buf.WriteString(r.Url.String())

	if r.ProtocolVer != nil {
		buf.WriteString(" ")
		buf.WriteString(r.ProtocolVer.String())
	}

	buf.WriteString("\n")

	for _, hv := range r.Headers {
		buf.WriteString(hv.String())
		buf.WriteString("\n")
	}

	buf.WriteString("\n")

	if r.Payload != nil {
		buf.WriteString(r.Payload.String())
	}

	return buf.String()
}
