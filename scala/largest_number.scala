object Solution {
  def largestNumber(nums: Array[Int]): String = {
    val s =
      nums.sortWith(
        (a, b) => a.toString + b.toString > b.toString + a.toString
      )
    if (s(0) == 0) "0" else s.mkString
  }

  def main(args: Array[String]): Unit = {
    val a = Array(1, 2, 2, 34)
    println(largestNumber(a))

    val b = Array(0, 0)
    println(largestNumber(b))
  }
}
