def quicksort(arr: list[int]) -> list[int]:
    if len(arr) < 2:
        return arr

    left_array = []
    right_array = []

    pivot = arr[0]
    for num in arr[1:]:
        if pivot >= num:
            left_array.append(num)
        elif pivot < num:
            right_array.append(num)
    
    return quicksort(left_array) + [pivot] + quicksort(right_array)


lista = [9, 8, 7, 6, 5, 4, 3, 2, 1, 0]
print(quicksort(lista))
