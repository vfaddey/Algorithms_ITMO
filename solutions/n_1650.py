import heapq

def main():
    n = int(input())
    wealth = {}
    city_of = {}
    city_totals = {}

    for _ in range(n):
        name, city, w = input().split()
        w = int(w)
        wealth[name] = w
        city_of[name] = city
        city_totals[city] = city_totals.get(city, 0) + w

    m, k = map(int, input().split())
    moves = []  # (day, name, dest)
    all_cities = set(city_totals.keys())

    for _ in range(k):
        d, name, dest = input().split()
        d = int(d)
        moves.append((d, name, dest))
        all_cities.add(dest)
        if dest not in city_totals:
            city_totals[dest] = 0

    eff = {}
    for d, name, dest in moves:
        ed = d + 1
        if ed <= m:
            eff.setdefault(ed, []).append((name, dest))

    event_days = sorted(eff.keys())

    heap = []
    for city in all_cities:
        heapq.heappush(heap, (-city_totals[city], city))

    def top_two():
        """Возвращает (город_max, сумма_max, уникальна ли вершина)."""
        temp = []
        while heap:
            neg_t1, c1 = heapq.heappop(heap)
            t1 = -neg_t1
            if city_totals[c1] == t1:
                break
        else:
            return None, 0, True

        while heap:
            neg_t2, c2 = heapq.heappop(heap)
            t2 = -neg_t2
            if city_totals[c2] == t2:
                break
        else:
            heapq.heappush(heap, (-t1, c1))
            return c1, t1, True

        heapq.heappush(heap, (-t1, c1))
        heapq.heappush(heap, (-t2, c2))

        return (c1, t1, t1 > t2)

    max_city, max_sum, is_unique = top_two()
    day_start = 1
    days_count = {city: 0 for city in all_cities}

    for ed in event_days:
        if day_start > m:
            break
        day_end = ed - 1
        if day_end >= day_start and is_unique:
            days_count[max_city] += (day_end - day_start + 1)

        # применяем все перемещения, действующие с ed-го дня
        for name, dest in eff[ed]:
            old = city_of[name]
            w = wealth[name]
            city_totals[old] -= w
            heapq.heappush(heap, (-city_totals[old], old))
            city_totals[dest] += w
            heapq.heappush(heap, (-city_totals[dest], dest))
            city_of[name] = dest

        max_city, max_sum, is_unique = top_two()
        day_start = ed

    if day_start <= m and is_unique:
        days_count[max_city] += (m - day_start + 1)

    for city in sorted(all_cities):
        cnt = days_count.get(city, 0)
        if cnt > 0:
            print(city, cnt)

if __name__ == "__main__":
    main()
