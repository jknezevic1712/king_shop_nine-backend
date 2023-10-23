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
