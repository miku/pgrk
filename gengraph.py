#!/usr/bin/env python
# coding: utf-8

"""
Generate random directed graphs for testing `pagerank`.

With p = 0.9999999999999999, you'll likely get a complete digraph.
"""
from __future__ import print_function
import argparse
import json
import random
import sys


def main():
    parser = argparse.ArgumentParser(prog='gengraph.py')
    parser.add_argument('-n', type=int, default=10, help='number of vertices')
    parser.add_argument('-p', type=float, default=0.67, help='density measure')
    parser.add_argument('-c', action='store_true', help='allow edges to self')

    args = parser.parse_args()

    if args.p >= 1.0:
        raise ValueError('p only defined on [0, 1)')

    n = args.n
    c = args.c
    rr = random.random
    ri = random.randint
    edges = 0

    for origin in xrange(n):
        destinations = set()
        p = args.p
        while rr() < p:
            destinations.add(ri(0, n))
            edges += 1
            p = p * p
        if c:
            parts = [origin] + list(destinations)
        else:
            parts = [origin] + list(destinations - set([origin]))
        print('\t'.join(map(str, parts)), file=sys.stdout)
    print(json.dumps(dict(nodes=n, edges=edges, p=args.p)), file=sys.stderr)


if __name__ == '__main__':
    main()
