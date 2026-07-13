"""
Chapter 04: Examples
"""

unsorted_list = [4, 2, 7, 1, 3]


def bubble_sort(list):
    unsorted_list_index = len(unsorted_list) - 1
    is_sorted = False
    counter = 0

    while not is_sorted:
        is_sorted = True

        for i in range(unsorted_list_index):
            counter += 1
            if list[i] > list[i + 1]:
                counter += 1
                list[i], list[i + 1] = list[i + 1], list[i]
                is_sorted = False

        unsorted_list_index -= 1

    print(f"Steps: {counter}")
    print(f"Sorted List: {list}")
    return counter


print("******** O(N^2): Bubble Sort ********")
steps = bubble_sort(unsorted_list)
assert steps == 16


def quadratic_has_duplicate_values(list):
    counter = 0

    for i in range(len(list)):
        counter += 1

        for j in range(len(list)):
            counter += 1

            if i != j and list[i] == list[j]:
                print("Has Duplicate Values")
                print(f"Steps: {counter}")
                return counter
            
    print("Does Not Have Duplicate Values")
    print(f"Steps: {counter}")
    return counter


print("******** O(N^2): Duplicate Values ********")
duplicated_list = [1, 5, 3, 9, 1, 4]
steps = quadratic_has_duplicate_values(duplicated_list)
assert steps == 6


def linear_has_duplicate_values(list):
    existing_values = [0] * 10
    counter = 0

    for i in range(len(list)):
        counter += 1

        if existing_values[list[i]] == 1:
            print("Has Duplicate Values")
            print(f"Steps: {counter}")
            return counter
        else:
            existing_values[list[i]] = 1
        
    print("Does Not Have Duplicate Values")
    print(f"Steps: {counter}")
    return counter


print("******** O(N): Duplicate Values ********")
duplicated_list = [1, 5, 3, 9, 1, 4]
steps = linear_has_duplicate_values(duplicated_list)
assert steps == 5