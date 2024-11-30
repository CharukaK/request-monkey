package ast

type Identifier struct {
    Text string
}

func (i *Identifier) Type() NodeType {
    return IdentifierType 
}

func (i *Identifier) String() string {
    return i.Text
}


