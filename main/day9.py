f = str(input())

curly_stack = []

in_garbage = False
ignore = False
score = 0
for i in range(len(f)):
    c = f[i]

    if c == "<" and not in_garbage:
        in_garbage = True
        continue

    if in_garbage:
        if ignore:
            ignore = False
            continue

        if c == "!":
            ignore = True
            continue

        if c == ">":
            in_garbage = False

        continue

    if c == "{":
        curly_stack.append("{")
    if c == "}":
        score += len(curly_stack)
        curly_stack.pop()

print(score)
