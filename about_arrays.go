package go_koans

func aboutArrays() {
  fruits := [4]string{"apple", "orange", "mango"}

  assert(fruits[0] == "apple") // indexes begin at 0
  assert(fruits[1] == "orange") // one is indeed the loneliest number
  assert(fruits[2] == "mango") // it takes two to ...tango?
  assert(fruits[3] == "") // there is no spoon, only an empty value

  // len is not number of array items that are not null, because
  // golang auto-initializes last array value to "".
  assert(len(fruits) == 4) // the length is what the length is
  assert(cap(fruits) == 4) // it can hold no more

  // adding that last "" to prove that golang loves auto-initializing. 
  // i'm used to nil/null as a third state. what if i want to know how
  // many indices in my array hold meaningful data? maybe cap() can also
  // be used for this as well as getting len of slice's origin array.
  assert(fruits == [4]string{"apple", "orange", "mango", ""}) // comparing arrays is not like comparing apples and oranges

  tasty_fruits := fruits[1:3]           // defining oneself as a variation of another
  assert(tasty_fruits[0] == "orange") // slices of arrays share some data
  assert(tasty_fruits[1] == "mango") // albeit slightly askewed

  assert(len(tasty_fruits) == 2) // its length is manifest
  assert(cap(tasty_fruits) == 3) // but its capacity is surprising!

  tasty_fruits[0] = "lemon" // are their shared roots truly identical?

  assert(fruits[0] == "apple") // has this element remained the same?
  assert(fruits[1] == "lemon") // how about the second?
  assert(fruits[2] == "mango") // surely one of these must have changed
  assert(fruits[3] == "") // but who can know these things

  veggies := [...]string{"carrot", "pea"}

  assert(len(veggies) == 2) // array literals need not repeat an obvious length
}
