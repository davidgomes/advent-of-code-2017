def burst(grid_map, x, y, direction, counter):
    if (x, y) in grid_map and grid_map[(x, y)] == "#":
        if direction == "up":
            direction = "right"
        elif direction == "right":
            direction = "down"
        elif direction == "down":
            direction = "left"
        elif direction == "left":
            direction = "up"
    else:
        if direction == "up":
            direction = "left"
        elif direction == "left":
            direction = "down"
        elif direction == "down":
            direction = "right"
        elif direction == "right":
            direction = "up"

    if (x, y) not in grid_map:
        grid_map[(x, y)] = "#"
        counter = counter + 1
    else:
        if grid_map[(x, y)] == "#":
            grid_map[(x, y)] = "."
        else:
            grid_map[(x, y)] = "#"
            counter = counter + 1

    if direction == "up":
        return x, y - 1, direction, counter
    elif direction == "right":
        return x + 1, y, direction, counter
    elif direction == "down":
        return x, y + 1, direction, counter
    elif direction == "left":
        return x - 1, y, direction, counter

def main():
    with open("day22.in") as f:
        grid = f.readlines()

    grid_map = {}

    for y in range(len(grid)):
        for x in range(len(grid[0])):
            if grid[y][x] == "#":
                grid_map[(x, y)] = "#"

    # 5380, 5404, 5521, 5535
    x = 13
    y = 12
    direction = "up"
    counter = 0

    for i in range(10000):
        print(x, y)
        x, y, direction, counter = burst(grid_map, x, y, direction, counter)

    return counter
