
res = 0
k = 12

def find(start, s, end):
  index, largest = start, s[start]
  for i in range(start, end):
    if s[i] > largest:
      largest = s[i]
      index = i 
  return largest, index + 1

with open("input.txt", "r") as f:
  for s in f:
    s = s.strip()

    if not s:
      break
    
    str_, start = "", 0
    
    for i in range(0, k):
      end = len(s) - (k - (i + 1))
      ch, start = find(start, s, end)
      str_ += ch

    res += int(str_)

print(res)