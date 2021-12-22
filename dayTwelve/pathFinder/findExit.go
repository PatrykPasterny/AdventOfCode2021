package pathfinder

import (
	"fmt"
	"strings"
	"unicode"
)

type GraphNode struct {
	Value    string
	Children []*GraphNode
}

func FindAllExitPathsForSingleSmallCaveEntrance(data []string) int {
	pathsMap := createPathsGraph(data)
	root := pathsMap["start"]
	visitedSmallCases := make(map[string]int, 0)

	return countAllPathsWithZeroDoubleSmallCave(root, visitedSmallCases)
}

func FindAllExitPathsForDoubleSmallCaveEntrance(data []string) int {
	pathsMap := createPathsGraph(data)
	root := pathsMap["start"]
	visitedSmallCases := make(map[string]int, 0)

	return countAllPathsWithOneDoubleSmallCave(root, visitedSmallCases, 2, false)
}

func createPathsGraph(data []string) (pathsMap map[string]*GraphNode) {
	pathsMap = make(map[string]*GraphNode, 0)
	for _, val := range data {
		nodes := strings.Split(val, "-")
		if nodeKey, ok := pathsMap[nodes[0]]; ok {
			if nodeVal, ok := pathsMap[nodes[1]]; ok {
				nodeKey.Children = append(nodeKey.Children, nodeVal)
				nodeVal.Children = append(nodeVal.Children, nodeKey)
			} else {
				childNode := GraphNode{Value: nodes[1], Children: make([]*GraphNode, 0)}
				nodeKey.Children = append(nodeKey.Children, &childNode)
				childNode.Children = append(childNode.Children, nodeKey)
				pathsMap[nodes[1]] = &childNode
			}
		} else {
			if nodeVal, ok := pathsMap[nodes[1]]; ok {
				childNode := GraphNode{Value: nodes[0], Children: make([]*GraphNode, 0)}
				nodeVal.Children = append(nodeVal.Children, &childNode)
				childNode.Children = append(childNode.Children, nodeVal)
				pathsMap[nodes[0]] = &childNode
			} else {
				childNode := GraphNode{Value: nodes[1], Children: make([]*GraphNode, 0)}
				pathsMap[nodes[0]] = &GraphNode{Value: nodes[0], Children: []*GraphNode{&childNode}}
				childNode.Children = append(childNode.Children, pathsMap[nodes[0]])
				pathsMap[nodes[1]] = &childNode
			}
		}
	}

	return
}
func countAllPathsWithZeroDoubleSmallCave(currentNode *GraphNode, visited map[string]int) int {
	if _, ok := visited[currentNode.Value]; ok {
		return 0
	}
	if currentNode.Value == "start" {
		visited[currentNode.Value] = 1
	} else {
		if unicode.IsLower(rune(currentNode.Value[0])) {
			if _, ok := visited[currentNode.Value]; !ok {
				visited[currentNode.Value] = 1
			}
		}
	}
	if currentNode.Value == "end" {
		return 1
	}

	result := 0
	for i := range currentNode.Children {
		if _, ok := visited[currentNode.Children[i].Value]; !ok {
			visitedCopy := make(map[string]int, 0)
			for k, v := range visited {
				visitedCopy[k] = v
			}

			result += countAllPathsWithZeroDoubleSmallCave(currentNode.Children[i], visitedCopy)
		}
	}

	return result
}

func countAllPathsWithOneDoubleSmallCave(currentNode *GraphNode, visited map[string]int, numberOfSingleSmallCavesVisits int, wasNTimesInSmallCave bool) int {
	if _, ok := visited[currentNode.Value]; wasNTimesInSmallCave && ok {
		return 0
	}
	if currentNode.Value == "start" {
		visited[currentNode.Value] = 1
	} else {
		if unicode.IsLower(rune(currentNode.Value[0])) {
			if _, ok := visited[currentNode.Value]; ok {
				wasNTimesInSmallCave = true
				visited[currentNode.Value] += 1
			} else {
				visited[currentNode.Value] = 1
			}
		}
	}
	if currentNode.Value == "end" {
		return 1
	}

	result := 0
	for i := range currentNode.Children {
		if _, ok := visited[currentNode.Children[i].Value]; !ok || (!wasNTimesInSmallCave) && currentNode.Children[i].Value != "start" {
			visitedCopy := make(map[string]int, 0)
			for k, v := range visited {
				visitedCopy[k] = v
			}

			result += countAllPathsWithOneDoubleSmallCave(currentNode.Children[i], visitedCopy, numberOfSingleSmallCavesVisits, wasNTimesInSmallCave)
		}
	}

	return result
}

func printMap(pathsMap map[string]*GraphNode) {
	for _, v := range pathsMap {
		fmt.Println(v.Value)
		for _, ch := range v.Children {
			fmt.Print(*ch)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
