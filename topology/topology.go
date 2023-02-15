package topology

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"front-end/rpctest/rpctest"
)

const (
	FILENAME          = "/home/djp/mininet/bin/mytopo.py"
	GraphPartitioners = "/home/djp/GraphPartitioners/data.txt"
)

type element Node

type Edge struct {
	Start  element
	End    element
	Intf1  int
	Intf2  int
	Weight uint
}

type Node struct {
	Name   string
	Weight uint
}

// -------------fat-tree-------------
type FatTree struct {
	cores []Node
	pods  []Pod
}

type Pod struct {
	aggregations []Node
	grounds      []Ground
}

type Ground struct {
	access Node
	hosts  []Node
}

// -------------fat-tree-------------

// ------------spine-leaf------------
type SpineLeaf struct {
	Spines []Node
	Leaves []Leaf
}

type Leaf struct {
	Leaf  Node
	Hosts []Node
}

// ------------spine-leaf------------

type MyEmunet struct {
	Nodes      map[string]element
	Edges      map[string]Edge
	Intf_index map[string]int
	Topo_type  string
}

func new_emunet(topo_type string) MyEmunet {
	nodes := make(map[string]element)
	edges := make(map[string]Edge)
	intf_index := make(map[string]int)

	return MyEmunet{nodes, edges, intf_index, topo_type}
}

func (n *MyEmunet) addnode(name string, weight ...int) error {
	if (len(weight) != 1) && (len(weight) != 0) {
		return errors.New("weight bit must be 1 or no weight")
	} else {
		if _, ok := n.Nodes[name]; ok {
			return errors.New("node" + name + "exist")
		}
		w := 1
		if len(weight) == 1 {
			w = weight[0]
		}

		n.Nodes[name] = element(Node{
			Name:   name,
			Weight: uint(w),
		})

		n.Intf_index[name] = 0
	}

	return nil
}

func (n *MyEmunet) addedge(node1 string, node2 string, weight ...int) error {
	_, ok1 := n.Nodes[node1]
	_, ok2 := n.Nodes[node2]
	if !(ok1 && ok2) {
		fmt.Println(node1 + "and" + node2 + "are not both exist")
		return errors.New(node1 + "and" + node2 + "are not both exist")
	}

	if (len(weight) != 1) && (len(weight) != 0) {
		fmt.Println("weight bit must be 1 or no weight")
		return errors.New("weight bit must be 1 or no weight")
	} else {
		if _, ok := n.Nodes[node1+"-"+node2]; ok {
			fmt.Println("edge" + node1 + "-" + node2 + "exist")
			return errors.New("edge" + node1 + "-" + node2 + "exist")
		}

		w := 1
		if len(weight) == 1 {
			w = weight[0]
		}

		n.Edges[node1+"-"+node2] = Edge{
			Start: element(Node{
				Name:   node1,
				Weight: n.Nodes[node1].Weight,
			}),
			End: element(Node{
				Name:   node2,
				Weight: n.Nodes[node2].Weight,
			}),
			Intf1:  n.Intf_index[node1],
			Intf2:  n.Intf_index[node2],
			Weight: uint(w),
		}

		fmt.Println("add edge", node1+"-"+node2)

		/*n.edges[node2+"-"+node1] = Edge{
			start: element(Node{
				name:   node2,
				weight: n.nodes[node2].weight,
			}),
			end: element(Node{
				name:   node1,
				weight: n.nodes[node1].weight,
			}),
			intf1:  n.intf_index[node2],
			intf2:  n.intf_index[node1],
			weight: uint(w),
		}*/

		n.Intf_index[node1]++
		n.Intf_index[node2]++
	}

	return nil
}

func (net MyEmunet) transform() rpctest.Emunet {
	emunet := rpctest.Emunet{}
	for _, node := range net.Nodes {
		emunet.Pods = append(emunet.Pods, &rpctest.Pod{
			Name: node.Name,
		})
	}
	for _, edge := range net.Edges {
		emunet.Links = append(emunet.Links, &rpctest.Link{
			Name: edge.Start.Name + "-" + edge.End.Name,
			Node1: &rpctest.Pod{
				Name: edge.Start.Name,
			},
			Node2: &rpctest.Pod{
				Name: edge.End.Name,
			},
			Node1Inf: strconv.Itoa(edge.Intf1),
			Node2Inf: strconv.Itoa(edge.Intf2),
		})
	}
	emunet.Type = net.Topo_type
	return emunet
}

func makerange(min, max int) []uint {
	a := make([]uint, max-min)
	for i := range a {
		a[i] = uint(min) + uint(i)
	}
	return a
}

func Make_Topo(topology_string string) (rpctest.Emunet, error) {
	split_string := strings.Split(topology_string, ",")
	topotype := split_string[0]
	parameter := split_string[1:]
	if topotype == "minimal" {
		return Generate_Minimal_Topo(), nil
	} else if topotype == "test" {
		if err := args_judgement(topotype, parameter); err != nil {
			return rpctest.Emunet{}, err
		} else {
			return GenerateTestTopo(), nil
		}
	} else if topotype == "linear" {
		// TODO
		return rpctest.Emunet{}, nil
	} else if topotype == "tree" {
		if err := args_judgement(topotype, parameter); err != nil {
			return rpctest.Emunet{}, err
		} else {
			return Generate_Tree_Topo(parameter), nil
		}
	} else if topotype == "fat-tree" {
		if err := args_judgement(topotype, parameter); err != nil {
			return rpctest.Emunet{}, err
		} else {
			return Generate_Fat_Tree_Topo(parameter[0], false), nil
		}
	} else if topotype == "spine-leaf" {
		if err := args_judgement(topotype, parameter); err != nil {
			return rpctest.Emunet{}, err
		} else {
			return Generate_Spine_Leaf_Topo(parameter, false), nil
		}
	} else {
		return rpctest.Emunet{}, errors.New("the topology type input doesn't match topology library (minimal, test, linear, tree, or fat-tree)")
	}
}

func args_judgement(type_string string, args []string) error {
	var type_arg_number int
	if type_string == "minimal" || type_string == "test" {
		type_arg_number = 0
	} else if type_string == "tree" {
		type_arg_number = 2
	} else if type_string == "fat-tree" {
		type_arg_number = 1
	} else if type_string == "spine-leaf" {
		type_arg_number = 4
	}

	if len(args) != type_arg_number {
		return errors.New("the parameter of" + type_string + "topology must be" + strconv.Itoa(type_arg_number))
	}

	if type_string == "spine-leaf" {
		// the first parameter of spine-leaf is spine number
		// the next is leaf number, host number and how many host for a leaf switch
		arg2, _ := strconv.Atoi(args[1])
		arg3, _ := strconv.Atoi(args[2])
		arg4, _ := strconv.Atoi(args[3])
		if arg3 != arg2*arg4 {
			return errors.New("the arg2 times arg4 must equal to arg3")
		}
	}

	for _, arg := range args {
		if _, err := strconv.Atoi(arg); err != nil {
			return errors.New("the parameter of " + type_string + " topology must be integer number")
		} else {
			if type_string == "fat-tree" {
				k, _ := strconv.Atoi(arg)
				if k%2 != 0 || k <= 2 {
					return errors.New("the k of fat-tree topology must be even and greater than 2")
				}
			} else if type_string == "spine-leaf" {
				k, _ := strconv.Atoi(arg)
				if k <= 1 {
					return errors.New("the spine, leaf, host and host every leaf must be inteter and greater than 1")
				}
			} else if type_string == "tree" {
				k, _ := strconv.Atoi(arg)
				if k <= 1 {
					return errors.New("the m and n of tree topology must be greater than 1")
				}
			}
		}
	}

	return nil
}

func Generate_Spine_Leaf_Topo(args []string, random bool) rpctest.Emunet {
	var spine_leaf SpineLeaf
	mynet := new_emunet("spine-leaf")
	spine_number, _ := strconv.Atoi(args[0])
	leaf_number, _ := strconv.Atoi(args[1])
	//host_number, _ := strconv.Atoi(args[2])
	host_per_leaf, _ := strconv.Atoi(args[3])

	switch_count := 1
	var i, j int
	for i = 1; i <= spine_number; i++ {
		node := Node{
			Name:   "s" + strconv.Itoa(switch_count),
			Weight: uint(leaf_number),
		}
		spine_leaf.Spines = append(spine_leaf.Spines, node)
		mynet.addnode("s"+strconv.Itoa(switch_count), leaf_number)
		switch_count++
	}

	for i = 1; i <= leaf_number; i++ {
		var leaf Leaf
		leaf.Leaf = Node{
			Name:   "s" + strconv.Itoa(switch_count),
			Weight: uint(spine_number + host_per_leaf),
		}
		mynet.addnode("s"+strconv.Itoa(switch_count), spine_number+host_per_leaf)
		switch_count++
		host_count := 1
		for j = 1; j <= host_per_leaf; j++ {
			leaf.Hosts = append(leaf.Hosts, Node{
				Name:   "h" + strconv.Itoa(host_count) + leaf.Leaf.Name,
				Weight: 1,
			})
			mynet.addnode("h"+strconv.Itoa(host_count)+leaf.Leaf.Name, 1)
			host_count++
		}
		spine_leaf.Leaves = append(spine_leaf.Leaves, leaf)
	}

	for _, spine := range spine_leaf.Spines {
		for _, leaf := range spine_leaf.Leaves {
			//fmt.Println(spine.Name + "-" + leaf.Leaf.Name)
			r := rand.Intn(5) + 1
			weight := uint(1)
			if random {
				weight = uint(r)
			}
			mynet.addedge(spine.Name, leaf.Leaf.Name, int(weight))
		}
	}

	for _, leaf := range spine_leaf.Leaves {
		for _, host := range leaf.Hosts {
			r := rand.Intn(5) + 1
			weight := uint(1)
			if random {
				weight = uint(r)
			}
			mynet.addedge(leaf.Leaf.Name, host.Name, int(weight))
		}
	}

	//fmt.Println(mynet)
	return mynet.transform()
}

func Generate_Fat_Tree_Topo(arg string, random bool) rpctest.Emunet {
	n, _ := strconv.Atoi(arg)
	var fat_tree FatTree

	mynet := new_emunet("fat-tree")
	switch_count := 1
	var i int
	for i = 1; i <= n*n/4; i++ {
		node := Node{
			Name:   "s" + strconv.Itoa(switch_count),
			Weight: uint(n),
		}
		fat_tree.cores = append(fat_tree.cores, node)
		mynet.addnode("s"+strconv.Itoa(switch_count), n)
		switch_count++
	}
	for i = 1; i <= n; i++ {
		var pod Pod
		for range makerange(1, n/2+1) {
			node := Node{
				Name:   "s" + strconv.Itoa(switch_count),
				Weight: uint(n),
			}
			pod.aggregations = append(pod.aggregations, node)
			mynet.addnode("s"+strconv.Itoa(switch_count), n)
			switch_count++

			var ground Ground
			node = Node{
				Name:   "s" + strconv.Itoa(switch_count),
				Weight: uint(n),
			}
			ground.access = node
			mynet.addnode("s"+strconv.Itoa(switch_count), n)
			switch_count++

			host_count := 1
			for range makerange(1, n/2+1) {
				node = Node{
					Name:   "h" + strconv.Itoa(host_count) + ground.access.Name,
					Weight: 1,
				}
				ground.hosts = append(ground.hosts, node)
				mynet.addnode("h"+strconv.Itoa(host_count)+ground.access.Name, 1)
				host_count++

			}
			pod.grounds = append(pod.grounds, ground)
		}
		fat_tree.pods = append(fat_tree.pods, pod)
	}

	for i, core := range fat_tree.cores {
		aggregation_count := i * 2 / n
		for _, pod := range fat_tree.pods {
			r := rand.Intn(5) + 1
			weight := uint(1)
			if random {
				weight = uint(r)
			}
			mynet.addedge(core.Name, pod.aggregations[aggregation_count].Name, int(weight))
		}
	}

	for _, pod := range fat_tree.pods {
		for _, aggregation := range pod.aggregations {
			for _, ground := range pod.grounds {
				r := rand.Intn(5) + 1
				weight := uint(1)
				if random {
					weight = uint(r)
				}
				mynet.addedge(aggregation.Name, ground.access.Name, int(weight))
			}
		}
	}

	for _, pod := range fat_tree.pods {
		for _, ground := range pod.grounds {
			for _, host := range ground.hosts {
				r := rand.Intn(5) + 1
				weight := uint(1)
				if random {
					weight = uint(r)
				}
				mynet.addedge(ground.access.Name, host.Name, int(weight))
			}
		}
	}

	return mynet.transform()
}

func Generate_Tree_Topo(args []string) rpctest.Emunet {
	mynet := new_emunet("tree")
	m, _ := strconv.Atoi(args[0])
	n, _ := strconv.Atoi(args[1])
	nodes := make([]element, 0)
	cut_point := (1 - int(math.Pow(float64(n), float64(m)))) / (1 - n)      // 最后一个switch节点的序号，从1开始
	total_number := (1 - int(math.Pow(float64(n), float64(m+1)))) / (1 - n) // 总节点数
	for _, i := range makerange(1, cut_point+1) {
		mynet.addnode("s"+strconv.Itoa(int(i)), 1)
		nodes = append(nodes, element(Node{
			Name:   "s" + strconv.Itoa(int(i)),
			Weight: 1,
		}))
	}
	for _, i := range makerange(cut_point+1, total_number+1) {
		parent_index := (int(i) + n - 2) / n
		mynet.addnode("h"+strconv.Itoa((int(i)+n-2)%n+1)+nodes[parent_index-1].Name, 1)
		nodes = append(nodes, element(Node{
			Name:   "h" + strconv.Itoa((int(i)+n-2)%n+1) + nodes[parent_index-1].Name,
			Weight: 1,
		}))
	}

	for _, i := range makerange(2, total_number+1) {
		mynet.addedge(nodes[(int(i)+n-2)/n-1].Name, nodes[int(i)-1].Name, 1)
		mynet.addedge(nodes[int(i)-1].Name, nodes[(int(i)+n-2)/n-1].Name, 1)
	}

	return mynet.transform()
}

func Generate_Minimal_Topo() rpctest.Emunet {
	mynet := new_emunet("minimal")
	mynet.addnode("h1s1")
	mynet.addnode("h2s1")
	mynet.addnode("s1")

	mynet.addedge("h1s1", "s1")
	mynet.addedge("h2s1", "s1")
	return mynet.transform()
}

func GenerateTestTopo() rpctest.Emunet {
	mynet := new_emunet("test")
	mynet.addnode("s1")
	mynet.addnode("h1s1")
	mynet.addnode("h2s1")

	mynet.addedge("h1s1", "s1")
	mynet.addedge("s1", "h2s1")

	return mynet.transform()
}
