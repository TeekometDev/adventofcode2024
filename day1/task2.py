def count_apperance(number, list):
    count = 0
    for num in list:
        if num == number:
            count += 1
    return count

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

sum = 0
for num in list1:
    if (num != 0):
        sum += num * count_apperance(num, list2)

print("Task2: " + str(sum))