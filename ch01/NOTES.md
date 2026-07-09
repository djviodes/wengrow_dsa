# Chapter 1: Why Data Structures Matter

- Efficiency is one of many measures of code quality and the first step in writing fast, efficient code is to understand what data structures are and how different data structures affect the speed of our code.

- Data is a broad term that refers to all types of information, down to the most basic numbers and strings.

- Data structures refer to how data is organized.

    - Data organization can significantly impact the efficiency of your code.

- The Array: The Foundational Data Structure

    - The _size_ of an array is how many data elements the array holds.
    - The _index_ of an array is the number that identifies where a piece of data lives inside the array.

        - Ex:

            ```
            | "apples" | "bananas" | "cucumbers" | "dates" | "elderberries" |   size = 5
              index 0     index 1      index 2     index 3      index 4
            ```

- Many data structures are used in four basic ways, which we refer to as operations:

    1. Read
    2. Search
    3. Insert
    4. Delete

- When we measure how "fast" an operation takes, we refer to how many steps it takes.

    - Measuring speed is also known as measuring the time complexity.

- A computer can read from an array in just one step.

    - When the computer reads a value at a particular index of an array, it can jump straight to that index because of the combination of the following about computers:

        1. A computer can jump to any memory address in one step.
        2. Whenever a computer allocates an array, it also makes note at which memory address the array begins.

- To search for a value within the array, the computer has no choice but to inspect each cell one at a time.

    - This is known as linear search.

        - Linear search takes a maximum of N steps where N represents the total number of cells in an array.

- The efficiency of inserting a new piece of data into an array depends on where within the array you're inserting it.

    - Inserting at the end of an array takes just one step since the computer keeps track of the array's size.
    - Inserting at the beginning or middle of an array requires shifting the following cells down one.

        - A worst-case scenario can take N + 1 steps for an array containing N elements.

- The efficiency of deleting a piece of data from an array follows the time complexity of inserting for the same reasons.

    - Deleting at the end of the array takes one step.
    - Deleting at the beginning of the array takes N steps.

- Sets: How a Single Rule Can Affect Efficiency

    - A set is a data structure that does not allow duplicate values to be contained within it.

        - Useful when you need to ensure that you don't have duplicate data.
        - For an array based set, this one constraint causes the set to have a different efficiency for one of the four primary operations.

            - Every insertion into a set first requires a search.

                - This takes N + 1 steps if we're inserting into the end of a set, and 2N + 1 if we're inserting into the beginning.
