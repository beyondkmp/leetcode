function sortArray(nums: number[]): number[] {
  const aux: number[] = new Array<number>(nums.length);

  const merge = (lo: number, mid: number, hi: number) => {
    let i = lo;
    let j = mid + 1;
    for (let k = lo; k <= hi; k++) {
      aux[k] = nums[k];
    }

    for (let k = lo; k <= hi; k++) {
      if (i > mid) {
        nums[k] = aux[j++];
      } else if (j > hi) {
        nums[k] = aux[i++];
      } else if (aux[j] < aux[i]) {
        nums[k] = aux[j++];
      } else {
        nums[k] = aux[i++];
      }
    }
  }


  const sort = (lo: number, hi: number) => {
    if (hi <= lo) {
      return;
    }
    const mid: number = lo + Math.trunc((hi - lo) / 2);
    sort(lo, mid);
    sort(mid + 1, hi);
    merge(lo, mid, hi);
  }
  sort(0, nums.length - 1);
  return nums;
};

console.log(sortArray([1, 9, 8, 7, 6, 0, 0, 1]));
