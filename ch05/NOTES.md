# Chapter 5: Optimizing Code With and Without Big O

- There may be times when two competing algorithms are described in the same way using Big O, yet one algorithm is faster than the other.

- Selection Sort is another sorting algorithm with the following steps:

    1. We check each cell of the array from left to right to determine which value is least.

        ```
        Lowest value so far is 2
        | 2 | 6 | 1 | 3 |
          ^
        ```

        ```
        Lowest value is still 2
        | 2 | 6 | 1 | 3 |
              ^
        ```

        ```
        Lowest value is now 1
        | 2 | 6 | 1 | 3 |
                  ^
        ```

        ```
        Lowest value is still 1
        | 2 | 6 | 1 | 3 |
                      ^
        ```

    2. Once we determine which index contains the lowest value, we swap its value with the value we began the pass-through with.

        ```
        | 1 | 6 | 2 | 3 |
        ```

    3. We repeat pass-throughs of steps 1 and 2 until we reach a pass-through that would start at the end of the array.

- Selection Sort contains two types of steps: Comparisons and Swaps.

    - For N elements, we make:

        - (N - 1) + (N - 2) + (N - 3) + ... + 1 comparisons

    - As for swaps, though, we only need to make a maximum of one swap per pass-through.

        - Compare this to Bubble Sort, where in the worst case scenario we have to make a swap for each and every comparison.

            | N Elements | Max # of Steps in Bubble Sort | Max # of Steps in Selection Sort |
            |---|---|---|
            | 5 | 20 | 14 |
            | 10 | 90 | 54 |
            | 20 | 380 | 199 |
            | 40 | 1560 | 819 |
            | 80 | 6320 | 3239 |

* <u>Big O ignores constants!</u>

    - Due to this, Selection Sort and Bubble Sort are described by Big O as O(N^2), even though Selection Sort takes half of the time.

        - There is no O(N^2/2)

    - We ignore constants because Big O Notation only concerns itself with general categories of algorithm speeds.

        - Efficiencies in different categories are so different that it doesn't really matter whether the O(N) algorithm is actually O(2N), O(N/2), or even O(100N).

- A step is considered to be significant if as N approaches infinity, the number of times that the step executes grows without bound.
