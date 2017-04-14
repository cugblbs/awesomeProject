package main

type Vertex struct {
	Lat,Long float64
}

var m = map[string] Vertex {
	"Bell Labs":Vertex{40.68433, -74.39967},
	"Google":Vertex{37.42202, -122.08408},
}

func main() {
	for name,v := range m{
		println(name, v.Lat, v.Long)
	}
}
