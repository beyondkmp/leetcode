use std::cmp;
impl Solution {
    pub fn max_profit_assignment(difficulty: Vec<i32>, profit: Vec<i32>, worker: Vec<i32>) -> i32 {
        let mut max_profits = vec![0; 100000 + 1];
        for (i, v) in difficulty.iter().enumerate() {
            let index = *v as usize;
            max_profits[index] = cmp::max(max_profits[index], profit[i])
        }

        let mut index = 1;
        while index < max_profits.len() {
            max_profits[index] = cmp::max(max_profits[index], max_profits[index - 1]);
            index += 1;
        }

        let mut result = 0;
        for v in &worker {
            result += max_profits[*v as usize];
        }

        result
    }
}
