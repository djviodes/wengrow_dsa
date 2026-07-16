# Chapter 7: Big O in Everyday Code

- Determining the efficiency of our code is the first step in optimizing it.

- Whenever we have two distinct datasets that have to interact with each other through multiplication, we have to identify both sources separately when we describe the efficiency.

    - O(N • M)

        - O(N) <= O(N • M) <= O(N^2)

- An algorithm that is O(2^N) is incredibly slow.
