# cleancode

This project aims to provide some hints of clean coding style.

It consists of two modules, i.e., color and app,
each of which has several coding stages.

## Color Module

Color Module, a warper of go-colorful package,
supports encodings and conversions of RGB and HSL color spaces.
More color spaces can be easily extended.
It has the following cleaning stages.

### Format

Nowadays, code can be automatically formatted by
various tools based on the programming languages.
For instance, the code in format folder can be automatically
formatted by [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports).

cd color/format
goimports -w rgb2hsl.go

Similarly, readers can find formatter for other languages,
e.g., [clang-format](https://clang.llvm.org/docs/ClangFormat.html) for C++,
[UnCrustify](https://github.com/uncrustify/uncrustify) for
C/C++/C#/Objective-C/D/Java/Pawn/VALA,
[JS Beautify](https://www.npmjs.com/package/js-beautify) for JavaScript. 


### Lint

## App Module

App Module is an web application
that provides conversions supported by the Color Module.
