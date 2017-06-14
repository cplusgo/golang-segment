package golang_segment

type TireNode struct {
	value     int32
	children  map[int32]*TireNode
	frequency int32
}

func NewTireNode(value int32) *TireNode {
	ins := &TireNode{
		value:     value,
		children:  make(map[int32]*TireNode),
		frequency: 0,
	}
	return ins
}

func (this *TireNode) addFrequency() {
	this.frequency++
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

func (this *TireNode) getChild(value int32) *TireNode {
	return this.children[value]
}
