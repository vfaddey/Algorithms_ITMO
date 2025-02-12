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