Студент: Валуев Фаддей Денисович\
ISU: 408352


# Задача 1005. Куча камней.

## Условие

У вас есть несколько камней известного веса $w_1, …, w_n$. Напишите программу, которая распределит камни в две кучи так, что разность весов этих двух куч будет минимальной.

### Исходные данные
Ввод содержит количество камней $n$ $(1 ≤ n ≤ 20)$ и веса камней $w_1, …, w_n$ $(1 ≤ w_i ≤ 100 000)$ — целые, разделённые пробельными символами.
### Результат
Ваша программа должна вывести одно число — минимальную разность весов двух куч.


## Решение 

```go
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	rocks := make([]int, n)
	var total int
	for i := 0; i < n; i++ {
		fmt.Scan(&rocks[i])
		total += rocks[i]
	}

	minDiff := total
	for mask := 0; mask < (1 << n); mask++ {
		var sum1 int
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				sum1 += rocks[i]
			}
		}
		sum2 := total - sum1
		diff := abs(sum1 - sum2)
		if diff < minDiff {
			minDiff = diff
		}
	}

	fmt.Println(minDiff)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
```

## Объяснение алгоритма

Я решил задачу полным перебором всех вариантов размещения камней. Из-за этого сложность алгоритма получилась $O(2^n)$. 


# Задача 1322. Шпион.

## Условие

Спецслужбы обнаружили действующего иностранного агента. Шпиона то есть. Установили наблюдение и выяснили, что каждую неделю он через Интернет посылает кому-то странные нечитаемые тексты. Чтобы выяснить, к какой информации получил доступ шпион, требуется расшифровать информацию. Сотрудники спецслужб проникли в квартиру разведчика, изучили шифрующее устройство и выяснили принцип его работы.
На вход устройства подается строка текста $S1 = s_1s_2...s_N$. Получив ее, устройство строит все циклические перестановки этой строки, то есть $S2 = s_2s_3...s_Ns_1, ..., S_N = s_Ns_1s_2...s_{N-1}$. Затем множество строк $S_1, S_2, ..., S_N$ сортируется лексикографически по возрастанию. И в этом порядке строчки выписываются в столбец, одна под другой. Получается таблица размером N × N. В какой-то строке K этой таблицы находится исходное слово. Номер этой строки вместе с последним столбцом устройство и выдает на выход.
Например, если исходное слово $S_1 = abracadabra$, то таблица имеет такой вид:\
aabracadabr = $S_{11}$\
abraabracad = $S_8$\
abracadabra = $S_1$\
acadabraabr = $S_4$\
adabraabrac = $S_6$\
braabracada = $S_9$\
bracadabraa = $S_2$\
cadabraabra = $S_5$\
dabraabraca = $S_7$\
raabracadab = $S_{10}$\
racadabraab = $S_3$\
И результатом работы устройства является число 3 и строка rdarcaaaabb.
Это все, что известно про шифрующее устройство. А вот дешифрующего устройства не нашли. Но поскольку заведомо известно, что декодировать информацию можно (а иначе зачем же ее передавать?), Вам предложили помочь в борьбе с хищениями секретов и придумать алгоритм для дешифровки сообщений. А заодно и реализовать дешифратор.


### Исходные данные
В первой и второй строках находятся соответственно целое число и строка, возвращаемые шифратором. Длина строки и число не превосходят 100000. Строка содержит лишь следующие символы: a-z, A-Z, символ подчеркивания. Других символов в строке нет. Лексикографический порядок на множестве слов задается таким порядком символов:
ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz
Символы здесь выписаны в порядке возрастания.
### Результат
Выведите декодированное сообщение в единственной строке.

## Решение

Сначала я написал код на Golang, но он не уложился по времени (в Go долгая сортировка), и я переписал на python.

```go
package main

import (
	"fmt"
	"sort"
)

type charI struct {
	letter rune
	index  int
}

func main() {
	var n int
	var input string
	fmt.Scan(&n)
	fmt.Scan(&input)

	n--

	var s []charI
	for i, c := range input {
		s = append(s, charI{letter: c, index: i})
	}

	sort.SliceStable(s, func(i, j int) bool {
		return s[i].letter < s[j].letter
	})

	j := n
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[j].letter))
		j = s[j].index
	}
}
```

```python
n = int(input()) - 1
input_str = input()

s = [(char, i) for i, char in enumerate(input_str)]
s.sort(key=lambda x: x[0])

j = n
result = []
for _ in range(len(s)):
    result.append(s[j][0])
    j = s[j][1]

print("".join(result))
```


## Объяснение алгоритма

На просторах интернета, мной был найден алгоритм преобразования Барроуза–Уилера. У нас каждая строка первого столбца связана с соответствующей строкой в последнем столбце. Последний столбец нам дан, поэтому для нахождения первого столбца необходимо отсортировать ее. Однако очень важно запомнить изначальный индекс, потому что он будет задавать правило построения исходной строки. Создадим структуру, которая хранит символ и индекс этого символа в строке, подающейся на вход. Затем отсортируем строку и выведем результат. Сложность такого алгоритма -- это сложность алгоритма сортировки => $O(N\log N)$ 


# Задача 1207. Медиана на плоскости.

## Условие

На плоскости находятся $N$ точек ($N$ чётно). Никакие три точки не лежат на одной прямой. Ваша задача — выбрать две точки так, что прямая линия, проходящая через них, делит множество точек на две части одинакового размера.

### Исходные данные
Первая строка содержит целое число $N$ $(4 ≤ N ≤ 10 000)$. Каждая из следующих $N$ строк содержит пары целых чисел $x_i, y_i$ $(−10^6 ≤ xi, yi ≤ 10^6)$ — координаты $i$-й точки.
### Результат
Выведите номера выбранных точек.


## Решение

```go
package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x, y float64
	idx  int
	ang  float64
}

func main() {

	var n int
	var x, y float64
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	points := make([]Point, n)
	var sumx, sumy float64
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		points[i] = Point{x: x, y: y, idx: i + 1}
		sumx += x
		sumy += y
	}
	cx := sumx / float64(n)
	cy := sumy / float64(n)
	for i := 0; i < n; i++ {
		points[i].ang = math.Atan2(points[i].y-cy, points[i].x-cx)
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].ang < points[j].ang
	})

	halfCount := (n - 2) / 2
	eps := 1e-9

	for i := 0; i < n; i++ {
		j := (i + n/2) % n
		A := points[i]
		B := points[j]
		dx := B.x - A.x
		dy := B.y - A.y
		left := 0
		right := 0
		for k := 0; k < n; k++ {
			if k == i || k == j {
				continue
			}
			cp := dx*(points[k].y-A.y) - dy*(points[k].x-A.x)
			if cp > eps {
				left++
			} else if cp < -eps {
				right++
			}
		}
		if left == halfCount && right == halfCount {
			fmt.Println(A.idx, B.idx)
			return
		}
	}
}
```

## Объяснение алгоритма

Изначально я хотел отсортировать точки и просто взять $\frac{N}{2} и \frac{N}{2} - 1$ точки. Однако такое решение подходило не для всех случаев. Тогда я решил найти центр (точка со значениями $c_x, c_y$ -- средними значениями координат данных точек), отсортировать массив точек, но не по координатам, а по углу, относительно центра. Дальше нужно пройтись по парам точек. Для каждой $i$-й точки берем $i+\frac{N}{2}$ точку. Задаем прямую $AB$ и проверяем, делится ли плоскость пополам. Сложность алгоритма получается $O(n^2)$.


# Результаты выполнения на Тимусе

## [Ссылка на папку](https://disk.yandex.ru/d/NNbLHj17bK5frw)


# Вывод

Я узнал, как обратное преобразование BWT восстанавливает исходную строку, используя сортировку символов и переходы по индексам. Во второй задаче я понял, что сортировка точек по углу вокруг центра и выбор «противоположных» через $\frac{n}{2}$ позволяют найти прямую, делящую множество на две равные части.