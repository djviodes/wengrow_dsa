# Chapter 8: Blazing Fast Lookup with Hash Tables

- Most programming languages include a data structure called a hash table, and it has an amazing superpower: fast reading.

    - A hash table is a list of paired values with the first item in each pair being the key, and the second item being the value.
    - Looking up a value in a hash table has an efficiency of O(1) on average.

- The process of taking characters and converting them to numbers is known as hashing.

    - The code that is used to convert those letters into particular numbers is called a hash function.

        - Hash functions only need to meet one criterion to be valid:

            - A hash function must convert the same string to the same number every single time it is applied.

- Under the hood, a hash table stores its data in a bunch of cells in a row, similar to an array.

    - For every key-value pair, each value is stored at the index of the key, after the key has been hashed.

        - This makes lookups easy since all we have to do is:

            1. Hash the key we're looking up.
            2. Look inside of the cell at the address hashed from the key and return the value stored there.

- We can only get O(1) lookups from hash tables if certain conditions are true:

    1. We know the value's key.
    2. We use a key to find the value.

- Trying to add data to a cell that is already filled is known as a collision.

    - One classic approach for handling collisions is known as separate chaining where instead of placing a single value in a cell, it places it in a reference array.

        - This array contains subarrays where the first index is the key and the second index is the value.
        - The computer takes the following steps:

            1. It hashes the key.
            2. It looks up the associated cell and takes note of the array.
            3. It searches through the array linearly, looking at index 0 of each subarray until it finds the right key.

- A hash table efficiency depends on three factors:

    1. How much data we're storing in the hash table.
    2. How many cells are available in the hash table.
    3. Which hash function we're using.

- A good hash table strikes a balance of avoiding collisions while not consuming lots of memory.

    - A rule of thumb: for every 7 data elements stored in a hash table, it should have 10 cells.

        - The ratio of data to cells is called the load factor.

            - The ideal load factor is 0.7 (7 elements / 10 cells)
