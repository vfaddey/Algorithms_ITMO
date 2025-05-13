def main():
    data = input().split()
    it = iter(data)
    n = int(next(it))
    m = int(next(it))
    S = int(next(it)) - 1
    V = float(next(it))
    amount = [0.0] * n
    amount[S] = V

    # читаем описание пунктов обмена
    exch = []
    for _ in range(m):
        A = int(next(it)) - 1
        B = int(next(it)) - 1
        RAB = float(next(it))
        CAB = float(next(it))
        RBA = float(next(it))
        CBA = float(next(it))
        exch.append((A, B, RAB, CAB, RBA, CBA))

    eps = 1e-12
    # relax–итерации (m раз достаточно, потому что простая последовательность не длиннее m)
    for _ in range(m):
        updated = False
        for A, B, RAB, CAB, RBA, CBA in exch:
            # обмен A -> B
            if amount[A] > CAB + eps:
                gained = (amount[A] - CAB) * RAB
                if gained > amount[B] + eps:
                    amount[B] = gained
                    updated = True
            # обмен B -> A
            if amount[B] > CBA + eps:
                gained = (amount[B] - CBA) * RBA
                if gained > amount[A] + eps:
                    amount[A] = gained
                    updated = True
        # если уже увеличили исходную валюту — достаточно
        if amount[S] > V + eps:
            print("YES")
            return
        if not updated:
            break

    # финальная проверка на «улучшающий цикл»
    for A, B, RAB, CAB, RBA, CBA in exch:
        if amount[A] > CAB + eps and (amount[A] - CAB) * RAB > amount[B] + eps:
            print("YES")
            return
        if amount[B] > CBA + eps and (amount[B] - CBA) * RBA > amount[A] + eps:
            print("YES")
            return

    print("NO")


if __name__ == "__main__":
    main()
