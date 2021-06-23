function sortArray(unorders: number[]): number[] {
  const aux: number[] = [...unorders];
  const merge = (nums: number[], lo: number, mid: number, hi: number) => {
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


  const sort = (nums: number[], lo: number, hi: number) => {
    if (hi <= lo) {
      return;
    }
    const mid:number = lo + Math.trunc((hi - lo) / 2);
    sort(nums, lo, mid);
    sort(nums, mid + 1, hi);
    merge(nums, lo, mid, hi);
  }
  sort(unorders, 0, unorders.length - 1);
  return unorders;
};

console.log(sortArray([1, 9, 8, 7, 6 , 0, 0, 1]));
