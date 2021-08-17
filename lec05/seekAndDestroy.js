seekAndDestroy = (arr, ...args) => {
    return arr.filter(value => !args.includes(value))
}

console.log(seekAndDestroy(["foo", "bar", 1], "foo", 1))