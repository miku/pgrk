#!/usr/bin/env python
# coding: utf-8

"""
Visualize random directed graphs with dot.

Input format:

    0   2   3   4
    1   5
    2   0   1
    3   1   5
    4

"""

import fileinput
import random
import string


def random_string(length=16):
    return ''.join(random.choice(string.letters) for _ in range(length))


def to_dot():
    print('digraph %s {' % random_string(length=5))
    for line in fileinput.input():
        nodes = line.strip().split()
        for node in nodes[1:]:
            print('    %s -> %s;' % (nodes[0], node))
    print('}')

if __name__ == '__main__':
    to_dot()
