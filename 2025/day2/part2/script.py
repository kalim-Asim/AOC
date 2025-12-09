import bisect 

sum = 0
arr = []

def precompute():
    for d in range(1, 100000):
        str_, s = str(d), str(d)
        while len(s) <= 10:
          s += str_
          if len(s) <= 10:
            num = int(s)
            arr.append(num)

precompute()
res = sorted(set(arr))

def lower_bound(x):
    return bisect.bisect_left(res, x)

def upper_bound(x):
    return bisect.bisect_right(res, x)

def count(l, r):
    cnt = 0
    lb, ub = lower_bound(l), upper_bound(r)

    for i in range(lb, ub):
        cnt += res[i]
    return cnt

with open("input.txt", "r") as f:
    text = f.read().strip()
    pairs = text.split(",")
    
    for p in pairs:
        l, r = map(int, p.split("-"))
        sum += count(l, r)
        
print(sum)
