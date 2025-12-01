- **Goal:** Create and pass a fixed-size **array** to a function.
- **Tasks:**
  1.  Define a **fixed-size array** of 5 integers (e.g., `[5]int`) and initialize it with any values.
  2.  Write a function called **`sumArray`** that takes a **fixed-size array** of 5 integers (`[5]int`) as an argument and returns the sum of all elements.
  3.  In your `main` function, call **`sumArray`** with your array and print the result.
  4.  _(Hint: Remember that fixed-size arrays are passed by value in Go.)_

---

- **Goal:** Use a tagless **`switch`** statement and practice basic type handling.
- **Tasks:**
  1.  Define a variable **`score`** of type **`int`** (e.g., set it to 85).
  2.  Write a function called **`getGrade`** that takes the integer `score` and returns a string representing the grade.
  3.  Use a **tagless `switch`** statement to check the score against ranges:
      - 90-100: "A"
      - 80-89: "B"
      - 70-79: "C"
      - Below 70: "D"
  4.  In `main`, print the grade returned by **`getGrade`**.

---

- **Goal:** Use different forms of the **`for`** loop and modify array elements.
- **Tasks:**
  1.  Define a **fixed-size array** of 8 integers (e.g., set them all to 0 initially).
  2.  Use a **standard `for` loop** (`for i := 0; i < len(arr); i++`) to assign a value to each element, where the value is the index multiplied by 2.
  3.  Use a **`for...range` loop** to print the index and the value of every element in the modified array.
  4.  Use a **`for` loop** to count how many elements in the array are greater than 10. Print the count.

---

- **Goal:** Practice functions with multiple return values and simulating success/failure.
- **Tasks:**
  1.  Write a function called **`findLargest`** that takes a **fixed-size array** of 4 integers (`[4]int`) and returns two values: the largest number found (`int`) and a boolean indicating if the array had any positive numbers.
  2.  Inside **`findLargest`**, use a **`for` loop** to iterate over the array, find the maximum value, and determine if any value was positive.
  3.  In `main`, call **`findLargest`** and use the boolean return value to print a message.
      - If `true`: print the largest number found.
      - If `false`: print "No positive numbers were found in the array."

---
