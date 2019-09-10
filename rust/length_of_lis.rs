use std::cmp::max;
pub fn length_of_lis(nums: Vec<i32>) -> i32 {
    let mut dp = vec![1; nums.len()];

    for i in 0..nums.len() {
        for j in 0..i {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j] + 1)
            }
        }
    }

    match dp.iter().max() {
        None => 0,
        Some(&x) => x,
    }
}
fn main() {
    let v = vec![100, 32, 57, 17, 75, 45, 108];
    println!("{}", length_of_lis(v))
}
