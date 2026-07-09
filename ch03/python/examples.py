"""
Chapter 03: Examples
"""

things = ["apples", "baboons", "cribs", "dulcimers"]


def print_things():
    counter = 0

    for thing in things:
        counter += 1
        print("Here's a thing: %s" % thing)

    print(f"Steps: {counter}")
    return counter


print("******** O(N): Print Each Thing ********")
steps = print_things()
assert steps == 4


def is_prime(number):
    counter = 0

    for i in range(2, number):
        counter += 1
        if number % i == 0:
            print(f"Steps: {counter}")
            return False, counter

    print(f"Steps: {counter}")
    return True, counter


print("******** O(N): Is Prime ********")
result, steps = is_prime(7)
assert result is True
assert steps == 5
