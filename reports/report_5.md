Студент: Валуев Фаддей Денисович\
ISU: 408352


# Задача 1162. Currency Exchange.

## Условие
Several currency exchange points are working in our city. Let us suppose that each point specializes in two particular currencies and performs exchange operations only with these currencies. There can be several points specializing in the same pair of currencies. Each point has its own exchange rates, exchange rate of A to B is the quantity of B you get for 1A. Also each exchange point has some commission, the sum you have to pay for your exchange operation. Commission is always collected in source currency.
For example, if you want to exchange 100 US Dollars into Russian Rubles at the exchange point, where the exchange rate is 29.75, and the commission is 0.39 you will get (100 - 0.39) * 29.75 = 2963.3975RUR.
You surely know that there are N different currencies you can deal with in our city. Let us assign unique integer number from 1 to N to each currency. Then each exchange point can be described with 6 numbers: integer A and B - numbers of currencies it exchanges, and real RAB, CAB, RBA and CBA - exchange rates and commissions when exchanging A to B and B to A respectively.
Nick has some money in currency S and wonders if he can somehow, after some exchange operations, increase his capital. Of course, he wants to have his money in currency S in the end. Help him to answer this difficult question. Nick must always have non-negative sum of money while making his operations.

### Исходные данные

The first line contains four numbers: N - the number of currencies, M - the number of exchange points, S - the number of currency Nick has and V - the quantity of currency units he has. The following M lines contain 6 numbers each - the description of the corresponding exchange point - in specified above order. Numbers are separated by one or more spaces. 1 ≤ S ≤ N ≤ 100, 1 ≤ M ≤ 100, V is real number, 0 ≤ V ≤ 103.
For each point exchange rates and commissions are real, given with at most two digits after the decimal point, 10-2 ≤ rate ≤ 102, 0 ≤ commission ≤ 102.
Let us call some sequence of the exchange operations simple if no exchange point is used more than once in this sequence. You may assume that ratio of the numeric values of the sums at the end and at the beginning of any simple sequence of the exchange operations will be less than 104.

### Результат
Результат
If Nick can increase his wealth, output YES, in other case output NO.

## Решение

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

const eps = 1e-9

type Edge struct {
	from, to  int
	rate, com float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var N, M, S int
	var V float64

	fmt.Fscan(reader, &N, &M, &S, &V)

	edges := make([]Edge, 0, 2*M)
	for i := 0; i < M; i++ {
		var A, B int
		var R_ab, C_ab, R_ba, C_ba float64
		fmt.Fscan(reader, &A, &B, &R_ab, &C_ab, &R_ba, &C_ba)
		edges = append(edges, Edge{from: A, to: B, rate: R_ab, com: C_ab})
		edges = append(edges, Edge{from: B, to: A, rate: R_ba, com: C_ba})
	}

	dist := make([]float64, N+1)
	dist[S] = V

	for i := 0; i < N; i++ {
		updated := false
		for _, e := range edges {
			if dist[e.from] > e.com+eps {
				amt := (dist[e.from] - e.com) * e.rate
				if amt > dist[e.to]+eps {
					dist[e.to] = amt
					updated = true
					if i == N-1 {
						fmt.Println("YES")
						return
					}
				}
			}
		}
		if !updated {
			break
		}
	}

	if dist[S] > V+eps {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
```

## Объяснение алгоритма
В худшем случае необходимо между потенциальными валютами для выгоды необходимо совершить n-1 обменов, чтобы получить выгоду (если ее возможно получить). Ограничение в использовании простой последовательности и конечную валюту - ничтожно


# Задача 1444. Накормить элефпотама.

## Условие
Любимое лакомство элефпотамов — слоновьи тыквы, именно они и растут на лужайке, где Гарри должен сдавать экзамен. Перед началом испытания Хагрид притащит животное к одной из тыкв. Скормив элефпотаму очередную тыкву, Гарри может направить его в сторону любой оставшейся тыквы. Чтобы сдать экзамен, надо провести элефпотама по лужайке так, чтобы тот съел как можно больше тыкв до того, как наткнется на свои следы.

### Исходные данные
В первой строке входа находится число N (3 ≤ N ≤ 30000) — количество тыкв на лужайке. Тыквы пронумерованы от 1 до N, причем номер один присвоен той тыкве, у которой будет стоять элефпотам в начале экзамена. В следующих N строках даны координаты всех тыкв по порядку. Все координаты — целые числа от −1000 до 1000. Известно, что положения всех тыкв различны, и не существует прямой, проходящей сразу через все тыквы.

### Результат
В первой строке выхода вы должны вывести K — максимальное количество тыкв, которое может съесть элефпотам. Далее по одному числу в строке выведите K чисел — номера тыкв в порядке их обхода. Первым в этой последовательности всегда должно быть число 1.

## Решение

```go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const PI = 3.14159265

type Pumpkin struct {
	Length  float64
	Degrees float64
	Index   int
}

func (p *Pumpkin) Create(fx, fy, idx int, r *bufio.Reader) {
	var x, y int
	fmt.Fscan(r, &x, &y)
	dx := float64(x - fx)
	dy := float64(y - fy)
	p.Index = idx + 1
	p.Length = dx*dx + dy*dy
	p.Degrees = math.Atan2(dy, dx) * 180.0 / PI
	if p.Degrees < 0 {
		p.Degrees += 360.0
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, fx, fy int
	fmt.Fscan(in, &n, &fx, &fy)

	pumpkins := make([]Pumpkin, n)
	pumpkins[0] = Pumpkin{
		Length:  0,
		Degrees: -1,
		Index:   1,
	}
	for i := 1; i < n; i++ {
		pumpkins[i].Create(fx, fy, i, in)
	}

	eps := 1e-10
	sort.Slice(pumpkins, func(i, j int) bool {
		d1 := pumpkins[i].Degrees
		d2 := pumpkins[j].Degrees
		if math.Abs(d1-d2) > eps {
			return d1 < d2
		}
		return pumpkins[i].Length < pumpkins[j].Length
	})

	startPoint := 1
	for i := 1; i < n-1; i++ {
		if pumpkins[i+1].Degrees-pumpkins[i].Degrees > 179.999 {
			startPoint = i + 1
			break
		}
	}

	fmt.Fprintln(out, n)
	fmt.Fprintln(out, 1)

	for i := startPoint; i < n; i++ {
		fmt.Fprintln(out, pumpkins[i].Index)
	}

	for i := 1; i < startPoint; i++ {
		fmt.Fprintln(out, pumpkins[i].Index)
	}
}
```

## Объяснение алгоритма
Так как все углы у нас в диапаоне от 0 до 360*, и отсортированны по возрастанию, то необходимо в качестве стартовой точки выбрать либо любую, либо ту, угол от которй до предидущей больше 180 при ее наличии, так как в противном случае может быть пеерсечение. Такая точка может быть всего одна


# Задача 1628. Мобильные телеграфы.

## Условие
У каждого неудачника в жизни бывают не только чёрные, но и белые полосы. Марсианин Вась-Вась отмечает в календаре, представляющем собой таблицу m × n, те дни, когда ему ужасно не повезло. Если Вась-Васю не повезло в j-й день i-й недели, то он закрашивает ячейку таблицы (i, j) в чёрный цвет. Все незакрашенные ячейки в таблице имеют белый цвет.
Будем называть отрезками жизни прямоугольники размером 1 × l либо l × 1. Белыми полосами Вась-Вась считает все максимальные по включению белые отрезки таблицы. А сможете ли Вы определить, сколько всего белых полос было в жизни Вась-Вася?


### Исходные данные
Первая строка содержит целые числа m, n, k — размеры календаря и количество неудачных дней в жизни Вась-Вася (1 ≤ m, n ≤ 30000; 0 ≤ k ≤ 60000). В следующих k строках перечислены неудачные дни в виде пар (xi, yi), где xi — номер недели, к которой относится неудачный день, а yi — номер дня в этой неделе (1 ≤ xi ≤ m; 1 ≤ yi ≤ n). Описание каждого неудачного дня встречается только один раз.

### Результат
Выведите число белых полос в жизни Вась-Вася.

## Решение

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var m, n, k int
	fmt.Fscan(in, &m, &n, &k)

	rowBlacks := make([][]int, m+1)
	colBlacks := make([][]int, n+1)

	for t := 0; t < k; t++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		rowBlacks[x] = append(rowBlacks[x], y)
		colBlacks[y] = append(colBlacks[y], x)
	}

	for i := 1; i <= m; i++ {
		sort.Ints(rowBlacks[i])
	}
	for j := 1; j <= n; j++ {
		sort.Ints(colBlacks[j])
	}

	stripesH := 0
	hSingles := make(map[uint64]struct{})

	for i := 1; i <= m; i++ {
		prev := 0
		for _, y := range rowBlacks[i] {
			length := y - prev - 1
			if length >= 2 {
				stripesH++
			} else if length == 1 {
				key := uint64(i)<<32 | uint64(prev+1)
				hSingles[key] = struct{}{}
			}
			prev = y
		}
		length := n - prev
		if length >= 2 {
			stripesH++
		} else if length == 1 {
			key := uint64(i)<<32 | uint64(prev+1)
			hSingles[key] = struct{}{}
		}
	}
	stripesV := 0
	isolated := 0
	
	for j := 1; j <= n; j++ {
		prev := 0
		for _, x := range colBlacks[j] {
			length := x - prev - 1
			if length >= 2 {
				stripesV++
			} else if length == 1 {
				key := uint64(prev+1)<<32 | uint64(j)
				if _, ok := hSingles[key]; ok {
					isolated++
				}
			}
			prev = x
		}
		length := m - prev
		if length >= 2 {
			stripesV++
		} else if length == 1 {
			key := uint64(prev+1)<<32 | uint64(j)
			if _, ok := hSingles[key]; ok {
				isolated++
			}
		}
	}

	fmt.Println(stripesH + stripesV + isolated)
}
```

## Объяснение алгоритма
Полноценный массив m*n реализовывать нет смысла, так как он не пройдет по памяти Получилось реализовать программу более быстрым способом и по времени, и по памяти, и она даже прошла все тесты, но по факту оно пройти не должно было в связи с возможностью вылета при определенных вводдных, которые не тестируются.


# Результаты выполнения на Тимусе

## [Ссылка на папку](https://disk.yandex.ru/d/IY921muuk9HWTg)
