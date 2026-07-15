# Chapter 6: Optimizing for Optimistic Scenarios

- Being able to consider all scenarios, instead of just the worst case scenario, is an important skill that can help you choose the appropriate algorithm for every situation.

- Insertion Sort reveals the power of analyzing scenarios beyond the worst case. It consists of the following steps:

    1. In the first pass-through, we temporarily remove the value at index 1 and store it in a temporary variable.

        ```
             | 4 |
        | 8 |     | 2 | 3 |
        ```

        - In subsequent pass-throughs, we remove the values at the subsequent indexes.

    2. We then begin the shifting phase, where we take the value to the left of the gap and compare it to the value in the temporary variable.

        ```
             | 4 |
        | 8 |     | 2 | 3 |
          ^
        ```

        - If the value to the left of the gap is greater than the temporary variable, then we shift that value to the right.

            ```
            | 4 |
                 | 8 | 2 | 3 |
            ```

        - As we shift values to the right, inherently the gap moves leftward.

            - As soon as we encounter a value that is lower than the temporarily removed variable, or we reach the left end of the array, this shifting phase is over.

    3. We then insert the temporarily removed value into the current gap.

        ```
        | 4 | 8 | 2 | 3 |
        ```

    4. Steps 1 through 3 represent a single pass-through which we repeat until a pass-through begins at the final index of the array.

- There are four types of steps associated with Insertion Sort that each have an efficiency:

    1. Comparisons: O(N^2/2)
    2. Shifts: O(N^2/2)
    3. Removals: O(N - 1)
    4. Insertions: O(N - 1)

        - Total Efficiency: O(N^2 + 2N - 2)

* Big O Notation only takes into account the highest order of N when we have multiple orders added together.

    - So Insertion Sort is O(N^2)

        - This means Bubble Sort, Selection Sort, and Insertion Sort are all O(N^2)

- It is critical we take into account the average-case scenario since they occur most frequently.

- Here is a table comparing Selection Sort and Insertion Sort:

    | | Best Case | Average Case | Worst Case |
    |---|---|---|---|
    | Selection Sort | N^2/2 | N^2/2 | N^2/2 |
    | Insertion Sort | N | N^2/2 | N^2 |
