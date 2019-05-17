package dom

import "testing"

func TestID(t *testing.T) {
	a := map[string]string{"id": "#title", "class": ".footer .navbar"}
	n := NewNodeElem("test", a, make([]Node, 0))

	n.NodeType.ID()
}
