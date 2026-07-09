"""
Chapter 02: Examples
"""

ordered_array = [2, 4, 6, 7, 9, 12, 14, 17, 21]


def linear_search(value):
    counter = 0

    for i in range(len(ordered_array)):
        counter += 1
        if ordered_array[i] == value:
            print(f"Steps: {counter}")
            return i, counter

    print(f"Steps: {counter}")
    return -1, counter


print("******** Linear Search ********")
index, steps = linear_search(14)
assert index == 6
assert steps == 7


def binary_search(value):
    lower_bound = 0
    upper_bound = len(ordered_array) - 1
    counter = 0

    while lower_bound <= upper_bound:
        counter += 1
        midpoint = (upper_bound + lower_bound) // 2

        value_at_midpoint = ordered_array[midpoint]

        if value == value_at_midpoint:
            print(f"Steps: {counter}")
            return midpoint, counter
        elif value < value_at_midpoint:
            upper_bound = midpoint - 1
        else:
            lower_bound = midpoint + 1

    print(f"Steps: {counter}")
    return -1, counter


print("******** Binary Search ********")
index, steps = binary_search(14)
assert index == 6
assert steps == 2


def naive_insertion(value):
    new_size = len(ordered_array) + 1
    counter = 0

    for i in range(len(ordered_array)):
        counter += 1
        if ordered_array[i] > value:
            ordered_array[len(ordered_array):len(ordered_array)] = [None] * (new_size - len(ordered_array))

            for j in range(len(ordered_array) - 1, i, -1):
                ordered_array[j] = ordered_array[j - 1]
                counter += 1

            ordered_array[i] = value
            counter += 1

            print(f"Steps: {counter}")
            return counter

    ordered_array[len(ordered_array):len(ordered_array)] = [value] * (new_size - len(ordered_array))
    counter += 1
    print(f"Steps: {counter}")
    return counter


print("******** Naive Insertion ********")
steps = naive_insertion(12)
assert steps == 11
assert ordered_array == [2, 4, 6, 7, 9, 12, 12, 14, 17, 21]


def binary_insertion(value):
    lower_bound = 0
    upper_bound = len(ordered_array) - 1
    counter = 0

    while lower_bound <= upper_bound:
        counter += 1
        midpoint = (upper_bound + lower_bound) // 2

        value_at_midpoint = ordered_array[midpoint]

        if value < value_at_midpoint:
            upper_bound = midpoint - 1
        else:
            lower_bound = midpoint + 1

    ordered_array[len(ordered_array):len(ordered_array)] = [None]

    for i in range(len(ordered_array) - 1, lower_bound, -1):
        ordered_array[i] = ordered_array[i - 1]
        counter += 1

    ordered_array[lower_bound] = value
    counter += 1

    print(f"Steps: {counter}")
    return counter


ordered_array = [2, 4, 6, 7, 9, 12, 14, 17, 21]
print("******** Binary Insertion ********")
steps = binary_insertion(12)
assert steps == 7
assert ordered_array == [2, 4, 6, 7, 9, 12, 12, 14, 17, 21]