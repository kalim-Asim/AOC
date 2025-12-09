import bisect 

sum = 0
arr = []

def precompute():
    for d in range(1, 100000):
        s = str(d)
        s += s
        num = int(s)
        arr.append(num)

precompute()
arr.sort()

# some no., lb >= x
def lower_bound(x):
    return bisect.bisect_left(arr, x)

# some no. ub > x
def upper_bound(x):
    return bisect.bisect_right(arr, x)

def count(l, r):
    res = 0
    lb, ub = lower_bound(l), upper_bound(r)
    for i in range(lb, ub):
        res += arr[i]
    return res

with open("input.txt", "r") as f:
    text = f.read().strip()
    pairs = text.split(",")
    
    for p in pairs:
        l, r = map(int, p.split("-"))
        sum += count(l, r)
        
print(sum)
