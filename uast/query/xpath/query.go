package xpath

import (
	"fmt"
	"strings"

	"github.com/antchfx/xpath"

	"gopkg.in/bblfsh/sdk.v2/uast"
	"gopkg.in/bblfsh/sdk.v2/uast/nodes"
)

var _ xpath.NodeNavigator = &nodeNavigator{}

// newNavigator creates a new xpath.nodeNavigator for the specified html.node.
func newNavigator(root nodes.External) *nodeNavigator {
	n := &node{n: root, typ: rootNode}
	return &nodeNavigator{root: n, cur: n}
}

// A nodeType is the type of a node.
type nodeType uint

const (
	// rootNode is a document object that, as the root of the document tree,
	// provides access to the entire XML document.
	rootNode nodeType = iota
	// objectNode is an element.
	objectNode
	fieldNode
	// valueNode is the text content of a node.
	valueNode
)

type node struct {
	typ nodeType

	n    nodes.External
	kind nodes.Kind
	obj  nodes.ExternalObject

	tag  [2]string
	sub  []*node
	par  *node
	pari int // index in parent's sub array
}

// nodeNavigator is for navigating JSON document.
type nodeNavigator struct {
	root, cur *node
}

func (a *nodeNavigator) Current() nodes.External {
	return a.cur.n
}

func (a *nodeNavigator) NodeType() xpath.NodeType {
	switch a.cur.typ {
	case valueNode:
		return xpath.TextNode
	case rootNode:
		return xpath.RootNode
	case objectNode, fieldNode:
		return xpath.ElementNode
	default:
		panic(fmt.Sprintf("unknown node type %v", a.cur.typ))
	}
}

func (a *nodeNavigator) LocalName() string {
	return a.cur.tag[1]
}

func (a *nodeNavigator) Prefix() string {
	return a.cur.tag[0]
}

func (a *nodeNavigator) Value() string {
	switch a.cur.typ {
	case valueNode:
		return fmt.Sprint(a.cur.n.Value())
	}
	return ""
}

func (a *nodeNavigator) Copy() xpath.NodeNavigator {
	n := *a
	return &n
}

func (a *nodeNavigator) MoveToRoot() {
	a.cur = a.root
}

func (a *nodeNavigator) MoveToParent() bool {
	n := a.cur.par
	if n == nil {
		return false
	}
	a.cur = n
	return true
}

func (x *nodeNavigator) MoveToNextAttribute() bool {
	return false
}

func toNode(n nodes.External, field string) *node {
	if n == nil {
		// TODO: what about nil attributes?
		return nil
	}
	nd := &node{n: n, kind: n.Kind()}

	wrap := func(nd *node) *node {
		if field == "" {
			return nd
		}
		// wrap node into field-node
		f := &node{
			n: nd.n, kind: nd.kind,
			typ: fieldNode, tag: [2]string{"", field},
			sub: []*node{nd},
		}
		nd.par = f
		nd.pari = 0
		return f
	}

	switch nd.kind {
	case nodes.KindNil:
		return nil // TODO
	case nodes.KindObject:
		if typ := uast.TypeOf(n); typ != "" {
			if i := strings.Index(typ, ":"); i >= 0 {
				nd.tag = [2]string{typ[:i], typ[i+1:]}
			} else {
				nd.tag = [2]string{"", typ}
			}
		}
		nd.obj, _ = nd.n.(nodes.ExternalObject)
		nd.typ = objectNode
		return wrap(nd)
	case nodes.KindArray:
		arr, _ := nd.n.(nodes.ExternalArray)
		// array == sub nodes of this field
		f := &node{
			n: nd.n, kind: nd.kind,
			typ: fieldNode, tag: [2]string{"", field},
		}
		if arr == nil {
			f.sub = []*node{}
			return f
		}
		sz := arr.Size()
		f.sub = make([]*node, 0, sz)
		for i := 0; i < sz; i++ {
			v := arr.ValueAt(i)
			s := toNode(v, "")
			s.par = f
			s.pari = i
			f.sub = append(f.sub, s)
		}
		return f
	default:
		// value
		nd.typ = valueNode
		return wrap(nd)
	}
}

func (a *nodeNavigator) MoveToChild() bool {
	switch a.cur.typ {
	case rootNode:
		// return the same node, but without the root type
		n := toNode(a.cur.n, "")
		if n == nil {
			return false
		}
		n.par = a.cur
		a.cur = n
		return true
	case objectNode:
		// node is an object, children are wrapped into a tag with the name = field
		if a.cur.obj == nil {
			return false
		}
		cur := a.cur
		if cur.sub == nil {
			// project fields
			obj := cur.obj
			keys := obj.Keys()
			cur.sub = make([]*node, 0, len(keys))
			for i, k := range keys {
				v, ok := obj.ValueAt(k)
				if !ok {
					continue
				}
				vn := toNode(v, k)
				vn.par = cur
				vn.pari = i
				cur.sub = append(cur.sub, vn)
			}
		}
		if len(cur.sub) == 0 {
			return false
		}
		a.cur = cur.sub[0]
		return true
	case fieldNode:
		n := a.cur.sub[0]
		if n == nil {
			return false
		}
		a.cur = n
		return true
	}
	return false
}

func (a *nodeNavigator) isSub() bool {
	return a.cur.par != nil && a.cur.pari < len(a.cur.par.sub)
}
func (a *nodeNavigator) MoveToFirst() bool {
	if a.isSub() {
		par := a.cur.par
		if n := par.sub[0]; n != nil {
			a.cur = n
		}
	}
	return true
}

func (a *nodeNavigator) MoveToNext() bool {
	if a.isSub() {
		par := a.cur.par
		if i := a.cur.pari + 1; i < len(par.sub) {
			a.cur = par.sub[i]
			return true
		}
	}
	return false
}

func (a *nodeNavigator) MoveToPrevious() bool {
	if a.isSub() {
		par := a.cur.par
		if i := a.cur.pari - 1; i >= 0 && i < len(par.sub) {
			a.cur = par.sub[i]
			return true
		}
	}
	return false
}

func (a *nodeNavigator) MoveTo(other xpath.NodeNavigator) bool {
	node, ok := other.(*nodeNavigator)
	if !ok || node.root != a.root {
		return false
	}
	a.cur = node.cur
	return true
}
