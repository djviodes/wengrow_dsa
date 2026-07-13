# Chapter 4: Speeding Up Your Code With Big O

- If you find that Big O labels your algorithm as slow, you can now take a step back and try to figure out if there's a way to optimize it by trying to get it to fall under a faster category of Big O.

- Sorting algorithms have been the subject of extensive research in computer science.

    - They solve the problem:

        - Given an array of unsorted values, how can we get them so they end up in ascending order?

- Bubble Sort is a basic sorting algorithm and follows these steps:

    1. Point to two consecutive values in the array and compare the first item with the second.

        ```
        | 2 | 1 | 3 | 5 |
          ^   ^
        ```

    2. If the two items are out of order, swap them.

        ```
        | 1 | 2 | 3 | 5 |
          ^   ^
        ```

    3. Move the "pointers" one cell to the right.

        ```
        | 1 | 2 | 3 | 5 |
              ^   ^
        ```

    4. Repeat steps 1 through 3 until we reach the end of the array, or if we reach the values that have already been sorted.

    5. We then move the two pointers back to the first two values of the array, and execute another pass-through of the array by running steps 1 through 4 again.

        - We keep on executing these pass-throughs until we have a pass through in which we did not perform any swaps.

- Bubble Sort takes N^2 steps, in Big O, we say that Bubble Sort has an efficiency of O(N^2)

    - | N Data Elements | # of Bubble Sort Steps | N^2 |
      |---|---|---|
      | 5 | 20 | 25 |
      | 10 | 90 | 100 |
      | 20 | 380 | 400 |
      | 40 | 1560 | 1600 |
      | 80 | 6320 | 6400 |