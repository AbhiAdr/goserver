# Setup 
This document demonstrates the development of a simple Go package and the standard setup while working with go. Download the go according to your operating system. This steps are followed in ubuntu.

## Download & set up a workspace

[Download](https://golang.org/dl/) the archive and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:

```sh
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or ( $HOME/.profile | $HOME/.bashrc) :

>export PATH=$PATH:/usr/local/go/bin
>export GOPATH=$HOME/gocode
>export PATH=$PATH:$HOME/gocode/bin
```sh
$ mkdir gocode
$ cd gocode
```

A workspace is a directory hierarchy with three directories at its root:

-   **src** contains Go source files,
-   **pkg** contains package objects, and
-   **bin** contains executable commands.

## Working with source repository 
create **src** folder in your work space and folder for first program.
```sh
$ mkdir src
$ cd src
```

If you planning keep your code in a source repository somewhere, then you should use the root of that source repository as your base path.

By this way it will be easy for you to manage your code.

>mkdir -p $GOPATH/src/github.com/user
>mkdir -p $GOPATH/src/github.com/user/PROJECT_NAME

Example
```sh
mkdir -p $GOPATH/src/github.com/AbhiAdr/hello
cd $GOPATH/src/github.com/AbhiAdr/hello
```
create file hello.go
```
vi hello.go

package main

import "fmt"

func main() {
	fmt.Printf("Hello, world.\n")
}
```
## Install Package

It then installs that binary to the workspace's bin directory as hello (or, under Windows, hello.exe). In our example, that will be $GOPATH/bin/hello, which is $HOME/gocode/bin/hello.
```sh
go install 
$ hello
Hello, world.
```
### Your First libarary 

Let's write a library and use it from the hello program.

>Again, the first step is to choose a package path (we'll use $GOPATH/src/github.com/user/stringutil) and create the package directory:
```
mkdir stringutil 
cd stringutil

vi reverse.go

// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

Now change the hello word program
```
cd ..
cd hello/
vi hello.go
package main

import (
	"fmt"

	"github.com/user/stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
```
### Build package

This won't produce an output file.
it will create a executable file in current folder
```sh
$ go build
$ ls -la
./
../
hello*
hello.go
```
>To make it globaly , you must use go install.
> go install

it will create a package object inside the pkg directory of the workspace.