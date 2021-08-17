isPalindrome = str => {
    
    str = str.replace( /\s/g, '').replace(/[\u2000-\u206F\u2E00-\u2E7F\\'!"#$%&()*+,\-.\/:;<=>?@\[\]^_`{|}~]/g, '')
    str = str.toLowerCase()
    var len = str.length
    for(let i=0; i< len; i++) {
        if (str[i] != str[len - 1- i]) {
            return false
        }
    }
    return true
}

console.log(isPalindrome('“five|\_/|four”'))