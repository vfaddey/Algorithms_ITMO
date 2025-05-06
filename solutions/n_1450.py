from collections import deque

N, M = map(int, input().split())
adj = [[] for _ in range(N+1)]
indegree = [0] * (N+1)

for _ in range(M):
    u, v, c = map(int, input().split())
    adj[u].append((v, c))
    indegree[v] += 1

S, F = map(int, input().split())
q = deque(i for i in range(1, N+1) if indegree[i] == 0)
topo = []

while q:
    u = q.popleft()
    topo.append(u)
    for v, _ in adj[u]:
        indegree[v] -= 1
        if indegree[v] == 0:
            q.append(v)

NEG_INF = -10**18
dp = [NEG_INF] * (N+1)
dp[S] = 0

for u in topo:
    if dp[u] == NEG_INF:
        continue
    for v, w in adj[u]:
        if dp[u] + w > dp[v]:
            dp[v] = dp[u] + w

if dp[F] == NEG_INF:
    print("No solution")
else:
    print(dp[F])
