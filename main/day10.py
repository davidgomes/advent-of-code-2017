import math
MAX = 256
f = list(map(int, input().split(",")))
my_list = [i for i in range(MAX)]

c = 0
skip_size = 0
for l in f:
    to_reverse = []
    for u in range(l):
        pos = (c + u) % MAX
        to_reverse.append(pos)

    for i in range(int(math.floor(l / 2))):
        my_list[to_reverse[i]], my_list[to_reverse[l - i - 1]] = my_list[to_reverse[l - i - 1]], my_list[to_reverse[i]]

    c = (c + l + skip_size) % MAX
    skip_size = skip_size + 1

print(my_list)
