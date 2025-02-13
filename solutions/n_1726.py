from math import floor

n = int(input())
xs = []
ys = []
for i in range(n):
    x, y = map(int, input().split())
    xs.append(x)
    ys.append(y)
xs.sort()
ys.sort()
sum = 0
for i in range(n):
    sum += xs[i] * (2*i - n + 1)
    sum += ys[i] * (2*i - n + 1)

sum *= 2
walks = n * (n-1)
print(floor(sum/walks))