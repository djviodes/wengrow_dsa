"""
Chapter 06: Examples
"""

unsorted_list = [4, 2, 7, 1, 3]


def insertion_sort(list):
    counter = 0

    for i in range(1, len(list)):
        counter += 1
        temp = list[i]
        position = i - 1

        while position >= 0:
            counter += 1

            if list[position] > temp:
                counter += 1
                list[position + 1] = list[position]
                position -= 1
            else:
                break

        counter += 1
        list[position + 1] = temp

    print(f"Sorted List: {list}")
    print(f"Steps: {counter}")
    return counter


print("******** Insertion Sort ********")
steps = insertion_sort(unsorted_list)
assert steps == 22


def nonbreak_intersection(first_list, second_list):
    result = []
    counter = 0

    for i in range(len(first_list)):
        for j in range(len(second_list)):
            counter += 1

            if first_list[i] == second_list[j]:
                counter += 1
                result.append(first_list[i])

    print(f"Result: {result}")
    print(f"Steps: {counter}")
    return counter


print("******** Nonbreak Intersection ********")
steps = nonbreak_intersection([3, 1, 4, 2], [4, 5, 3, 6])
assert steps == 18


def break_intersection(first_list, second_list):
    result = []
    counter = 0

    for i in range(len(first_list)):
        for j in range(len(second_list)):
            counter += 1

            if first_list[i] == second_list[j]:
                counter += 1
                result.append(first_list[i])
                break

    print(f"Result: {result}")
    print(f"Steps: {counter}")
    return counter


print("******** Break Intersection ********")
steps = break_intersection([3, 1, 4, 2], [4, 5, 3, 6])
assert steps == 14