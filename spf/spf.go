package main

import "sort"
import "fmt"

// Queue is a basic priority queue implementation, where the node with the
// lowest priority is kept as first element in the queue
type Queue struct {
	keys  []string
	nodes map[string]int
}

// Len is part of sort.Interface
func (q *Queue) Len() int {
	return len(q.keys)
}

// Swap is part of sort.Interface
func (q *Queue) Swap(i, j int) {
	q.keys[i], q.keys[j] = q.keys[j], q.keys[i]
}

// Less is part of sort.Interface
func (q *Queue) Less(i, j int) bool {
	a := q.keys[i]
	b := q.keys[j]

	return q.nodes[a] < q.nodes[b]
}

// Set updates or inserts a new key in the priority queue
func (q *Queue) Set(key string, priority int) {
	// inserts a new key if we don't have it already
	if _, ok := q.nodes[key]; !ok {
		q.keys = append(q.keys, key)
	}

	// set the priority for the key
	q.nodes[key] = priority

	// sort the keys array
	sort.Sort(q)
}

// Next removes the first element from the queue and retuns it's key and priority
func (q *Queue) Next() (key string, priority int) {
	// shift the key form the queue
	key, keys := q.keys[0], q.keys[1:]
	q.keys = keys

	priority = q.nodes[key]

	delete(q.nodes, key)

	return key, priority
}

// IsEmpty returns true when the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.keys) == 0
}

// Get returns the priority of a passed key
func (q *Queue) Get(key string) (priority int, ok bool) {
	priority, ok = q.nodes[key]
	return
}

// NewQueue creates a new empty priority queue
func NewQueue() *Queue {
	var q Queue
	q.nodes = make(map[string]int)
	return &q
}


type node struct {
    key  string
    cost int
}

// Graph is a rappresentation of how the points in our graph are connected
// between each other
type Graph map[string]map[string]int

// Path finds the shortest path between start and target, also returning the
// total cost of the found path.
func (g Graph) Path(start, target string) (path []string, cost int, err error) {
    if len(g) == 0 {
        err = fmt.Errorf("cannot find path in empty map")
        return
    }

    // ensure start and target are part of the graph
    if _, ok := g[start]; !ok {
        err = fmt.Errorf("cannot find start %v in graph", start)
        return
    }
    if _, ok := g[target]; !ok {
        err = fmt.Errorf("cannot find target %v in graph", target)
        return
    }

    explored := make(map[string]bool)   // set of nodes we already explored
    frontier := NewQueue()              // queue of the nodes to explore
    previous := make(map[string]string) // previously visited node

    // add starting point to the frontier as it'll be the first node visited
    frontier.Set(start, 0)

    // run until we visited every node in the frontier
    for !frontier.IsEmpty() {
        // get the node in the frontier with the lowest cost (or priority)
        aKey, aPriority := frontier.Next()
        n := node{aKey, aPriority}

        // when the node with the lowest cost in the frontier is target, we can
        // compute the cost and path and exit the loop
        if n.key == target {
            cost = n.cost

            nKey := n.key
            for nKey != start {
                path = append(path, nKey)
                nKey = previous[nKey]
            }
            
            break
        }

        // add the current node to the explored set
        explored[n.key] = true

        // loop all the neighboring nodes
        for nKey, nCost := range g[n.key] {
            // skip alreadt-explored nodes
            if explored[nKey] {
                continue
            }

            // if the node is not yet in the frontier add it with the cost
            if _, ok := frontier.Get(nKey); !ok {
                previous[nKey] = n.key
                frontier.Set(nKey, n.cost+nCost)
                continue
            }

            frontierCost, _ := frontier.Get(nKey)
            nodeCost := n.cost + nCost

            // only update the cost of this node in the frontier when
            // it's below what's currently set
            if nodeCost <= frontierCost {
                previous[nKey] = n.key
                frontier.Set(nKey, nodeCost)
            }
        }
    }

    // add the origin at the end of the path
    path = append(path, start)

    // reverse the path because it was populated
    // in reverse, form target to start
    for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
        path[i], path[j] = path[j], path[i]
    }

    return
}

func main() {

	g := Graph{
        
        "a":{"b":20,"c":80},
        "b":{"a": 20,"d":80},
        "c": {"a": 80,"d":20},
        "d":{"b":80,"c":20},
    }
    //fmt.Println("Graph",g)
    path, cost, _ := g.Path("a","d")
    fmt.Printf("path: %v, cost: %v", path, cost)
    fmt.Println("\n")
}

