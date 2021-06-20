package main

import "fmt"

func getNeck(request GameRequest) string {
	head := Coordinate{
		X: request.Battlesnake.Head.X,
		Y: request.Battlesnake.Head.Y,
	}

	unsafe := ""

	for i := range request.Battlesnake.Body {
		if head.X+1 == request.Battlesnake.Body[i].X {
			unsafe = "right"
		}
		if head.X-1 == request.Battlesnake.Body[i].X {
			unsafe = "left"
		}
		if head.Y+1 == request.Battlesnake.Body[i].Y {
			unsafe = "up"
		}
		if head.Y-1 == request.Battlesnake.Body[i].Y {
			unsafe = "down"
		}
	}

	fmt.Printf("unsafe:%q", unsafe)
	return unsafe
}
