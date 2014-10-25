README
======

Simple command line [PageRank](http://infolab.stanford.edu/~backrub/google.html) calculator.

The input file represents a directed graph as adjacency list, where each node
must be identified by an integer (translation from and to your domain must
happen elsewhere).

Installation
------------

    $ go get github.com/miku/pgrk/cmd/{pgrk,pgrk-gen,pgrk-dot}

Example
-------

![Shiny graph](http://i.imgur.com/0ZQYLFl.png)

Example input file:

    $ cat example.in 
    0   1
    1   2
    2   3   4   5
    4   2
    5   1


Usage:

    $ pagerank example.in | sort -k2,2 -nr
    2013/11/11 12:54:27 Ranking with beta='0.850000', epsilon='0.000100'
    2013/11/11 12:54:27 Pagerank iteration #1 delta=0.661111
    2013/11/11 12:54:27 Pagerank iteration #2 delta=0.361250
    2013/11/11 12:54:27 Pagerank iteration #3 delta=0.217029
    2013/11/11 12:54:27 Pagerank iteration #4 delta=0.163127
    2013/11/11 12:54:27 Pagerank iteration #5 delta=0.059629
    2013/11/11 12:54:27 Pagerank iteration #6 delta=0.020982
    2013/11/11 12:54:27 Pagerank iteration #7 delta=0.017126
    2013/11/11 12:54:27 Pagerank iteration #8 delta=0.008490
    2013/11/11 12:54:27 Pagerank iteration #9 delta=0.002520
    2013/11/11 12:54:27 Pagerank iteration #10 delta=0.001689
    2013/11/11 12:54:27 Pagerank iteration #11 delta=0.001104
    2013/11/11 12:54:27 Pagerank iteration #12 delta=0.000350
    2013/11/11 12:54:27 Pagerank iteration #13 delta=0.000179
    2013/11/11 12:54:27 Pagerank iteration #14 delta=0.000132
    2013/11/11 12:54:27 Pagerank iteration #15 delta=0.000045
    2   0.33477170103317816
    1   0.20154082325712538
    5   0.13963495553328562
    4   0.13963495553328562
    3   0.13963495553328562
    0   0.0447826091098397



Options:

    $ pagerank 
    Usage: pagerank [OPTIONS] FILE
    File format: TSV (NODE OUTBOUND [OUTBOUND, ...])
      -c=0.0001: convergence criteron
      -w=0.85: walk probability


Benchmark
---------

* Computing the pagerank of an 10,004,750 node sparse graph takes about 30s (tested on i5-3470, 4 core, 3.2GHz machine, 16G RAM).
* `{"p": 0.8, "nodes": 30000000, "edges": 46735510}`: 1m55.010s
* `{"p": 0.8, "nodes": 50000000, "edges": 77879276}`: 3m6.173s
* `{"p": 0.99, "nodes": 50000000, "edges": 268891019}`: 4m10.350s


Utils
-----

Generate and visualize random directed graphs, with `pgrk-gen`
and `pgrk-dot`:

    $ pgrk-gen -n 30 -p 0.75 > test.in
    $ cat test.in | pgrk-dot | dot -Tpdf -o test.pdf && open test.pdf
    $ pgrk test.in 2> /dev/null | sort -nrk2,2 | head -5
    20  0.10566870915778451
    18  0.10238092345567337
    6   0.10012509291487659
    28  0.06610747553255986
    16  0.06485242652853762

![Another Shiny graph](http://i.imgur.com/hzzKtzq.png)


Credits
-------

* [Thomas Dimson](https://github.com/cosbynator) for [WikiRank](https://github.com/cosbynator/WikiRank)
