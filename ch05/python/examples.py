"""
Chapter 05: Examples
"""

unsorted_list = [4, 2, 7, 1, 3]


def selection_sort(list):
    counter = 0

    for i in range(len(list)):
        lowest_number_index = i

        for j in range(i + 1, len(list)):
            counter += 1

            if list[j] < list[lowest_number_index]:
                lowest_number_index = j

        if lowest_number_index != i:
            counter += 1
            temp = list[i]
            list[i] = list[lowest_number_index]
            list[lowest_number_index] = temp

    print(f"Sorted List: {list}")
    print(f"Steps: {counter}")
    return counter


print("******** Selection Sort ********")
steps = selection_sort(unsorted_list)
assert steps == 12


def print_numbers_version_one(upper_limit):
    number = 2
    counter = 0

    while number <= upper_limit:
        counter += 1
        # If number is even, print it:
        if number % 2 == 0:
            print(number)

        number += 1

    return counter


print("******** Print Numbers Version One ********")
steps = print_numbers_version_one(100)
assert steps == 99


def print_numbers_version_two(upper_limit):
    number = 2
    counter = 0

    while number <= upper_limit:
        counter += 1
        print(number)

        # Increase number by 2, which, by definition
        # is the next even number:
        number += 2

    return counter


print("******** Print Numbers Version Two ********")
steps = print_numbers_version_two(100)
assert steps == 50