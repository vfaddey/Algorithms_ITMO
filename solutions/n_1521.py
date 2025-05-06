from collections import deque

n, k = map(int, input().split())
dq = deque(range(1, n + 1))

while dq:
    dq.rotate(-(k - 1))
    print(dq.popleft())
