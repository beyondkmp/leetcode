use std::cmp;
use std::i32;
pub fn min_sub_array_len(s: i32, nums: Vec<i32>) -> i32 {
    let mut left = 0;
    let mut result = i32::MAX;
    let mut sum = 0;
    for (right, v) in nums.iter().enumerate() {
        sum += v;
        while sum >= s {
            result = cmp::min(result, (right - left + 1) as i32);
            sum -= nums[left];
            left += 1;
        }
    }

    match result {
        i32::MAX => 0,
        _ => result,
    }
}

fn main() {
    let s = 7;
    let nums = vec![2, 3, 1, 2, 4, 3];
    println!("the result is {}", min_sub_array_len(s, nums))
}
