def sort(arr: list) -> list:
    n = len(arr)
    

    for i in range(n-1):
        for j in range(0, n-i-1):
            if arr[j] > arr[j+1]:
                
                arr[j], arr[j+1] = arr[j+1], arr[j]
                
    return arr

arr = [5,2,1,1,7,6,6]
print(sort(arr))