# Clean Code

This project aims to provide some hints of clean coding style.

It consists of two modules, i.e., color and app,
each of which has several coding stages.

## Color Module

Color Module, a warper of go-colorful package,
supports encodings and conversions of RGB and HSL color spaces.
It has the following cleaning stages.

> This color module is for demonstration purpose only.
> Applications should use go-colorful directly.

### Format

Nowadays, code can be automatically formatted by
various tools based on the programming languages.
For instance, the code in format folder can be automatically
formatted by [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports).

```
$ cd color/format
$ goimports -w rgb2hsl.go
```

Similarly, readers can find many formatters for other languages,
e.g., [clang-format](https://clang.llvm.org/docs/ClangFormat.html) for C++,
[Uncrustify](https://github.com/uncrustify/uncrustify) for
C/C++/C#/Objective-C/D/Java/Pawn/VALA,
[JS Beautify](https://www.npmjs.com/package/js-beautify) for JavaScript.

### Lint

After formatting code, various linters can be used to provide hints
on improving code based on some general guidelines and common sense of writing good code.
In this project, [golint](https://github.com/golang/lint) is used on the formatted code above.

```
$ cd color/format
$ golint rgb2hsl.go
rgb2hsl.go:10:1: error should be the last type when returning multiple items
rgb2hsl.go:10:1: exported function Convert should have comment or be unexported
rgb2hsl.go:12:8: should replace errors.New(fmt.Sprintf(...)) with fmt.Errorf(...)
rgb2hsl.go:13:7: if block ends with a return statement, so drop this else and outdent its block
```

Based on the results of golint,
[rgb2hsl.go](https://github.com/lilissun/cleancode/blob/master/color/format/rgb2hsl.go)
in the *format* folder can be rewritten as
[rgb2hsl.go](https://github.com/lilissun/cleancode/blob/master/color/lint/rgb2hsl.go)
in the *lint* folder.

### Name

Linters can help us to identify common mistakes in coding.
However, linters only provide a baseline for clean code.
As can be seen from the code in *lint* folder,
the function name, variable names and variable types are chosen badly.
When *Convert(c)* is call in some application,
programmers can misunderstand the functionality of *Convert*
and the structure of *c*.

Hence, in the *name* folder, we propose two methods to refine the code.
1. We define two structures *RGBColor* and *HSLColor* to hold the color values.
2. We define the function *ToHSL* (*ToRGB* resp.)
as a member method in *RGBColor* (*HSLColor* resp.) for color conversion.

Other member methods can be defined as well,
e.g., constructor and equality checking function. 

## App Module

App Module is an web application
that provides conversions supported by the Color Module.
