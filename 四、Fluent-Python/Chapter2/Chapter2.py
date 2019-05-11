# fruits = ['grape', 'raspberry', 'apple', 'banana']
#
# print(sorted(fruits))
# print(fruits)
#
# print(sorted(fruits, reverse=True))
# print(sorted(fruits, key=len))
# print(sorted(fruits, key=len, reverse=True))
#
# print(fruits)
# fruits.sort()
# print(fruits)


'''2-17'''
import bisect

haystack = [1, 4, 5, 6, 8, 11, 20, 29, 49, 87]
needle = 24
index = bisect.bisect(haystack, needle)
haystack.insert(index, needle)

print(index, '\n', haystack)


