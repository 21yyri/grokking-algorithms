def search_smallest(array: list[int]) -> int:
    # Checks for the index for the smallest
    # number on the array passed as argument
    
    smallest_num = array[0]
    smallest_num_index = 0

    for index in range(1, len(array)):
        if array[index] < smallest_num:
            smallest_num = array[index]
            smallest_num_index = index

    return smallest_num_index


def select_sort(array: list[int]) -> list[int]:
    # Sorts the array by iterating over it to get 
    # the smallest value with complexity of O(n^2) 

    sorted_array = []

    for _ in range(len(array)):
        sorted_array.append(
            array.pop(
                search_smallest(array)
            )
        )
    
    return sorted_array


if __name__ == "__main__":
    unsorted_array = [i for i in range(1, 1_000)][::-1]
    sorted_array = select_sort(unsorted_array)

    print(sorted_array)
