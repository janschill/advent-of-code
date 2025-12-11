package main

import (
	"fmt"
	"log"
	"math"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/janschill/advent-of-code-2025/puzzles/helpers"
)

func main() {
	input := filepath.Join(puzzleDir(), "input.txt")
	lines := helpers.MustLines(input)
	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 2: ", part2(lines))
}

type Point struct {
	X float64
	Y float64
	Z float64
}

func euclidianDistance(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dz := p2.Z - p1.Z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

type Edge struct {
	d float64
	i int // index first point
	j int // index second point
}

func distances(points []Point) []Edge {
	edges := make([]Edge, 0, len(points)*(len(points)-1)/2)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, Edge{
				d: euclidianDistance(points[i], points[j]),
				i: i,
				j: j,
			})
		}
	}
	sort.Slice(edges, func(a, b int) bool {
		return edges[a].d < edges[b].d
	})
	return edges
}

type DSU struct {
	parent []int
	size   []int
}

func newDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &DSU{parent: parent, size: size}
}

func (d *DSU) find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) union(a, b int) {
	ra, rb := d.find(a), d.find(b)
	if ra == rb {
		return
	}
	if d.size[ra] < d.size[rb] {
		ra, rb = rb, ra
	}
	d.parent[rb] = ra
	d.size[ra] += d.size[rb]
}

// connect 1000 closest junction boxes
// multiply the sizes of the three largest circuits
// use straight-line distance to find all distances
// sort them
// connect the first 1000 closest connections
// connecting creates a circuit
// keep track of existing circuits
// when connecting two new junction boxes, check circuits
func part1(lines []string) (total int) {
	var points []Point
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(parts[0], 64)
		y, _ := strconv.ParseFloat(parts[1], 64)
		z, _ := strconv.ParseFloat(parts[2], 64)
		points = append(points, Point{
			X: x,
			Y: y,
			Z: z,
		})
	}
	edges := distances(points)[:1000]
	dsu := newDSU(len(points))
	for _, edge := range edges {
		dsu.union(edge.i, edge.j)
	}
	counts := make(map[int]int)
	for i := range points {
		root := dsu.find(i)
		counts[root]++
	}
	var sizes []int
	for _, c := range counts {
		sizes = append(sizes, c)
	}
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })
	if len(sizes) < 3 {
		return 0
	}

	fmt.Println("Sizes of largest circuits:", sizes[0], sizes[1], sizes[2])

	return sizes[0] * sizes[1] * sizes[2]
}

func part2(lines []string) (total int) {
	return
}

func puzzleDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to determine puzzle directory")
	}
	return filepath.Dir(filename)
}
