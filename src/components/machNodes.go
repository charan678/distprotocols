package components

type MachNode interface{

}

type machNode struct {
	name string
}

func NewMachNode(name string) MachNode{
	return &machNode{name: name}
}


