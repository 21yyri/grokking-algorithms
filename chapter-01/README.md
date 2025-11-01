# CH1 - Binary Search
The main idea is to search for the index of an item while reducing steps when looking for said item inside an ordenated list by sequentially breaking the array in half for each iteration.

The way to start implementing this is by:
- Defining the `low` variable, which represents the start of the array under analysis and will always start with 0;
- Define the `high` variable, which will represent the end of the array and initially assumes the value of `len(array) - 1`;

The algorithm runs every time `low` is lower or equal to `high`, and breaks any time that condition isn't met. It is made to run sequentially, so it´s variables need to readjust their values in order to process the array adequately.

A third variable, referred to as `half` is in need to start comparing. It is given by the sum of the `low` and `high` variables, divided by two. For every iteration, given by the condition for it to run, the program must:
- Define the `half` variable for the current array in analysis;
- Compare it to the target value;
- When array´s `half` value is bigger than target, then `high` is now `half - 1`;
- Else, the `low` value is now going to be `half + 1`;
- If the target is not found by the function, a null value shall be returned;