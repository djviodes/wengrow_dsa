"""
Chapter 06: Exercises
"""


def worst_case_contains_x(string):
    found_x = False
    counter = 0

    for i in range(len(string)):
        counter += 1

        if string[i] == "X":
            found_x = True

    print(f"Found X: {found_x}")
    print(f"Steps: {counter}")
    return counter


print("******** Worst Case Contains X ********")
x_string = "TestX"
steps = worst_case_contains_x(x_string)
assert steps == 5


def best_and_average_case_contains_x(string):
    counter = 0

    for i in range(len(string)):
        counter += 1

        if string[i] == "X":
            print("Found X: True")
            print(f"Steps: {counter}")
            return counter

    print("Found X: False")
    print(f"Steps: {counter}")
    return counter


print("******** Best Case Contains X ********")
x_string = "XTest"
steps = best_and_average_case_contains_x(x_string)
assert steps == 1


print("******** Average Case Contains X ********")
x_string = "TeXst"
steps = best_and_average_case_contains_x(x_string)
assert steps == 3