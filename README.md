README
======

Simple command line PageRank calculator.

The input file represents a directed graph as adjacency list, where each node
must be identified by an integer (translation from and to your domain must
happen elsewhere).

Example input file:

    $ cat example.in 
    0   1
    1   2
    2   3   4   5
    4   2
    5   1


Usage:

    $ pagerank example.in |sort -k2,2 -nr
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
