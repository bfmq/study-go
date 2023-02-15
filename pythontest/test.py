list20 = (x for x in range(21))
even20 = [print(x) for x in list20 if x % 2 == 0]


from itertools import product
res = 0
for x, y, z in product(range(1, 5), range(1, 5), range(1, 5)):
    if x != y and x != z and y != z:
        res += 1
        print("%d%d%d", x, y, z)
print(res)


from collections import defaultdict
def list2map(source):
    res = defaultdict(int)
    for i in source:
        res[i] += 1
    return res
