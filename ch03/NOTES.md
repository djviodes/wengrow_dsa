# Chapter 3: O Yes! Big O Notation

- Big O achieves consistency by focusing on the number of steps an algorithm takes.

    - An example would be linear search taking as many steps as there are elements in an array.

        - O(N) also known as linear time.

    - Another example would be reading from a standard array which only takes one step.

        - O(1) also known as constant time.

    * If there are N data elements, how many steps will the algorithm take?

- The soul of Big O: How will an algorithm's performance change as the data increases?

- Big O generally refers to the worst-case scenario unless specified otherwise.

- In Big O terms, we describe binary search as having a time complexity of O(logN).

    - Otherwise known as log time.
    - Simply put, O(logN) describes an algorithm that increases one step each time the data is doubled.
    - O(logN) is really shorthand for O(log_2N)

        | N Elements | O(N) | O(logN) |
        |---|---|---|
        | 8 | 8 | 3 |
        | 16 | 16 | 4 |
        | 32 | 32 | 5 |
        | 64 | 64 | 6 |
        | 128 | 128 | 7 |
        | 256 | 256 | 8 |
        | 512 | 512 | 9 |
        | 1024 | 1024 | 10 |
