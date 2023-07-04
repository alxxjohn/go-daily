package main

import (
	"fmt"
	"contatiner/queue"
)

func main(){

}


func bfs(graph, start string){

	visit := make([int]struct{})
	queue := queue.New()
	queue.Enqueue(start)

	for queue {

		node, _ := queue.Get()

		for elem, _ := range graph[node]{

			if elem not in visit {

				new = append(visit, elem)
			}
		}
	}
	return visit

}