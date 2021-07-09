function multiply(num1: string, num2: string): string {
  let product: number[] = new Array<number>(num1.length + num2.length);
  product.fill(0);

  for (let i = num1.length - 1; i >= 0; i--) {
    for (let j = num2.length - 1; j >= 0; j--) {
      product[i + j + 1] += +num1[i] * +num2[j];
    }
  }


  let carry = 0;
  for (let i = product.length - 1; i >= 0; i--) {
    let tmp = product[i] + carry;
    carry = ~~(tmp / 10);
    product[i] = tmp % 10;
  }

  let firstZeroIndex = product.findIndex(p => p !== 0);
  if (firstZeroIndex === -1) {
    return "0";
  }
  return product.slice(firstZeroIndex).join('')
};

console.log(multiply("123", "45"));
