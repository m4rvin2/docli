# Confirming that Docli is installed

Before going any further, let's make sure that you have Docli installed. Go ahead and create a `main.go` file and paste the following content inside:

```go
package main

import (
    "github.com/alecthomas/repr"
    "github.com/celicoo/docli/v2"
)

func main() {
    args := docli.Args()
    repr.Println(args)
}
```

Now build and run it:

```bash
$ go build
$ ./main

docli.args{
}
```

**Note**: the output is the [Abstract Syntax Structure](https://en.wikipedia.org/wiki/Abstract_syntax_tree) of the command-line arguments, and it's empty because you didn't pass any arguments. If you run it again passing arguments you should see a different output.
