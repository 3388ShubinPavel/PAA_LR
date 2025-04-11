package main

import (
	"fmt"
	"sort"
)

type Node struct {
	children       map[rune]*Node
	suffixLink     *Node
	terminalLink   *Node
	patternIndices []int
	patternLength  int
	id             int
}

type AhoCorasick struct {
	root           *Node
	patterns       []string
	patternLengths []int
	nodeID         int
	verbose        bool
}

type Pair struct {
	Pos          int
	PatternIndex int
}

func NewAhoCorasick(patterns []string, verbose bool) *AhoCorasick {
	ac := &AhoCorasick{
		root: &Node{
			children: make(map[rune]*Node),
			id:       0,
		},
		patterns:       patterns,
		patternLengths: make([]int, len(patterns)),
		nodeID:         1,
		verbose:        verbose,
	}
	for i, p := range patterns {
		ac.patternLengths[i] = len(p)
	}
	ac.buildTrie()
	ac.buildSuffixAndTerminalLinks()
	return ac
}

func (ac *AhoCorasick) buildTrie() {
	if ac.verbose {
		fmt.Println("Building trie:")
	}
	for i, pattern := range ac.patterns {
		current := ac.root
		if ac.verbose {
			fmt.Printf("  Inserting pattern '%s' (ID %d)\n", pattern, i+1)
		}
		for _, c := range pattern {
			if _, exists := current.children[c]; !exists {
				if ac.verbose {
					fmt.Printf("    Creating new node %d for character '%c'\n", ac.nodeID, c)
				}
				current.children[c] = &Node{
					children: make(map[rune]*Node),
					id:       ac.nodeID,
				}
				ac.nodeID++
			}
			current = current.children[c]
			if ac.verbose {
				fmt.Printf("    Moved to node %d\n", current.id)
			}
		}
		current.patternIndices = append(current.patternIndices, i+1)
		current.patternLength = len(pattern)
		if ac.verbose {
			fmt.Printf("    Pattern %d ends at node %d\n", i+1, current.id)
		}
	}
}

func (ac *AhoCorasick) buildSuffixAndTerminalLinks() {
	if ac.verbose {
		fmt.Println("\nBuilding suffix links:")
	}
	queue := []*Node{}
	ac.root.suffixLink = ac.root
	ac.root.terminalLink = nil

	for _, child := range ac.root.children {
		child.suffixLink = ac.root
		queue = append(queue, child)
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for char, child := range current.children {
			temp := current.suffixLink
			for temp != ac.root && temp.children[char] == nil {
				temp = temp.suffixLink
			}
			if tempNode, exists := temp.children[char]; exists {
				child.suffixLink = tempNode
				if ac.verbose {
					fmt.Printf("  Set suffix link for node %d -> %d\n", child.id, tempNode.id)
				}
			} else {
				child.suffixLink = ac.root
				if ac.verbose {
					fmt.Printf("  Set suffix link for node %d -> root\n", child.id)
				}
			}
			queue = append(queue, child)
		}

		if current.suffixLink != nil && len(current.suffixLink.patternIndices) == 0 {
			current.terminalLink = current.suffixLink.terminalLink
		} else {
			current.terminalLink = current.suffixLink
		}
		if ac.verbose {
			if current.terminalLink != nil {
				fmt.Printf("  Terminal link for node %d -> %d\n", current.id, current.terminalLink.id)
			} else {
				fmt.Printf("  Terminal link for node %d -> nil\n", current.id)
			}
		}
	}
}

func (ac *AhoCorasick) PrintAutomaton() {
	fmt.Println("\nAutomaton structure:")
	queue := []*Node{ac.root}
	visited := make(map[int]bool)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if visited[node.id] {
			continue
		}
		visited[node.id] = true

		fmt.Printf("Node %d:\n", node.id)
		fmt.Printf("  Suffix link: %d\n", node.suffixLink.id)
		if node.terminalLink != nil {
			fmt.Printf("  Terminal link: %d\n", node.terminalLink.id)
		} else {
			fmt.Printf("  Terminal link: nil\n")
		}
		if len(node.patternIndices) > 0 {
			fmt.Printf("  Patterns: %v\n", node.patternIndices)
		}

		for char, child := range node.children {
			fmt.Printf("  Transition '%c' -> %d\n", char, child.id)
			queue = append(queue, child)
		}
	}
}

func (ac *AhoCorasick) Search(text string) []Pair {
	var results []Pair
	current := ac.root

	if ac.verbose {
		fmt.Printf("\nSearch process:\n")
	}

	for i := 0; i < len(text); i++ {
		c := rune(text[i])
		if ac.verbose {
			fmt.Printf("\nProcessing character '%c' at position %d\n", c, i+1)
			fmt.Printf("Current node: %d\n", current.id)
		}

		for current != ac.root && current.children[c] == nil {
			if ac.verbose {
				fmt.Printf("  Following suffix link %d -> %d\n", current.id, current.suffixLink.id)
			}
			current = current.suffixLink
		}

		if nextNode, exists := current.children[c]; exists {
			current = nextNode
			if ac.verbose {
				fmt.Printf("  Moved to node %d\n", current.id)
			}
		} else {
			current = ac.root
			if ac.verbose {
				fmt.Printf("  Returned to root\n")
			}
		}

		temp := current
		for temp != nil {
			if len(temp.patternIndices) > 0 {
				startPos := i - temp.patternLength + 2
				if ac.verbose {
					fmt.Printf("  Found pattern(s) %v at position %d\n", temp.patternIndices, startPos)
				}
				for _, idx := range temp.patternIndices {
					results = append(results, Pair{startPos, idx})
				}
			}
			temp = temp.terminalLink
			if temp != nil && ac.verbose {
				fmt.Printf("  Following terminal link to node %d\n", temp.id)
			}
		}
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Pos == results[j].Pos {
			return results[i].PatternIndex < results[j].PatternIndex
		}
		return results[i].Pos < results[j].Pos
	})

	return results
}
