import functools

def flip_v(pattern):
    lines = pattern.split("/")
    return "/".join(lines[::-1])

def flip_h(pattern):
    result = ""
    for line in pattern.split("/"):
        result += "{}/".format(line[::-1])
    return result[:-1]

def rotate90(pattern):
    columns = []

    for line in pattern.split("/"):
        columns = [line] + columns

    res = ""
    for i in range(len(columns)):
        for column in columns:
            res += column[i]

        if i != len(columns) - 1:
            res += "/"

    return res

def rotate180(pattern):
    return rotate90(rotate90(pattern))

def rotate270(pattern):
    return rotate90(rotate90(rotate90(pattern)))

def replace(new_lines, pos_x, pos_y, size, mini_pattern):
    mini_lines = mini_pattern.split("/")

    for y in range(pos_y, pos_y + size):
        if y >= len(new_lines):
            new_lines.append("")

        for x in range(pos_x, pos_x + size):
            new_lines[y] = new_lines[y] + mini_lines[y - pos_y][x - pos_x]

    return new_lines

def print_lines(lines):
    for line in lines:
        print(line)

    print()

def main():
    f_conversions = []

    with open("day21.in", "r") as f:
        f_conversions = list(map(lambda line: list(map(lambda x: x.strip(), line.split("=>"))), f.readlines()))

    conversions = {}

    for conversion in f_conversions:
        inp = conversion[0]
        out = conversion[1]

        conversions[inp] = out
        conversions[flip_v(inp)] = out
        conversions[flip_h(inp)] = out

        conversions[rotate90(inp)] = out
        conversions[rotate90(flip_v(inp))] = out
        conversions[rotate90(flip_h(inp))] = out

        conversions[rotate180(inp)] = out
        conversions[rotate180(flip_v(inp))] = out
        conversions[rotate180(flip_h(inp))] = out

        conversions[rotate270(inp)] = out
        conversions[rotate270(flip_v(inp))] = out
        conversions[rotate270(flip_h(inp))] = out

    pattern = ".#./..#/###"    # pattern = "#..#/..../..../#..#"
    lines = pattern.split("/")

    for i in range(5):
        new_lines = []

        step = 2 if len(lines) % 2 == 0 else 3

        mini_y = 0
        for y in range(step - 1, len(lines), step):
            mini_x = 0

            for x in range(step - 1, len(lines[y]), step):
                mini_pattern = ""

                for k in range(y - step + 1, y + 1):
                    mini_pattern += lines[k][x - step + 1:x + 1]
                    mini_pattern += "/"

                mini_pattern = mini_pattern[:-1]
                new_lines = replace(new_lines, mini_x, mini_y, len(conversions[mini_pattern].split("/")), conversions[mini_pattern])
                mini_x += len(conversions[mini_pattern].split("/"))

            mini_y = len(new_lines)

        lines = new_lines

    count = 0

    print(functools.reduce(lambda cur, line: cur + line.count("#"), lines, 0))

main()
