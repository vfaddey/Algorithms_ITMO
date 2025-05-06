import heapq

n = int(input())
T = list(map(int, input().split()))
nums = [input().strip() for _ in range(n)]
src, dst = 0, n-1


idx = {nums[i]: i for i in range(n)}

INF = 10**30
dist = [INF] * n
prev = [-1] * n

dist[src] = 0
pq = [(0, src)]

while pq:
    d, u = heapq.heappop(pq)
    if d > dist[u]:
        continue
    if u == dst:
        break
    s = nums[u]

    for i in range(10):
        orig = s[i]
        for digit in '0123456789':
            if digit == orig:
                continue
            v_str = s[:i] + digit + s[i+1:]
            v = idx.get(v_str)
            if v is None:
                continue
            nd = d + T[i]
            if nd < dist[v]:
                dist[v] = nd
                prev[v] = u
                heapq.heappush(pq, (nd, v))

    s_list = list(s)
    for i in range(10):
        for j in range(i+1, 10):
            if s_list[i] == s_list[j]:
                continue
            s_list[i], s_list[j] = s_list[j], s_list[i]
            v_str = ''.join(s_list)
            v = idx.get(v_str)
            if v is not None:
                nd = d + T[i]
                if nd < dist[v]:
                    dist[v] = nd
                    prev[v] = u
                    heapq.heappush(pq, (nd, v))
            s_list[i], s_list[j] = s_list[j], s_list[i]

if dist[dst] >= INF:
    print(-1)
    sys.exit(0)

path = []
cur = dst
while cur != -1:
    path.append(cur + 1)
    cur = prev[cur]
path.reverse()

print(dist[dst])
print(len(path))
print(*path)


