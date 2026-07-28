# Chapter 10: Recursively Recurse with Recursion

- Recursion is the term for a function calling itself.

- The case in which our function will not recurse is known as the base case.

    - Every recursive function needs at least one base case to prevent it from calling itself indefinitely.

- The computer uses a stack to keep track of which functions it's in the middle of calling.

    - This stack is known as the call stack.

        - The top element will be the most recently called function.

- In the case of infinite recursion, the computer keeps pushing the same function again and again onto the call stack.

    - The call stack grows to a point where there is no more room left in short term memory to hold all of the data.

        - This causes an error known as stack overflow.
