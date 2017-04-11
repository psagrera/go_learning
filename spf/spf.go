package main

import (
    
    "fmt"
    d "github.com/psagrera/dijkstra"
)                      

type Graph map[string]map[string]int

func main() {
   
    g := Graph{
        
        "a":{"b":20,"c":80},
        "b":{"a": 20, "c": 20},
        "c": {"a": 80, "b": 20},
    }
    fmt.Println("Graph",g)
    path, cost, _ := g.d.Path("a","c")
    fmt.Printf("path: %v, cost: %v", path, cost)
}
