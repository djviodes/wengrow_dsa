"""
Chapter 04: Exercises
"""

def quadratic_greatest_number(list):
    counter = 0

    for i in list:
        # Assume for now that i is the greatest:
        is_i_val_the_greatest = True

        for j in list:
            counter += 1
            # If we find another value that is greater than i,
            # i is not the greatest:
            if j > i:
                is_i_val_the_greatest = False

        # If, by the time we checked all the other numbers, i
        # is still the greatest, it means that i is the greatest number:
        if is_i_val_the_greatest:
            print(f"Greatest Number: {i}")
            print(f"Steps: {counter}")
            return counter
        

print("******** O(N^2): Greatest Number ********")
number_list = [3, 1, 6, 8, 2, 9]
steps = quadratic_greatest_number(number_list)
assert steps == 36


def linear_greatest_number(list):
    greatest_number = list[0]
    counter = 0

    for i in range(len(list)):
        counter += 1

        if list[i] > greatest_number:
            greatest_number = list[i]

    print(f"Greatest Number: {greatest_number}")
    print(f"Steps: {counter}")
    return counter


print("******** O(N): Greatest Number ********")
number_list = [3, 1, 6, 8, 2, 9]
steps = linear_greatest_number(number_list)
assert steps == 6