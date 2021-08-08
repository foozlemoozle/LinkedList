package LinkedList

type GraphNode struct {
	Paths IList
	Data  interface{}
}

type Path struct {
	From *GraphNode
	To   *GraphNode
	Data interface{}
}

type Graph struct {
	Head *GraphNode
}

func MakeGraph() *Graph {
	return &Graph{Head: &GraphNode{Paths: List()}}
}

func MakeGraphNode(data interface{}) *GraphNode {
	return &GraphNode{
		Data:  data,
		Paths: List(),
	}
}

func CreateTwoDirectionPath(a *GraphNode, b *GraphNode, meta interface{}) (*Path, *Path) {
	return CreatePath(a, b, meta), CreatePath(b, a, meta)
}

func CreatePath(from *GraphNode, to *GraphNode, meta interface{}) *Path {
	path := &Path{
		From: from,
		To:   to,
		Data: meta,
	}

	from.Paths.Enqueue(path)
	return path
}

func (graph *Graph) Find(data interface{}) (ISet, ISet, bool) {
	return graph.Head.find(data, List())
}

func (node *GraphNode) find(data interface{}, searched IList) (IList, IList, bool) {
	searched.Enqueue(node)

	if node.Data == data {
		//this node is the desired node, return this node data and empty path data (filled down the callstack)
		dataSet := List()
		dataSet.Enqueue(data)
		return dataSet, List(), true
	}

	iter := node.Paths.Iterator()
	for val, ok := iter.Current(); ok; val, ok = iter.MoveNext() {
		path := val.(*Path)

		//don't check a path that links to a node we've already searched
		if searched.Find(path.To) != -1 {
			continue
		}

		nodeData, pathData, ok := path.To.find(data, searched)
		//node found, append this data to found data, and return
		if ok {
			nodeData.Enqueue(node.Data)
			pathData.Enqueue(val.(*Path).Data)
			return nodeData, pathData, true
		}
	}

	//node not found, return false
	return nil, nil, false
}

func (graph *Graph) HasLoop() bool {
	return graph.Head.hasLoop(List())
}

func (node *GraphNode) hasLoop(searched IList) bool {
	if searched.Find(node) != -1 {
		return true
	}

	searched.Enqueue(node)

	iter := node.Paths.Iterator()
	for val, ok := iter.Current(); ok; val, ok = iter.MoveNext() {
		path := val.(*Path)

		if path.To.hasLoop(searched) {
			return true
		}
	}

	return false
}
