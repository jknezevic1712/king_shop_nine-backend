# Notes

## Struct types

### Types defined in separate packages are different even if they share name and underlying structure

package one

type User struct {
ID int64
Name string
Age int32
DateOfBirth string
}

#### Not same!

package two

type User struct {
ID int64
Name string
Age int32
DateOfBirth string
}

### Printf format specifiers

%d - for printing integers
%f - for printing floating-point numbers
%c - for printing characters
%s - for printing strings
%p - for printing memory addresses
%x - for printing hexadecimal values
