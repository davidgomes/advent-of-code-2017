with open("day8.in", "r") as o:
    f = o.readlines()

registers = {}

operators = {">": lambda a, b: a > b,
             "<": lambda a, b: a < b,
             "==": lambda a, b: a == b,
             ">=": lambda a, b: a >= b,
             "<=": lambda a, b: a <= b,
             "!=": lambda a, b: a != b}

for line in f:
    l = line.replace("\n", "").split(" ")

    cond_l = l[4]
    cond_o = l[5]
    cond_r = int(l[6])

    if cond_l not in registers:
        registers[cond_l] = 0

    l_value = registers[cond_l]

    cond_result = operators[cond_o](l_value, cond_r)

    if l[0] not in registers:
        registers[l[0]] = 0

    if cond_result:
        if l[1] == "inc":
            registers[l[0]] += int(l[2])
        else:
            registers[l[0]] -= int(l[2])

print(registers)

max_value = None
for r in registers:
    if not max_value:
        max_value = registers[r]
    else:
        max_value = max(max_value, registers[r])

print(max_value)
