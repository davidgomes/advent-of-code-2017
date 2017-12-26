
input = 382
N = 50000000
pos = 0
last = -1
for l in range(1,N+1):
    pos = (pos + 1 + input) % l
    if pos == 0:
        last = l
    #print(l, last)
print(last)
