package ast

type NodeType string

const (
	DocumentType NodeType = "Document"

	// variable
	VarType        NodeType = "Variable"
	IdentifierType NodeType = "VariableName"

	// values
	ValueType    NodeType = "Value"
	LiteralType  NodeType = "LiteralValue"
	RefereceType NodeType = "VarReference"

	// Request
	RequestType   NodeType = "Request"
	MethodType    NodeType = "Method"
	UrlType       NodeType = "URL"
	ProtocolVType NodeType = "HTTPVersion"

	// headers
	HeaderType      NodeType = "Header"
	HeaderKeyType   NodeType = "HeaderKey"
	HeaderValueType NodeType = "HeaderValue"

	// payload
	PayloadType NodeType = "Payload"
)

type Node interface {
	Type() NodeType
	String() string
}

type Statement interface {
	Node
	StatementNode()
}
