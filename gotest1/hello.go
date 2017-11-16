package main

import "fmt"

// Format codes
//  %v - value in default Format
//  %#v - value in golang format
//  %+v - value + field name for structs
//  %T  - type of iem//   tem
//   %%  - print percent sign %
// Boolean
//  %t   - boolean - word 'true' or 'false'
// Integer
//  %b    - base 2
//  %d    - base 10
//  %o    - base 8
//  %c    - character represented by unicode point
//  %q    - single quoted character literal escaped w/go syntax
//  %x    - base 16, lower case letters
//  %X    - base 16, upper case letters
// Floating and Complex
//  %b	decimalless scientific notation with exponent a power of two,
//  %e	scientific notation, e.g. -1.234456e+78
//  %E	scientific notation, e.g. -1.234456E+78
//  %f	decimal point but no exponent, e.g. 123.456
//       %9.2f    - specify width and precision
//  %F	synonym for %f
//  %g	%e for large exponents, %f otherwise. Precision is discussed below.
//  %G	%E for large exponents, %F otherwise
// String and slice of bytes (treated equivalently with these verbs):
//  %s	the uninterpreted bytes of the string or slice
//  %q	a double-quoted string safely escaped with Go syntax
//  %x	base 16, lower-case, two characters per byte
//  %X	base 16, upper-case, two characters per byte
// Pointer:
//  %p	base 16 notation, with leading 0x
// FLAGS
//  + (include sign), - (left justify), # (alternate format),
//  ' ' (space in output), 0 (leading 0 padding)
//    \n   newline,   \t  tab

func main() {
	fmt.Printf("Hello, world, testing 1, 2, 3.\n")
	fmt.Println("This is a test of a testing")
	// set value for z
	z := 5
	zp := junk(z)
	fmt.Println(zp)
	var num = 42
	fmt.Printf("Decimal: %d, Binary: %b\n", num, num)
	fmt.Printf("Decimal: %d, Octal: %o\n", num, num)
	fmt.Printf("Decimal: %d, Hexadecimal: %x, Hex alt %#x\n", num, num, num)
	fmt.Printf("Decimal: %d, Hexadecimal: %X, Hex alt %#X\n", num, num, num)
	fmt.Printf("float %9.2f\n", 32.45678)
	bval := false
	fmt.Printf("bool value: %t\n", bval)
	sval := "a simple % \tstring \nhere"
	fmt.Printf("string with s: %s\n", sval)
	fmt.Printf("string with q: %q\n", sval)

	for i := 0; i < 200; i++ {
		fmt.Printf("%d\t%q\n", i, i)
	}
}
