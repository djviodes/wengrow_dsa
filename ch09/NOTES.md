# Chapter 9: Crafting Elegant Code with Stacks and Queues

- Stacks and queues serve as temporary containers that can be used to form beautiful algorithms.

- A stack stores data in the same way arrays do but with three constraints:

    1. Data can be inserted only at the end of the stack.
    2. Data can be deleted only from the end of the stack.
    3. Only the last element of the stack can be read.

- Stack operations subscribe to the acronym LIFO (Last In First Out).

    - Inserting a new value into the stack is also called pushing onto the stack.
    - Removing elements from the top of the stack is called popping from the stack.

- The stack is an example of what is known as an abstract data type:

    - A data structure that is a set of theoretical rules that revolve around some other built-in data structure.

- A simple implementation of a stack could be the beginnings of a programming language's linter:

    - We prepare an empty stack, and then we read each character from left to right following these rules:

        1. If we find any character that isn't a type of brace, we ignore it and move on.
        2. If we find an opening brace, we push it onto the stack.
        3. If we find a closing brace, we pop the top element in the stack and inspect it.

            - If we couldn't pop an element because the stack was empty, or the popped element does not match the closing brace, then we have a syntax error.
            - If the popped element is a corresponding match for the current closing brace, then we've successfully closed the opening brace, and we can continue parsing the code.

        4. If we make it to the end of the line of code and there is still something left on the stack, then we have a syntax error.

- Constrained data structures like the stack are important for several reasons:

    - We can prevent potential bugs.
    - They give us a new mental model for tackling problems.

- A queue is another data structure designed to process temporary data.

    - Like stacks, they are arrays with three restrictions:

        1. Data can only be inserted at the end of the queue.
        2. Data can only be removed from the beginning of the queue.
        3. Only the element at the front of the queue can be read.
