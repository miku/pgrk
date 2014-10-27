package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"gopkg.in/fatih/set.v0"
)

func main() {
	n := flag.Int("n", 10, "number of vertices")
	p := flag.Float64("p", 0.67, "density measure")
	c := flag.Bool("c", false, "allow zero hop cycles")

	flag.Parse()

	if *p >= 1.0 {
		log.Fatal("p must be < 1")
	}

	edges := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < *n; i++ {
		dsts := set.New()
		var parts []string
		pp := *p
		for rand.Float64() < pp {
			d := r.Intn(*n)
			if d == i && !*c {
				continue
			}
			dsts.Add(d)
			edges++
			pp = pp * pp
		}
		parts = append(parts, strconv.Itoa(i))
		ss := set.IntSlice(dsts)
		for _, d := range ss {
			parts = append(parts, strconv.Itoa(d))
		}
		fmt.Printf("%s\n", strings.Join(parts, "\t"))
	}
}
