README
======

Simple command line PageRank calculator.

The input file represents a directed graph as adjacency list, where each node
must be identified by an integer (translation from and to your domain must
happen elsewhere).

Example input file:

    1   2
    2   3   4   5
    4   2
    5   1

Usage:

    $ pagerank example.in |sort -k2,2 -nr
    2013/11/11 10:55:37 Ranking with beta='0.850000', epsilon='0.000100'
    2013/11/11 10:55:37 Pagerank iteration #1 delta=0.472222
    2013/11/11 10:55:37 Pagerank iteration #2 delta=0.200694
    2013/11/11 10:55:37 Pagerank iteration #3 delta=0.064445
    2013/11/11 10:55:37 Pagerank iteration #4 delta=0.044575
    2013/11/11 10:55:37 Pagerank iteration #5 delta=0.032411
    2013/11/11 10:55:37 Pagerank iteration #6 delta=0.011727
    2013/11/11 10:55:37 Pagerank iteration #7 delta=0.005198
    2013/11/11 10:55:37 Pagerank iteration #8 delta=0.004099
    2013/11/11 10:55:37 Pagerank iteration #9 delta=0.002494
    2013/11/11 10:55:37 Pagerank iteration #10 delta=0.000892
    2013/11/11 10:55:37 Pagerank iteration #11 delta=0.000510
    2013/11/11 10:55:37 Pagerank iteration #12 delta=0.000382
    2013/11/11 10:55:37 Pagerank iteration #13 delta=0.000169
    2013/11/11 10:55:37 Pagerank iteration #14 delta=0.000060
    2   0.32908887088503813
    1   0.17789822317101825
    5   0.14656550722676537
    4   0.14656550722676537
    3   0.14656550722676537
    0   0.05331638426364749


Options:

    $ pagerank 
    Usage: pagerank [OPTIONS] FILE
    File format: TSV (NODE OUTBOUND [OUTBOUND, ...])
      -c=0.0001: convergence criteron
      -w=0.85: walk probability
