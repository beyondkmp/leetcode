fn largest_number(nums: Vec<i32>) -> String {
    let mut tmp = nums.clone();
    tmp.sort_by(|a, b| format!("{}{}", b, a).cmp(&format!("{}{}", a, b)));

    if tmp[0] == 0 {
        return "0".to_string();
    }

    tmp.iter().map(ToString::to_string).collect()
}
fn main() {
    let v = vec![10, 2];
    println!("the result is {}", largest_number(v));
    let v = vec![3, 30, 34, 5, 9];
    println!("the result is {}", largest_number(v));
    let v = vec![0, 0, 0];
    println!("the result is {}", largest_number(v));
}
