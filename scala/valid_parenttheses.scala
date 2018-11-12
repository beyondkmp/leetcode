object Solution {
  def isValid(s: String): Boolean = {
    val stack = scala.collection.mutable.Stack[Char]()
    for (ch <- s) ch match {
      case '(' | '{' | '[' => stack.push(ch)
      case _               => if (stack.isEmpty || math.abs(ch - stack.pop) > 2) return false
    }
    stack.isEmpty
  }

  def main(args: Array[String]): Unit = {
    val a = "((({[]})))"
    println(isValid(a))

    val b = "((({[]}))]"
    println(isValid(b))
  }
}
