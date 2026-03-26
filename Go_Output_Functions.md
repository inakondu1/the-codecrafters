# Go Output Functions

Go has three functions to output text:

    Print()
    Println()
    Printf()
print() : this is just  like gabage in gabage out. it print exactly waht the user inputed.

Println(): This Println print the output in a newline.

printf() : This is just like a short hand to a variables decared. 
| Verb   | Meaning                     | Example                                |
| ------ | --------------------------- | -------------------------------------- |
| `%d`   | Integer (decimal)           | `fmt.Printf("%d", 42)` → `42`          |
| `%f`   | Float (decimal)             | `fmt.Printf("%f", 3.14)` → `3.140000`  |
| `%.2f` | Float with 2 decimals       | `fmt.Printf("%.2f", 3.14159)` → `3.14` |
| `%s`   | String                      | `fmt.Printf("%s", "Hello")` → `Hello`  |
| `%t`   | Boolean                     | `fmt.Printf("%t", true)` → `true`      |
| `%v`   | Default format for any type | `fmt.Printf("%v", 42)` → `42`          |

this are the short hand that are used for everthing in golang when using printf()