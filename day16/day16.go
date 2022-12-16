package day16

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	From *Node
	To   *Node
	Time int
}

type Node struct {
	Name     string
	FlowRate int
	Edges    map[string]Edge
}

func loadData(input string) map[string]*Node {
	readFile, _ := os.Open(input)
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	measurements := make(map[string]*Node)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = strings.Replace(line, "Valve ", "", 1)
		line = strings.Replace(line, "has flow rate=", "", 1)
		line = strings.Replace(line, "; tunnels lead to valves", "", 1)
		line = strings.Replace(line, "; tunnel leads to valve", "", 1)
		line = strings.ReplaceAll(line, ",", "")

		data := strings.Split(line, " ")
		flowRate, _ := strconv.Atoi(data[1])

		currentNode := &Node{Name: data[0]}
		if node, ok := measurements[data[0]]; ok {
			currentNode = node
		}

		currentNode.FlowRate = flowRate
		currentNode.Edges = make(map[string]Edge)

		for _, nodeName := range data[2:] {
			if node, ok := measurements[nodeName]; ok {
				currentNode.Edges[nodeName] = Edge{From: currentNode, To: node, Time: 1}
			} else {
				newNode := &Node{Name: nodeName}

				measurements[nodeName] = newNode
				currentNode.Edges[nodeName] = Edge{From: currentNode, To: newNode, Time: 1}
			}
		}

		measurements[currentNode.Name] = currentNode
	}

	return measurements
}

func reduceNodes(nodes map[string]*Node) map[string]*Node {
	for _, node := range nodes {
		if node.FlowRate != 0 || node.Name == "AA" {
			continue
		}

		for _, edge := range node.Edges {
			for _, edge2 := range node.Edges {
				if edge == edge2 {
					continue
				}

				edge.To.Edges[edge2.To.Name] = Edge{From: edge.To, To: edge2.To, Time: edge.Time + edge2.Time}
				edge2.To.Edges[edge.To.Name] = Edge{From: edge2.To, To: edge.To, Time: edge.Time + edge2.Time}
				delete(edge.To.Edges, node.Name)
				delete(edge2.To.Edges, node.Name)
			}
		}

		delete(nodes, node.Name)
	}

	return nodes
}

func Part1() int {
	nodes := loadData("day16/example.txt")

	nodes = reduceNodes(nodes)

	fmt.Println(nodes)

	for _, node := range nodes {
		fmt.Println(node.Name, node.FlowRate, node.Edges)
	}

	return 0
}
