# Go Ruby-Runner

For running a ruby code string and getting its result in Go applications.

Given code string is saved as a temporary file and executed by `ruby`.

## Install

```bash
$ go get -u github.com/meinside/go-rbrunner
```

## Usage

```go
package main

import (
	"fmt"

	rb "github.com/meinside/go-rbrunner"
)

const code = `#!/usr/bin/env ruby

puts "args = #{ARGV[0]}, #{ARGV[1]}"`

func main() {
	result := rb.Run(code, "arg1", "arg2")

	fmt.Printf("result = %+v\n", result)
}
```
