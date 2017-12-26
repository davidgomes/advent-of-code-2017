import math

with open("day20.in", "r") as f:
    lines = f.readlines()

lens = []
minimum = 0

for i in range(len(lines)):
    a = list(map(int, lines[i].split(" ")[2][2:-1].replace("<", "").replace(">", "").split(",")))
    # lens.append(math.sqrt(a[0]**2 + a[1]**2 + a[2]**2))
    lens.append(abs(a[0]) + abs(a[1]) + abs(a[2]))
    minimum = i if lens[minimum] > lens[i] else minimum

print(minimum, lens[minimum])
