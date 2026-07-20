"""
Chapter 08: Examples
"""

list_1 = [1, 2, 3, 4, 5, 6]
list_2 = [2, 4, 6]


def no_hash_is_subset(list_1, list_2):
    larger_list = []
    smaller_list = []

    # Determine which array is smaller:
    if len(list_1) > len(list_2):
        larger_list = list_1
        smaller_list = list_2
    else:
        larger_list = list_2
        smaller_list = list_1

    # Iterate through smaller array
    for i in range(len(smaller_list)):
        # Assume temporarily that the current value from
        # smaller array is not found in larger array
        found_match = False

        # For each value in smaller array, iterate through
        # larger array:
        for j in range(len(larger_list)):
            # If the two values are equal, it means the current
            # value in smaller array is present in the larger array:
            if smaller_list[i] == larger_list[j]:
                found_match = True
                break

        # If the current value in smaller array doesn't exist
        # in larger array, return false:
        if found_match == False:
            return False
    
    # If we get tot he end of the loops, it means that all
    # values from smaller array are present in larger array:
    return True


print("******** No Hash Table Is Subset ********")
is_subset = no_hash_is_subset(list_1, list_2)
assert is_subset == True


def hash_is_subset(list_1, list_2):
    larger_list = []
    smaller_list = []
    hash_table = {}

    # Determine which array is smaller:
    if len(list_1) > len(list_2):
        larger_list = list_1
        smaller_list = list_2
    else:
        larger_list = list_2
        smaller_list = list_1

    # Store all items from larger_list inside hash_table:
    for num in larger_list:
        hash_table[num] = True

    # Iterate through each item in smaller_list and return False
    # if we encounter an item not inside hash_table:
    for num in smaller_list:
        if num not in hash_table:
            return False
        
    # If we got this far in our code without returning False,
    # it means that all the items in the smaller_list
    # must be contained within larger_list
    return True


print("******** Hash Table Is Subset ********")
is_subset = hash_is_subset(list_1, list_2)
assert is_subset == True