# Clean Code

This project aims to provide some hints on clean coding style.

It consists of two modules, i.e., color and app,
each of which has several coding stages.

## Color Module

Color Module, a warper of go-colorful package,
supports encodings and conversions of RGB and HSL color spaces.
It has the following cleaning stages.

> The color module is for demonstration purpose only.
> Applications should use go-colorful directly,
> and define RGBColor and HSLColor there if necessary.

### Format - use some tool to format your code

Nowadays, code can be automatically formatted by
various tools based on the programming languages.
For instance, the code in format folder can be automatically
formatted by [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports).

```
$ cd color/1.format
$ goimports -w rgb2hsl.go
```

Similarly, readers can find many formatters for other languages,
e.g., [clang-format](https://clang.llvm.org/docs/ClangFormat.html) for C++,
[Uncrustify](https://github.com/uncrustify/uncrustify) for
C/C++/C#/Objective-C/D/Java/Pawn/VALA,
[JS Beautify](https://www.npmjs.com/package/js-beautify) for JavaScript.

### Lint - highlight error-prone code

After formatting code, various linters can be used to provide hints
on improving code based on some general guidelines of writing good code.
In this project, [golint](https://github.com/golang/lint) is used on the formatted code above.

```
$ cd color/1.format
$ golint rgb2hsl.go
rgb2hsl.go:10:1: error should be the last type when returning multiple items
rgb2hsl.go:10:1: exported function Convert should have comment or be unexported
rgb2hsl.go:12:8: should replace errors.New(fmt.Sprintf(...)) with fmt.Errorf(...)
rgb2hsl.go:13:7: if block ends with a return statement, so drop this else and outdent its block
```

Based on the results of golint,
[rgb2hsl.go](https://github.com/lilissun/cleancode/blob/master/color/1.format/rgb2hsl.go)
in the *1.format* folder can be rewritten as
[rgb2hsl.go](https://github.com/lilissun/cleancode/blob/master/color/2.lint/rgb2hsl.go)
in the *2.lint* folder.

### Name - choose variable and function names wisely

Linters can help us to identify common mistakes in coding.
However, linters only provide a baseline for clean code.
As can be seen from the code in *lint* folder,
the function name, variable names and variable types are chosen badly.
When *Convert(c)* is call in some application,
programmers can misunderstand the functionality of *Convert*
and the structure of *c*.

Hence, in the *3.name* folder, we propose the following to refine the code.
1. We define two structures *RGBColor* and *HSLColor* to hold the color values.
2. We define the function *ToHSL* (*ToRGB* resp.)
as a member method in *RGBColor* (*HSLColor* resp.) for color conversion.

Other member methods can be defined as well,
e.g., constructor and equality checking function.

### Test - write your test cases, no joke

The necessities of tests can be well illustrated
by the test written in the *4.test* folder.
Even though we believed that the code is correct
(it is indeed correct in some sense),
the equality checking can fail because of the careless rounding.

In Go programming language, test suite is provided as a standard tool.
For instance, we can run the following commands.

```
$ cd color/4.test
$ go test
--- FAIL: TestConvertRGBtoHSL (0.00s)
        name_test.go:21: convert hsl=[{270 0.8 0.5}] to rgb=[{0.49999999999999967 0.09999999999999998 0.9}] but expect=[{0.5 0.1 0.9}]
FAIL
exit status 1
FAIL    bitbucket.org/lilissun/cleancode/color/test     0.006s
```

The fix can be found in the *color* folder.
We use a helper function called *round* to solve this problem.

## App Module

App Module is an web application
that provides conversions supported by the Color Module.

### Init - bad practice may not seem obvious

Initially, we implement everything very straight forward.
We have two functions to handle the conversions between RGB and HSL
in [rgb2hsl.go](https://github.com/lilissun/cleancode/blob/master/app/1.init/rgb2hsl.go) and
[hsl2rgb.go](https://github.com/lilissun/cleancode/blob/master/app/1.init/hsl2rgb.go) respectively.
They decode the color as json from the request, convert it accordingly,
and put the JSON encoding of the converted color in the response.

In [main.go](https://github.com/lilissun/cleancode/blob/master/app/1.init/main.go),
when the http request is received,
we use a if-else condition (or we can use a switch condition) to decide
which function should be called on the request payload.

Imagining that we have such a http server for hundreds of functions,
which is quite common in reality,
we can have a pretty long function filled by if-else conditions.
This can be problematic for maintenance in a long run.
(Yes. Even though it has such a simple logic form, we cannot put it in our code.)

### Lang - benefit from the language you use

Often, we choose a programming language for some reason, especially in industry.
For instance, C/C++ is suitable for high performance computing,
while more maintenance efforts are needed in large applications.
Python is suitable fast development with great community support,
while its application/library sometimes need C++ implementation for better performance.
Go programming language is famous for its packaging method and parallelization power.
Hence, it is extremely suitable for implementing servers.

In Go, we can easily define a mapping from strings to functions.
Hence, we can use the service name as key and use the actual service as value.
Whenever we received a request,
we search for the service in the map and call the service on the value.
Then, the registration of the services can be done
in the corresponding service implementation files.
By using the map, we no longer need the long if-else route function.
Instead, the route function now becomes a simple search in the map as follows.

```go
func route(name string, req []byte) ([]byte, error) {
	if proc, exist := router[name]; exist {
		return proc(req)
	}
	return nil, fmt.Errorf("service name=[%s] is unregistered", name)
}
```

### Monitor - log important things and check them regularly

In reality, services, once they are started,
need to run on servers for months/years continuously.
Meanwhile, we have hundreds of servers
running on hundreds of machines at the same time.
How can we make sure that everything works as expected?
Test cases may not be able to identify every bug.
Hence, we need to monitor our services on live.

Generally, monitors are developed based on well-defined logs.
Then, log collectors can parse the log and get the runtime information,
including processing time and runtime errors.
The collected logs can be presented using some web applications, e.g., Grafana, Kibana.

### App - the key of writing clean code is non-stoping refinement

Yes. There must be a better way.
The router should handle cases where duplicated keys are registered.
The number of handles running in parallel should be limited to the maximum number of processors available.
The logs should be written to local disks in rotated files with duration limit.
And many many more...
