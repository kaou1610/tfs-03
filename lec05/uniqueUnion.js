uniqueUnion = (...args) => {
    const arr = [].concat(...args)
    return [...new Set(arr)]

}

console.log(uniqueUnion([1, 3, 2], [3, 4, 5]))
