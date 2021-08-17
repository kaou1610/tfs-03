toSpinalCase = (str) => {
    // str = str.replace(/([a-z])([A-Z])/g, "$1 $2").replace(/\s+|_+/g, "-")
    // return str.toLowerCase()
    return str
        .split(/\s|_|(?=[A-Z])/)
        .join("-")
        .toLowerCase();
}

console.log(toSpinalCase('IAmHungry'))
