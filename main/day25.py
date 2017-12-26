def step(machine, state, current):
    if state == "A":
        if current not in machine or machine[current] == 0:
            machine[current] = 1
            current = current + 1
            state = "B"
        else: # machine[current] is 1
            machine[current] = 0
            current = current - 1
            state = "B"
    elif state == "B":
        if current not in machine or machine[current] == 0:
            machine[current] = 0
            current = current + 1
            state = "C"
        else:
            machine[current] = 1
            current = current - 1
            state = "B"
    elif state == "C":
        if current not in machine or machine[current] == 0:
            machine[current] = 1
            current = current + 1
            state = "D"
        else:
            machine[current] = 0
            current = current - 1
            state = "A"
    elif state == "D":
        if current not in machine or machine[current] == 0:
            machine[current] = 1
            current = current - 1
            state = "E"
        else:
            machine[current] = 1
            current = current - 1
            state = "F"
    elif state == "E":
        if current not in machine or machine[current] == 0:
            machine[current] = 1
            current = current - 1
            state = "A"
        else:
            machine[current] = 0
            current = current - 1
            state = "D"
    elif state == "F":
        if current not in machine or machine[current] == 0:
            machine[current] = 1
            current = current + 1
            state = "A"
        else:
            machine[current] = 1
            current = current - 1
            state = "E"

    return state, current

def main():
    machine = {}
    state = "A"
    current = 0

    for i in range(12629077):
        state, current = step(machine, state, current)

    count = 0
    for cell in machine:
        if machine[cell] == 1:
            count = count + 1

    print(count)
