// Command line pagerank
// ---------------------
//
//    $ cat example.in
//    0   1
//    1   2
//    2   3   4   5
//    4   2
//    5   1
//
//    $ pagerank example.in 2> /dev/null | sort -k2,2 -nr
//    2   0.33477170103317816
//    1   0.20154082325712538
//    5   0.13963495553328562
//    4   0.13963495553328562
//    3   0.13963495553328562
//    0   0.0447826091098397
package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
)

type Graph struct {
    Nodes []GraphNode
}

type GraphNode struct {
    OutboundNeighbors []int
}

// Taken from https://github.com/cosbynator/WikiRank
// Written by Thomas Dimson (tdimson@cs.stanford.edu)
func pageRankGraph(g Graph, walkProbability float64, convergenceCriteron float64) []float64 {
    beta, epsilon := walkProbability, convergenceCriteron
    log.Printf("Ranking with beta='%f', epsilon='%f'", beta, epsilon)
    n := len(g.Nodes)
    lastRank := make([]float64, n)
    thisRank := make([]float64, n)

    for iteration, lastChange := 1, math.MaxFloat64; lastChange > epsilon; iteration++ {
        thisRank, lastRank = lastRank, thisRank
        if iteration > 1 {
            // Clear out old values
            for i := 0; i < n; i++ {
                thisRank[i] = 0.0
            }
        } else {
            // Base case: everything uniform
            for i := 0; i < n; i++ {
                lastRank[i] = 1.0 / float64(n)
            }
        }

        // Single power iteration
        for i := 0; i < n; i++ {
            contribution := beta * lastRank[i] / float64(len(g.Nodes[i].OutboundNeighbors))
            for _, linkId := range g.Nodes[i].OutboundNeighbors {
                // fmt.Println(linkId)
                thisRank[linkId] += contribution
            }
        }

        // Reinsert leaked probability
        S := float64(0.0)
        for i := 0; i < n; i++ {
            S += thisRank[i]
        }
        leakedRank := (1.0 - S) / float64(n)
        lastChange = 0.0 // and calculate L1-difference too
        for i := 0; i < n; i++ {
            thisRank[i] += leakedRank
            lastChange += math.Abs(thisRank[i] - lastRank[i])
        }

        log.Printf("Pagerank iteration #%d delta=%f", iteration, lastChange)
    }

    return thisRank
}

func main() {

    log.SetOutput(os.Stderr)

    walkProbability := flag.Float64("w", 0.85, "walk probability")
    convergenceCriteron := flag.Float64("c", 0.0001, "convergence criteron")

    var PrintUsage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] FILE\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "File format: TSV (NODE OUTBOUND [OUTBOUND, ...])\n")
        flag.PrintDefaults()
    }

    flag.Parse()

    if flag.NArg() < 1 {
        PrintUsage()
        os.Exit(1)
    }

    fileName := flag.Args()[0]

    var in *os.File
    if fileName == "-" {
      in = os.Stdin
    } else {
      in, err := os.Open(fileName)
      if err != nil {
          fmt.Printf("%s\n", err)
          os.Exit(1)
      }
      defer func() {
          if err := in.Close(); err != nil {
              panic(err)
          }
      }()
    }
    scanner := bufio.NewScanner(in)
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    // intermediate store of outbound links, that are actually defined
    defined := make(map[int][]int)
    maximum := 0

    for scanner.Scan() {
        fields := strings.Fields(scanner.Text())
        head, _ := strconv.Atoi(fields[0])
        if head > maximum {
            maximum = head
        }
        rest := fields[1:]
        nodelist := make([]int, len(rest))
        for i, value := range rest {
            converted, _ := strconv.Atoi(value)
            if converted > maximum {
                maximum = converted
            }
            nodelist[i] = converted
        }
        defined[head] = nodelist
    }

    log.Printf("%d nodes\n", maximum)
    nodes := make([]GraphNode, maximum+1)

    for i, _ := range nodes {
        cached, present := defined[i]
        if present {
            nodes[i] = GraphNode{OutboundNeighbors: cached}
        } else {
            nodes[i] = GraphNode{}
        }
    }

    g := Graph{Nodes: nodes}
    rankVector := pageRankGraph(g, *walkProbability, *convergenceCriteron)

    for i, value := range rankVector {
        fmt.Printf("%d\t%s\n", i, strconv.FormatFloat(value, 'f', -1, 64))
    }
}
