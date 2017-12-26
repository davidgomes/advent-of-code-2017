import math
import operator
from functools import reduce

MAX=256

def knot_hash(raw):
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

    dense_hash = []
    for h in range(0, 16):
        to_xor = my_list[(16 * h):(16 * h + 16)]
        dense_hash.append(reduce(operator.xor, to_xor, 0))

    final_res = "".join(map(lambda c: '%0*x' % (2, c), dense_hash))
    return final_res

conv = { "0": "0000",
	 "1": "0001",
	 "2": "0010",
	 "3": "0011",
	 "4": "0100",
	 "5": "0101",
	 "6": "0110",
	 "7": "0111",
	 "8": "1000",
	 "9": "1001",
	 "a": "1010",
	 "b": "1011",
	 "c": "1100",
	 "d": "1101",
	 "e": "1110",
	 "f": "1111" }

def hex_string_to_bin(s): # a0c2017
    bin_res = ""
    for c in s:
        bin_res += "%s" % conv[c]
    return bin_res

res = 0
for i in range(128):
    hex_hash = knot_hash("jxqlasbh-%d" % i)

    print(hex_hash)
    bin_res = hex_string_to_bin(hex_hash)
    res += bin_res.count("1")

print(res)
