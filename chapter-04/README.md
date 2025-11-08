# Divide and Conquer Algorithm
The divide and conquer algorithm is a technique that tends to break down a problem into two or more sub-problems recursively until these problems turn into simple enough cases that can be solved.

It is present in many algorithms such as sorting (quicksort, merge sort), searching functions and more.

## Quicksort
This is a sorting algorithm that uses comparison and recursion to rearrange arrays. Works by:
- Selecting a pivot of your choice;
- Iterating over the array to see which are smaller or bigger than the pivot;
- Create separated arrays for what you find;
- Call the function on itself with the created arrays;
- Wait until they hit the base case, which is when the size of the array is less than two.
  
## Complexity
It varies on the pivot and array, but the worst case usually is O(n²), when the most desired result will be O(log n).