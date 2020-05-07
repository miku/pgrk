package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/miku/pgrk"
	"gopkg.in/fatih/set.v0"
)

func main() {
	n := flag.Int("n", 10, "number of vertices")
	p := flag.Float64("p", 0.67, "density measure")
	c := flag.Bool("c", false, "allow zero hop cycles")
	version := flag.Bool("v", false, "prints current version and exits")

	flag.Parse()

	if *version {
		fmt.Println(pgrk.Version)
		os.Exit(0)
	}

	if *p >= 1.0 {
		log.Fatal("p must be < 1")
	}

	edges := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < *n; i++ {
		dsts := set.New(set.NonThreadSafe)
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
