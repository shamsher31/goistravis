# goistravis
Check if your code is running on Travis CI

# How to install
<pre>
go get github.com/shamsher31/goistravis
</pre>

# How to use
<pre>
package main

import (
	"fmt"
	"github.com/shamsher31/goistravis"
)

func main() {
	fmt.Println(travis.Is())
  // true
}
</pre>

# Related
[gotisvdo](https://github.com/shamsher31/gotisvdo)
[goistext](https://github.com/ferhatelmas/goistext)
[goisimg](https://github.com/ferhatelmas/goisimg)

# Why
This package is inspired by [is-travis](https://www.npmjs.com/package/is-travis) npm module to check your code is running on Travis CI.

# License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
