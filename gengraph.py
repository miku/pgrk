#!/usr/bin/env python
# coding: utf-8

"""
Generate random directed graphs for testing `pagerank`.

With p = 0.9999999999999999, you'll likely get a complete digraph.
"""

import argparse
import random


def main():
    parser = argparse.ArgumentParser(prog='gengraph.py')
    parser.add_argument('-n', type=int, default=10, help='number of vertices')
    parser.add_argument('-p', type=float, default=0.67, help='density measure')
    parser.add_argument('-c', action='store_true', help='allow edges to self')

    args = parser.parse_args()

    if args.p >= 1.0:
        raise ValueError('p only defined on [0, 1)')

    for origin in range(args.n):
        destinations = set()
        p = args.p
        while random.random() < p:
            destinations.add(random.randint(0, args.n))
            p = p * p
        if args.c:
            line = [origin] + list(destinations)
        else:
            line = [origin] + list(destinations - set([origin]))
        print('\t'.join(map(str, line)))


if __name__ == '__main__':
    main()
