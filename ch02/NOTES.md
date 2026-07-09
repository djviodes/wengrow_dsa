# Chapter 2: Why Algorithms Matter

- An algorithm is simply a set of instructions for completing a specific task.

    - It is possible to have two different algorithms that accomplish the same task.

- An ordered array is an array that requires the values to be kept in order.

    - Ex:

        ```
        | 3 | 17 | 80 | 202 |
        ```

        If we want to drop 75 into this array, the computer must:

        1. Find the right place to insert the 75
        2. Shift the other values to make room for it

            ```
            | 3 | 17 | 80 |    | 202 |
            ```

            ```
            | 3 | 17 |    | 80 | 202 |
            ```

            ```
            | 3 | 17 | 75 | 80 | 202 |
            ```

    - For N elements in an ordered array, the insertion took N + 2 steps.
    - In certain situations a linear search over an ordered array can take fewer steps since we can stop when we reach a value larger than the target value.

        - If the target value is the final value, or larger than the final value, then we could still end up taking N steps.

    - Binary search is a method that allows for quicker search of something by splitting the search interval into two.

        - Ex:

            Say we have an ordered array containing 9 unknown elements and we are looking for the value 7. Here's how binary search would work:

            ```
            | ? | ? | ? | ? | ? | ? | ? | ? | ? |
            ```

            Step 1: Begin with the central cell and check its value

            ```
            | ? | ? | ? | ? | 9 | ? | ? | ? | ? |
            ```

                - Because the cell we uncovered was a 9, we can conclude that the 7 would be somewhere to the left:

            ```
            | ? | ? | ? | ? | 9 | X | X | X | X |
            ```

            Step 2: Of the cells to the left of the 9, inspect the middlemost value. Since there are two, we arbitrarily choose the left one.

            ```
            | ? | 4 | ? | ? | 9 | X | X | X | X |
            ```

                - Since the cell was a 4, we know 7 must be somewhere to the right:

            ```
            | X | 4 | ? | ? | 9 | X | X | X | X |
            ```

            Step 3: There are only two more cells where the 7 can be, so we arbitrarily choose the left one.

            ```
            | X | 4 | 6 | ? | 9 | X | X | X | X |
            ```

            Step 4: Inspect the final remaining cell. (If it's not there, that means there is no 7 in the ordered array.)

            ```
            | X | 4 | 6 | 7 | 9 | X | X | X | X |
            ```

                - We found the 7 in four steps

            - Binary search is only possible within an ordered array.
            - For smaller arrays, the difference between linear and binary search is negligible. But as arrays grow in size, so does the difference between the two algorithms.

                - A pattern emerges where for each time we double the size of the ordered array, the number of steps needed for the binary search increases by 1.

                    | Elements (N) | Binary search steps (worst case) |
                    |---|---|
                    | 1 | 1 |
                    | 2 | 2 |
                    | 4 | 3 |
                    | 8 | 4 |
                    | 16 | 5 |
                    | 32 | 6 |
                    | 64 | 7 |
                    | 128 | 8 |
                    | 256 | 9 |
                    | 512 | 10 |

                    ```
                    steps
                      10 |                                                     *
                       9 |                                        *
                       8 |                              *
                       7 |                       *
                       6 |                 *
                       5 |             *
                       4 |          *
                       3 |        *
                       2 |     *
                       1 |  *
                         +-----------------------------------------------------
                            1  2   4     8        16          32         ... 512  (N, doubling)
                    ```

                    That's the shape: steep for small N, flattening hard as N grows -- the hallmark of logarithmic growth.

            - We can use a binary search to speed up insertion into an ordered array as well.
