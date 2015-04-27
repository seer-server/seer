# Golang "flag" Alternative

This was a simple test/personal project to provide a more familiar (standard) flag parsing library. The API is similar (but not yet as complete) as the built in Go flag library.

# Usage

For basic types, there are four helper methods:

```go
import (
        "fmt"
        "os"

        "github.com/bbuck/flag"
)

func main() {
        // First noticable difference is the <long>, <short> flag description
        str := flag.String("string, s", "NONE", "Some string")
        num := flag.Int("int, i", 0, "Some integer")
        flt := flag.Float("float, f", 0.0, "Some float")
        bl := flag.Bool("verbose, v", false, "Some boolean")

        flag.Parse(os.Args[1:])

        fmt.Println("String = ", *str)
        fmt.Println("Int = ", *num)
        fmt.Println("Float = ", *flt)
        fmt.Println("Verbose = ", *bl)
}
```

This could be run with long names (with `--`):

```bash
sample --str String --int 10 --float 3.4 --verbose
```

or short names (with `-`):

```bash
sample -s String -i 10 -f 3.4 -v
```

Boolean short flags can also be chained, and the last flag specified in a chain can be one that receives an input value (but it cannot occur anywhere else in the series).

```bash
sample -vi 10
```

The previous will enable the verbose flag and pass in a value to `-i` (`--int`).

Any additional arguments can be accessed via `flag.Args()`

There is a way to handle setting custom variable values, but it needs to be modified to function like the standard flag library and use the Value interface.

# Disclaimer

This library is very much still a work in progress. I'd highly advise against your usage unless you're willing to fix them and make a pull request or fix them locally. I don't care, just be warned. I don't guarantee or claim any of this works in it's current state.

# License

Copyright (c) 2015 Brandon Buck

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

