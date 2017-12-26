import math
import operator
from functools import reduce

MAX = 256
raw = input()

f = []
for c in raw:
    f.append(ord(c))

f.extend([17, 31, 73, 47, 23])

my_list = [i for i in range(MAX)]

c = 0
skip_size = 0
for v in range(64):
    for l in f:
        to_reverse = []
        for u in range(l):
            pos = (c + u) % MAX
            to_reverse.append(pos)

        for i in range(int(math.floor(l / 2))):
            my_list[to_reverse[i]], my_list[to_reverse[l - i - 1]] = my_list[to_reverse[l - i - 1]], my_list[to_reverse[i]]

        c = (c + l + skip_size) % MAX
        skip_size = skip_size + 1

2dense_hash = []
for h in range(0, 16):
    to_xor = my_list[(16 * h):(16 * h + 16)]
    dense_hash.append(reduce(operator.xor, to_xor, 0))

hash_str = "".join(map(lambda c: '%0*x' % (2, c), dense_hash))

print(hash_str)
print("63960835bcdc130f0b66d7ff4f6a5a8e")
print(hash_str=="63960835bcdc130f0b66d7ff4f6a5a8e")
