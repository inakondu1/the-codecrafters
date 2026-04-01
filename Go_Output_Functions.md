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


## Go Formatting Verbs

Formatting Verbs for Printf()

Go offers several formatting verbs that can be used with the Printf() function.
General Formatting Verbs

```The following verbs can be used with all data types:
Verb 	Description
%v 	Prints the value in the default format
%#v 	Prints the value in Go-syntax format
%T 	Prints the type of the value
%% 	Prints the % sign
```

### Integer Formatting Verbs

The following verbs can be used with the integer data type


Verb 	Description


%b 	Base 2


%d 	Base 10


%+d 	

Base 10 

and always show sign


%o 	Base 8
%O 	Base 8, with leading 0o


%x 	Base 16, lowercase

%X 	Base 16, uppercase

%#x 	Base 16, with leading 0x

%4d 	Pad with spaces (width 4, right justified)

%-4d 	Pad with spaces (width 4, left justified)

%04d 	Pad with zeroes (width 4)

### String Formatting Verbs

The following verbs can be used with the string data type

Verb 	Description

%s 	Prints the value as plain string

%q 	Prints the value as a double-quoted string

%8s 	Prints the value as plain string (width 8, right justified)

%-8s 	Prints the value as plain string (width 8, left justified)

%x 	Prints the value as hex dump of byte values

% x 	Prints the value as hex dump with spaces


### Boolean Formatting Verbs

The following verb can be used with the boolean data type:

Verb 	Description

%t 	Value of the boolean operator in true or false format (same as using %v)

### Float Formatting Verbs

The following verbs can be used with the float data type:

Verb 	Description

%e 	Scientific notation with 'e' as exponent

%f 	Decimal point, no exponent

%.2f 	Default width, precision 2

%6.2f 	Width 6, precision 2

%g 	Exponent as needed, only necessary digits
