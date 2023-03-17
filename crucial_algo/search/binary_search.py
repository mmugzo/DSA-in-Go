def search(arr: list, target: int)->int:
    left, right = 0, len(arr) -1

    while left <= right:
        mid  = left + (right - left) //2
        if arr[mid] == target:
            return mid
        elif arr[mid] > target:
            right = mid - 1
        else:
            left = mid + 1
    -1

arr = [1,2,3,4,5,6,7]
target = 4

result = search(arr, target)
if result != -1:
    print(f"target found at index {result}")
else:
    print(f"target not found")