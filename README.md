# goistravis
Check if your code is running on Travis CI

### How to install
```go
go get github.com/shamsher31/goistravis
```

### How to use
```go
package main

import (
	"fmt"
	"github.com/shamsher31/goistravis"
)

func main() {
	fmt.Println(travis.Is())
    // true
}
```
###Aliasing Imports
If you already have package name ```travis``` you can use following.
```go
package main

import (
	"fmt"
	pckTravis "github.com/shamsher31/goistravis"
)

func main() {
	fmt.Println(pckTravis.Is())
    // true
}
```
### Related
[goisvdo](https://github.com/shamsher31/goisvdo)<br>
[goistext](https://github.com/ferhatelmas/goistext)<br>
[goisimg](https://github.com/ferhatelmas/goisimg)<br>

### Why
This package is inspired by [is-travis](https://www.npmjs.com/package/is-travis) npm module .

### License
MIT © [Shamsher Ansari](https://github.com/shamsher31)
