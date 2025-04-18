#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

static int n;
static int64_t k;
static int *fenw;
static int logn;

static void fenw_add(int i, int v) {
    for (; i <= n; i += i & -i)
        fenw[i] += v;
}

static int fenw_sum(int i) {
    int s = 0;
    for (; i > 0; i -= i & -i)
        s += fenw[i];
    return s;
}

static int fenw_find_kth(int k_th) {
    int pos = 0;
    for (int step = logn; step > 0; step >>= 1) {
        int nxt = pos + step;
        if (nxt <= n && fenw[nxt] < k_th) {
            k_th -= fenw[nxt];
            pos = nxt;
        }
    }
    return pos + 1;
}

int main() {
    if (scanf("%d %lld", &n, &k) != 2) {
        return 0;
    }

    logn = 1;
    while ((logn << 1) <= n) {
        logn <<= 1;
    }

    fenw = malloc((n + 1) * sizeof(int));
    if (!fenw) {
        perror("malloc");
        return 1;
    }
    for (int i = 1; i <= n; i++) {
        fenw[i] = 0;
    }
    for (int i = 1; i <= n; i++) {
        fenw_add(i, 1);
    }

    int remaining = n;
    int64_t current = 0;
    for (int removed = 0; removed < n; removed++) {
        int64_t steps = (k - 1) % remaining;
        current = (current + steps) % remaining;

        int k_th = (int)(current + 1);
        int pos = fenw_find_kth(k_th);

        printf("%d\n", pos);
        fenw_add(pos, -1);


        remaining--;
    }

    free(fenw);
    return 0;
}
