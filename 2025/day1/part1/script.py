N = 100
ans, cur = 0, 50

with open("input.txt", "r") as f:
    for line in f:
        line = line.strip() #removes trailing space

        if not line:
            break

        dir = line[0]
        num = int(line[1:])

        if dir == "L":
                cur = (cur - num + N) % N
        else:
                cur = (cur + num) % N

        ans += cur == 0

print(ans)