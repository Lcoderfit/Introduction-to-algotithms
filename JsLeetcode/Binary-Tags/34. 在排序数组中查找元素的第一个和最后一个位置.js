/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var searchRange = function(nums, target) {
    if(nums.length == 0) {
        return [-1, -1];
    }
    var i = 0; 
    var j = nums.length - 1;
    while(i < j) {
        var mid = Math.floor((i + j) / 2);
        if(nums[mid] < target) {
            i = mid + 1;
        } else {
            j = mid;
        }
    };
    if(nums[i] != target) {
        return [-1, -1];
    }
    res = [];
    res.push(i);
    i, j = 0, nums.length - 1
    while(i < j) {
        mid = Math.floor((i+j+1)/2);
        if(nums[mid] > target) {
            j = mid -1
        } else {
            i = mid
        }
    }
};