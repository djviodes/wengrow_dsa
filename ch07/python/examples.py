"""
Chapter 07: Examples
"""

def average_of_even_numbers(list):
    sum = 0.0
    count_of_even_numbers = 0
    counter = 0

    for number in list:
        counter += 1

        if number % 2 == 0:
            counter += 1
            sum += number
            count_of_even_numbers += 1

    print(f"Average: {sum / count_of_even_numbers}")
    print(f"Steps: {counter}")
    return counter


print("******** Average of Even Numbers ********")
numbers_list = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
steps = average_of_even_numbers(numbers_list)
assert steps == 15


def word_builder(list):
    collection = []
    counter = 0

    for i in range(len(list)):
        for j in range(len(list)):
            counter += 1

            if i != j:
                counter += 1
                collection.append(list[i] + list[j])

    print(f"Words: {collection}")
    print(f"Steps: {counter}")
    return counter


print("******** Word Builder ********")
words_list = ["a", "b", "c", "d"]
steps = word_builder(words_list)
assert steps == 28


def list_sample(list):
    first = list[0]
    middle = list[int(len(list) / 2)]
    last = list[-1]
    counter = 0

    counter += 1

    print(f"List Sample: {[first, middle, last]}")
    print(f"Steps: {counter}")
    return counter


print("******** List Sample ********")
full_list = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
steps = list_sample(full_list)
assert steps == 1


def average_celcius(fahrenheit_readings):
    celsius_numbers = []
    counter = 0

    for fahrenheit_reading in fahrenheit_readings:
        counter += 1

        celsius_conversion = (fahrenheit_reading - 32) / 1.8
        counter += 1

        celsius_numbers.append(celsius_conversion)
        counter += 1

    sum = 0.0

    for celsius_number in celsius_numbers:
        counter += 1

        sum += celsius_number
        counter += 1

    print(f"Mean Average: {sum / len(celsius_numbers)}")
    print(f"Steps: {counter}")
    return counter


print("******** Average Celcius ********")
fahrenheit_readings = [82.0, 81.2, 85.3, 84.0, 78.6, 79.9, 80.7, 75.1]
steps = average_celcius(fahrenheit_readings)
assert steps == 40


def mark_inventory(clothing_items):
    clothing_options = []
    counter = 0

    for item in clothing_items:
        for size in range(1, 6):
            counter += 1
            clothing_options.append(item + " Size: " + str(size))

    print(f"Clothing Options: {clothing_options}")
    print(f"Steps: {counter}")
    return counter


print("******** Mark Inventory ********")
clothing_items = ["Purple Shirt", "Green Shirt"]
steps = mark_inventory(clothing_items)
assert steps == 10


def count_ones(matrix):
    ones_count = 0
    counter = 0

    for vector in matrix:
        for num in vector:
            counter += 1

            if num == 1:
                ones_count += 1

    print(f"Ones Count: {ones_count}")
    print(f"Steps: {counter}")
    return counter


print("******** Count Ones ********")
matrix = [
    [0, 1, 0, 1],
    [1, 0, 1, 0],
    [0, 0, 1, 1],
    [1, 1, 0, 0]
]
steps = count_ones(matrix)
assert steps == 16


def is_palindrome(string):
    left_index = 0
    right_index = len(string) - 1
    counter = 0

    while left_index < len(string) / 2:
        counter += 1

        if string[left_index] != string[right_index]:
            print(f"Is Palindrome: {False}")
            print(f"Steps: {counter}")
            return counter
        
        left_index += 1
        right_index -= 1

    print(f"Is Palindrome: {True}")
    print(f"Steps: {counter}")
    return counter


print("******** Is Palindrome ********")
palindrome = "racecar"
steps = is_palindrome(palindrome)
assert steps == 4


def two_number_products(list):
    products = []
    counter = 0

    for i in range(len(list)):
        for j in range(len(list)):
            counter += 1
            products.append(list[i] * list[j])

    print(f"Products: {products}")
    print(f"Steps: {counter}")
    return counter


print("******** Two Number Products ********")
num_list = [1, 2, 3, 4, 5]
steps = two_number_products(num_list)
assert steps == 25