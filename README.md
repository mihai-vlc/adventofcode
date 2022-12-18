
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
```
```output
79
```
