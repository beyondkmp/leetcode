object Solution {
  def largestNumber(nums: Array[Int]): String = {
    val s = nums.sortWith((a,b)=> a.toString + b.toString > b.toString+a.toString).mkString
    if (s.startsWith("0")) "0" else s
  }

  def main(args: Array[String]): Unit ={
    val a = Array(1,2,2,34)
    println(largestNumber(a))
  }
}

