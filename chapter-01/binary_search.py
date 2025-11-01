def search(array: list[int], target: int) -> int | None:
    low, high = 0, len(array) - 1
    while low <= high:
        half = (low + high) // 2
        if array[half] == target:
            return half

        if array[half] > target:
            high = half - 1
        else:
            low = half + 1
    
    return None


if __name__ == "__main__":
    array = [i for i in range(1, 1e8)]

    target = int(input())
    result = search(array, target)

    print(result)
    