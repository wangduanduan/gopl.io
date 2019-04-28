<!-- TOC -->

- [1. 包、变量、函数](#1-包变量函数)
  - [1.1. 函数可以返回任意数量的返回值](#11-函数可以返回任意数量的返回值)
  - [1.2. 命名返回值](#12-命名返回值)
  - [1.3. 基本类型](#13-基本类型)
  - [1.4. 零值](#14-零值)
  - [1.5. 类型转换](#15-类型转换)
  - [1.6. 类型推导](#16-类型推导)
- [2. 参考](#2-参考)

<!-- /TOC -->

# 1. 包、变量、函数

- 在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的[link](https://tour.go-zh.org/basics/3)

```go
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42))
}
// 函数如果声明需要两个数，那么传一个数就会报错
```

```go
s := 1 // 短变量声明 只能在函数内部使用，类型会从初始值中获取
var s, x string // 多个变量，类型相同时的声明
var s  = "" // 仅能声明并初始化一个变量
var s = "", y = 1 // 会报错
var s, y int = 1, 2 // ok
var s string = "" // 
```
## 1.1. 函数可以返回任意数量的返回值

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

## 1.2. 命名返回值
不建议使用，可读性差
```go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

## 1.3. 基本类型
```
bool
string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点
float32 float64

complex64 complex128
```
## 1.4. 零值
没有明确初始值的变量声明会被赋予它们的 零值。

- 数值类型为 0，
- 布尔类型为 false，
- 字符串为 ""（空字符串）。

## 1.5. 类型转换
表达式 T(v) 将值 v 转换为类型 T
```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
## 1.6. 类型推导

在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。

当右值声明了类型时，新变量的类型与其相同：
```go
var i int
j := i // j 也是一个 int
```
不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度：

```go
i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```
尝试修改示例代码中 v 的初始值，并观察它是如何影响类型的。

# 2. 参考