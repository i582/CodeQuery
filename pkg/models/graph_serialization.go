package models

//
// var encodeVisited = map[*Node]bool{}
// var allNodes []*Node
//
// func (n *Node) GobEncode() ([]byte, error) {
// 	encodeNodeArray := func(nodes []*Node, encoder *gob.Encoder) error {
// 		var nodesMap []int64
// 		for _, node := range nodes {
// 			nodesMap = append(nodesMap, node.ID)
// 		}
//
// 		err := encoder.Encode(nodesMap)
// 		if err != nil {
// 			return err
// 		}
//
// 		return nil
// 	}
//
// 	if encodeVisited[n] {
// 		return nil, nil
// 	}
// 	encodeVisited[n] = true
//
// 	w := new(bytes.Buffer)
// 	encoder := gob.NewEncoder(w)
// 	err := encoder.Encode(n.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = encoder.Encode(n.Data)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	err = encodeNodeArray(n.Next, encoder)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = encodeNodeArray(n.Prev, encoder)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return w.Bytes(), nil
// }
//
// func (n *Node) GobDecode(buf []byte) error {
// 	decodeNodeArray := func(nodes *NodesIDs, decoder *gob.Decoder) error {
// 		var nodesMap []int
//
// 		err := decoder.Decode(&nodesMap)
// 		if err != nil {
// 			return err
// 		}
//
// 		for _, nodeIndex := range nodesMap {
// 			*nodes = append(*nodes, allNodes[nodeIndex])
// 		}
//
// 		return nil
// 	}
//
// 	r := bytes.NewBuffer(buf)
// 	decoder := gob.NewDecoder(r)
// 	err := decoder.Decode(&n.ID)
// 	if err != nil {
// 		return err
// 	}
// 	err = decoder.Decode(&n.Data)
// 	if err != nil {
// 		return err
// 	}
//
// 	err = decodeNodeArray(&n.Prev, decoder)
// 	if err != nil {
// 		return err
// 	}
//
// 	err = decodeNodeArray(&n.Next, decoder)
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }
//
// func (g *Graph) GobEncode() ([]byte, error) {
// 	w := new(bytes.Buffer)
// 	encoder := gob.NewEncoder(w)
//
// 	err := encoder.Encode(g.NodesIDs)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	allNodes = g.NodesIDs
//
// 	graphMap := make(map[int64]NodesIDs, len(g.Graph))
//
// 	for node, nodes := range g.Graph {
// 		graphMap[node.ID] = nodes
// 	}
//
// 	err = encoder.Encode(graphMap)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return w.Bytes(), nil
// }
//
// func (g *Graph) GobDecode(buf []byte) error {
// 	r := bytes.NewBuffer(buf)
// 	decoder := gob.NewDecoder(r)
// 	err := decoder.Decode(&g.NodesIDs)
// 	if err != nil {
// 		return err
// 	}
//
// 	graphMap := make(map[int64]NodesIDs)
//
// 	err = decoder.Decode(&graphMap)
// 	if err != nil {
// 		return err
// 	}
//
// 	g.Graph = make(map[*Node]NodesIDs, len(graphMap))
//
// 	for index, nodes := range graphMap {
// 		g.Graph[g.NodesIDs[index]] = nodes
// 	}
//
// 	return nil
// }
