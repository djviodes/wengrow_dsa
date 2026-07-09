"""
Chapter 01: Examples
"""

grocery_list = ["apples", "bananas", "cucumbers", "dates", "elderberries"]
print(f"Size: {len(grocery_list)}")
print(f"Index 0: {grocery_list[0]} | Index 1: {grocery_list[1]} | Index 2: {grocery_list[2]} | Index 3: {grocery_list[3]} | Index 4: {grocery_list[4]}")


def search_for_value(value):
    counter = 0

    for i in range(len(grocery_list)):
        counter += 1
        if grocery_list[i] == value:
            print(f"Steps: {counter}")
            return i, counter

    print(f"Steps: {counter}")
    return -1, counter


print("******** Search for Value ********")
index, steps = search_for_value("elderberries")
assert index == 4
assert steps == 5


def insert_at_beginning(item):
    new_size = len(grocery_list) + 1
    grocery_list[len(grocery_list):len(grocery_list)] = [None] * (new_size - len(grocery_list))
    counter = 0

    for i in range(len(grocery_list) - 1, 0, -1):
        grocery_list[i] = grocery_list[i - 1]
        counter += 1

    grocery_list[0] = item
    counter += 1

    print(f"Steps: {counter}")
    return counter


print("******** Insert at Beginning ********")
steps = insert_at_beginning("figs")
assert steps == 6
assert grocery_list == ["figs", "apples", "bananas", "cucumbers", "dates", "elderberries"]
print(f"Size: {len(grocery_list)}")
print(f"Index 0: {grocery_list[0]} | Index 1: {grocery_list[1]} | Index 2: {grocery_list[2]} | Index 3: {grocery_list[3]} | Index 4: {grocery_list[4]} | Index 5: {grocery_list[5]}")


grocery_list = ["apples", "bananas", "cucumbers", "dates", "elderberries"]


def insert_at_end(item):
    new_size = len(grocery_list) + 1
    grocery_list[len(grocery_list):len(grocery_list)] = [item] * (new_size - len(grocery_list))

    print("Steps: 1")
    return 1


print("******** Insert at End ********")
steps = insert_at_end("pomegranates")
assert steps == 1
assert grocery_list == ["apples", "bananas", "cucumbers", "dates", "elderberries", "pomegranates"]
print(f"Size: {len(grocery_list)}")
print(f"Index 0: {grocery_list[0]} | Index 1: {grocery_list[1]} | Index 2: {grocery_list[2]} | Index 3: {grocery_list[3]} | Index 4: {grocery_list[4]} | Index 5: {grocery_list[5]}")


grocery_list = ["apples", "bananas", "cucumbers", "dates", "elderberries"]


def delete_at_beginning():
    counter = 0

    for i in range(0, len(grocery_list) - 1):
        grocery_list[i] = grocery_list[i + 1]
        counter += 1

    grocery_list[:] = grocery_list[:len(grocery_list) - 1]
    counter += 1

    print(f"Steps: {counter}")
    return counter


print("******** Delete at Beginning ********")
steps = delete_at_beginning()
assert steps == 5
assert grocery_list == ["bananas", "cucumbers", "dates", "elderberries"]
print(f"Size: {len(grocery_list)}")
print(f"Index 0: {grocery_list[0]} | Index 1: {grocery_list[1]} | Index 2: {grocery_list[2]} | Index 3: {grocery_list[3]}")


grocery_list = ["apples", "bananas", "cucumbers", "dates", "elderberries"]


def delete_at_end():
    grocery_list[:] = grocery_list[:len(grocery_list) - 1]

    print("Steps: 1")
    return 1


print("******** Delete at End ********")
steps = delete_at_end()
assert steps == 1
assert grocery_list == ["apples", "bananas", "cucumbers", "dates"]
print(f"Size: {len(grocery_list)}")
print(f"Index 0: {grocery_list[0]} | Index 1: {grocery_list[1]} | Index 2: {grocery_list[2]} | Index 3: {grocery_list[3]}")
