/**
 * Converts a Roman numeral to an integer.
 * @param {string} s - Input string containing Roman numerals.
 * @returns {number} - Converted integer value.
 */
const romanToInt = function(s) {
    const romanMap = new Map([
        ['I', 1],
        ['V', 5],
        ['X', 10],
        ['L', 50],
        ['C', 100],
        ['D', 500],
        ['M', 1000]
    ]);

    let result = 0;
    let prevValue = 0;

    for (const char of s) {
        const currValue = romanMap.get(char);

        // Compare the current value with the previous value
        if (currValue > prevValue) {
            // If current value is greater than the previous, subtract previous value twice
            // and then add the current value
            result += currValue - 2 * prevValue;
        } else {
            // Otherwise, simply add the current value
            result += currValue;
        }

        // Cache the current value as the previous value for the next iteration
        prevValue = currValue;
    }

    return result;
};
