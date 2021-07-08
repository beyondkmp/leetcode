function candy(ratings: number[]): number {
  if (ratings.length <= 1) {
    return 1;
  }

  const nums: number[] = new Array<number>(ratings.length);
  nums.fill(1);

  for (let i = 1; i < ratings.length; i++) {
    if (ratings[i] > ratings[i - 1]) {
      nums[i] = nums[i - 1] + 1;
    }
  }

  for (let i = ratings.length - 1; i >= 0; i--) {
    if (ratings[i - 1] > ratings[i] && nums[i - 1] <= nums[i]) {
      nums[i - 1] = nums[i] + 1;

    }
  }

  let result = 0;
  for (let i of nums) {
    result += i;
  }
  return result;

};

const test1 = [1,3,2,2,1];
console.log(candy(test1));
