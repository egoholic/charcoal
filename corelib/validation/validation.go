package validation

type Node struct {
	title    string
	messages []string
	children map[string]*Node
}

func NewNode(title string) *Node {
	return &Node{title, make([]string, 0), make(map[string]*Node)}
}

func (n *Node) IsValid() bool {
	mlen := len(n.messages)
	if mlen > 0 {
		return false
	}

	for _, node := range n.children {
		if !node.IsValid() {
			return false
		}
	}

	return true
}

func (n *Node) Title() string {
	return n.title
}

func (n *Node) Messages() []string {
	return n.messages
}

func (n *Node) Children() map[string]*Node {
	return n.children
}

func (n *Node) AddMessage(m string) {
	n.messages = append(n.messages, m)
}

func (n *Node) AddChild(ch *Node) {
	n.children[ch.Title()] = ch
}
