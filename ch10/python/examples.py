"""
Chapter 10: Examples
"""


def countdown(number):
    counter = 0

    for i in range(number, 0, -1):
        counter += 1
        print(i)

    print(f"Steps: {counter}")
    return counter


print("******** Countdown ********")
steps = countdown(10)
assert steps == 10


def recursive_countdown(number, counter):
    if number > 0:
        print(number)
        counter += 1
        number -= 1
        return recursive_countdown(number, counter)
    else:
        print(f"Steps: {counter}")
        return counter


print("******** Recursive Counter ********")
steps = recursive_countdown(10, 0)
print(steps)
assert steps == 10


def factorial(number):
    if number == 1:
        return 1
    else:
        return number * factorial(number - 1)


print("******** Factorial ********")
factorial_sum = factorial(5)
print(f"5 Factorial = {factorial_sum}")
assert factorial_sum == 120