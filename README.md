# Go Ruby-Runner

For running ruby code string and getting its result in Go applications.

## Install

```bash
$ go get -u github.com/meinside/go-rbrunner
```

## Usage

```go
package main

import (
	"fmt"

	rbrunner "github.com/meinside/go-rbrunner"
)

const code = `#!/usr/bin/env ruby

puts "args = #{ARGV[0]}, #{ARGV[1]}"`

func main() {
	result := rbrunner.Run(code, "arg1", "arg2")

	fmt.Printf("result = %+v\n", result)
}
```
