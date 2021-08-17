drop = (arr, callback) => {
    for (const key in arr) {
        if (callback(arr[key]) == true) {
            return arr.slice(key)
        }
    }
    return []
    
}

console.log(drop([1, 2, 3, 4], n => n === 0))