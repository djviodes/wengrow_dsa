"""
Chapter 09: Exercises
"""

from examples import Stack


def reverse_string(string):
    stack = Stack()

    for char in string:
        stack.push(char)

    reversed_chars = []

    while stack.read():
        reversed_chars.append(stack.pop())

    return "".join(reversed_chars)


print("******** Reverse String ********")
reversed_str = reverse_string("hello")
print(f"Reversed: {reversed_str}")
assert reversed_str == "olleh"
