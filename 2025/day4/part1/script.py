ans = 0
grid = []
dirs = [[-1, 0], [1, 0], [1, 1], [-1, 1], [1, -1], [0, -1], [0, 1], [-1, -1]]

with open("input.txt", "r") as f:
  for line in f.readlines():
    line = line.strip()
    if line == "":
      continue

    grid.append(list(line))

rows = len(grid)
cols = len(grid[0])

def valid(r, c):
  return r >= 0 and r < rows and c < cols and c >= 0

for r in range(0, rows):
  for c in range(0, cols):
    if grid[r][c] == '.':
      continue

    cnt = 0
    for dr, dc in dirs:
      nr, nc = r + dr, c + dc 
      if not valid(nr, nc):
        continue 
      cnt += grid[nr][nc] == '@'

    ans += cnt < 4

print(ans)