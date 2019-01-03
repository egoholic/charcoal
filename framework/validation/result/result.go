package result

type Node struct {
	domainName string
	issues     map[string][]string
	children   map[string]*Node
}

func (n *Node) DomainName() string          { return n.domainName }
func (n *Node) Issues() map[string][]string { return n.issues }
