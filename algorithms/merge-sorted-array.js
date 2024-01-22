/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
    // Copy elements from nums1 to a new array
    const auxArray = nums1.slice(0, m);

    let i = 0; // Index to iterate over auxArray
    let j = 0; // Index to iterate over nums2
    let k = 0; // Index to iterate over nums1

    // Merge elements from auxArray and nums2 in ascending order
    while (i < m && j < n) {
        if (auxArray[i] <= nums2[j]) {
            nums1[k] = auxArray[i];
            i++;
        } else {
            nums1[k] = nums2[j];
            j++;
        }
        k++;
    }

    // Add remaining elements from auxArray, if any
    while (i < m) {
        nums1[k] = auxArray[i];
        i++;
        k++;
    }

    // Add remaining elements from nums2, if any
    while (j < n) {
        nums1[k] = nums2[j];
        j++;
        k++;
    }
};