"""
Chapter 08: Exercises
"""


def list_intersection(list_1, list_2):
    larger_list = []
    smaller_list = []
    intersections = []
    hash_table = {}

    if len(list_1) > len(list_2):
        larger_list = list_1
        smaller_list = list_2
    else:
        larger_list = list_2
        smaller_list = list_1

    for num in larger_list:
        hash_table[num] = True

    for num in smaller_list:
        if num in hash_table:
            intersections.append(num)

    return intersections


print("******** List Intersection ********")
intersection = list_intersection([1, 2, 3, 4, 5], [0, 2, 4, 6, 8])
print(f"Intersection: {intersection}")
assert intersection == [2, 4]


def duplicate_chars(list):
    hash_table = {}

    for char in list:
        if char in hash_table:
            return char
        else:
            hash_table[char] = True

    return "N/A"


print("******** Duplicate Chars ********")
duplicate = duplicate_chars(["a", "b", "c", "d", "c", "e", "f"])
print(f"Duplicate: {duplicate}")
assert duplicate == "c"


def missing_letter(i_string):
    hash_table = {
        'a': False, 'b': False, 'c': False, 'd': False, 'e': False,
        'f': False, 'g': False, 'h': False, 'i': False, 'j': False,
        'k': False, 'l': False, 'm': False, 'n': False, 'o': False,
        'p': False, 'q': False, 'r': False, 's': False, 't': False,
        'u': False, 'v': False, 'w': False, 'x': False, 'y': False,
        'z': False
    }

    for char in i_string:
        hash_table[char] = True

    for key, value in hash_table.items():
        if value == False:
            return key

    return "N/A"


print("******** Missing Letter ********")
missing = missing_letter("the quick brown box jumps over the lazy dog")
print(f"Missing Letter: {missing}")
assert missing == "f"


def non_duplicated_string(i_string):
    lowest_index_single_char = 0
    hash_table = {}

    for char in i_string:
        if char not in hash_table:
            hash_table[char] = 1
        else:
            hash_table[char] += 1

    while lowest_index_single_char < len(i_string):
        if hash_table[i_string[lowest_index_single_char]] == 1:
            return i_string[lowest_index_single_char]
        else:
            lowest_index_single_char += 1

    return "N/A"


print("******** Non-Duplicated String ********")
non_duplicated = non_duplicated_string("minimum")
print(f"Non-Duplicated: {non_duplicated}")
assert non_duplicated == "n"
