var isValid = function(s) {
    let stack = []
    let chars = { '}': '{', ']': '[', ')': '(' }

    for (let char of s) {
        if(chars[char]) {
            if(stack.pop() !== chars[char]){
                return false
            }
        } else {
            stack.push(char)
        }
    }

    return stack.length  === 0
};