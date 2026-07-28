"""
Chapter 10: Exercises
"""


def array_parser(array):
    if len(array) == 0:
        return

    if type(array[0]) == int:
        print(array[0])
        array_parser(array[1:])
    else:
        array_parser(array[0])
        array_parser(array[1:])


print("******** Array Parser ********")
num_array = [1,
             2,
             3,
             [4, 5, 6],
             7,
             [8,
              [9, 10, 11,
                [12, 13, 14]
              ]
             ],
             [15, 16, 17, 18, 19,
               [20, 21, 22,
                 [23, 24, 25,
                   [26, 27, 29]
                 ], 30, 31
               ], 32
             ], 33
            ]
array_parser(num_array)