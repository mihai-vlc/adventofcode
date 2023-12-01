
# Advent of code 2022 solutions in golang

To execute the code

```
cd <day number>
go run main.go
```

### Debug console filter

```
!start,!type,!process,!detaching,!dlv,!DAP
```


```go
var a = " 79 "

var n, err = strconv.Atoi(strings.Trim(a, " "))

if err != nil {
	fmt.Println(err)
}
fmt.Println(n)
fmt.Println(n/7)
fmt.Println("A =", int('A'))
fmt.Println("a =", int('a'))
fmt.Println("E =", int('E'))
fmt.Println("96 =", string(rune(96)))
```
```output
79
11
A = 65
a = 97
E = 69
96 = 96
```
