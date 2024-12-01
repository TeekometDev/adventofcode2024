def difference(el1, el2):
    if el1 > el2:
        return el1 - el2
    return el2 - el1

f = open("input.txt", "r")
lines = f.readlines()
f.close()
list1 = []
list2 = []
for line in lines:
    symbols = line.rstrip().split(' ')
    symbols = list(filter(lambda x: x != '' and x != '\n', symbols))
    list1.append(int(symbols[0]))
    list2.append(int(symbols[1]))
list1.sort()
list2.sort()

sum = 0

for i in range(len(list1)):
    sum += difference(list1[i], list2[i])

print("Task1: " + str(sum))