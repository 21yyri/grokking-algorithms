# Hashtables
This is a data structure that relates a key, to a value registered.

## How it works
Works by passing the key into a hash function that will return an integer related to its position in an array, thus making accessing an element an O(1) operation. If a value is passed to the same key, the previous is overwritten.

- Key pass into a hash function, returns an integer;
- Integer returned is used as an index to an array;
- Insert operations overwrite the index returned from the hash.

## Collision Flaw
The structure might be flawed if its hash function isn't met for the occasion. The hash function may not be fitting to the keys intended to be inserted, and a two or more keys can have the same hash. When that happens on an insertion operation, the value overwrites that index position, causing loss of data.

To avoid collisions, you must have:
- A good hash function;
- A low load factor.

## Load factor
As its based on an array, hashtables have limited size, and when the array is almost full, it will most likely need a resize. This can handle collision problems, as there will be more and more space to allocate entries.

This is given by a reason between allocated and free indexes on the array. It is recommended that when an array reaches around 70% of its capacity, it should double itself.

## Cool idea for the future
This reminds me of cache databases. It should be interesting to learn hashing functions and create a cache database for learning purposes.