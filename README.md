# Reverse-Table-Selector
Used to select tables to be graded in Mr. Phelps AP Stats class

I am currently working on a GUI so it is not handled through terminal

# How it Works

Starts with an array of numbers with a user specified length

`[1, 2, 3, 4, 5, 6, 7, 8]`

A random number is generated 1 -> array length

`4`

The number 4 is now replaced with an X

`[1, 2, 3, x, 5, 6, 7, 8]`

If the number 4 is generated again, it becomes a number again

`[1, 2, 3, 4, 5, 6, 7, 8]`

The game continues until there is only one number left that is not an X

`[x, x, x, x, 5, x, x, x]`

# Upcoming

1: Add a GUI
2: Fix the channels in the goroutine (learning more with channels)
3: Anything else suggested
