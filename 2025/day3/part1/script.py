res = 0

with open("input.txt", "r") as f:
  for s in f:
    s = s.strip()

    if not s:
      break 

    if len(s) == 1:
      res += int(s)
      continue

    idx1, first = 0, '0'
    for i in range(0, len(s) - 1):
      if s[i] > first:
        first = s[i]
        idx1 = i 
    
    sec = '0'
    for i in range(idx1 + 1, len(s)):
      if s[i] > sec:
        sec = s[i]
        idx2 = i 
    
    str = first + sec 
    res += int(str)

print(res)