package golang_segment

type TireNode struct {
	value    int32
	children map[int32]*TireNode
}

func NewTireNode(value int32) *TireNode {
	ins := &TireNode{
		value:    value,
		children: make(map[int32]*TireNode),
	}
	return ins
}

func (this *TireNode) AddChild(value int32) *TireNode {
	node := this.children[value]
	if node == nil {
		node = &TireNode{
			value:    value,
			children: make(map[int32]*TireNode),
		}
		this.children[value] = node
	}
	return node
}

func (this *TireNode) getChild(value int32) *TireNode  {
	return this.children[value]
}
