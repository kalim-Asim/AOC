N = 100
ans, cur = 0, 50

with open("input.txt", "r") as f:
  for line in f:
    line = line.strip()

    if not line:
      break 

    dir = line[0]
    num = int(line[1:])

    ans += int(num / N) 
    num %= N 

    if dir == 'L':
      if cur != 0 and num > cur:
        ans += 1
      cur = (cur - num + N) % N 
    else: 
      if num > N - cur:
        ans += 1
      cur = (cur + num) % N
    ans += cur == 0

print(ans)   