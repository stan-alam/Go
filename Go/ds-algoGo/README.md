## Data Structures and Algorithms in GoLang

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%201.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%202.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%203.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%204.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%205.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%206.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%207.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%208.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%209.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2010.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2011.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2012.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2013.png" width="80%" height="80%">
</a>

```go
package main

import  (
  "fmt"
)

func powerSeries(x int) (int, int) {

  return x * x,  x * x * x

}
```

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2014.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2015.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2016A.png" width="80%" height="80%">
</a>

```Go
//cb 16A-B
```

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2016B.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2017_e.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2018.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2019.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2020A.png" width="80%" height="80%">
</a>

```Go
// 18.A
package main

import (
  "fmt"
)

type IProcess interface {
  process()
}
//adapter struct
type Adapter struct {
  adaptee Adaptee
}
```

```go
//19.A
func (adapter Adapter) process() {
    fmt.Println("Adapter Process")
    adapter.adaptee.convert()
}

type Adaptee struct {
  adapterType int
}

func (adaptee Adaptee) convert () {
  fmt.Println("Adaptee convert method")
}

func main() {
  var processor IProcess = Adapter{}
  processor.process()
}
```

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2020B.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2021.png" width="80%" height="80%">
</a>

<a>
  <img src="https://github.com/stan-alam/Go/blob/develop/Go/ds-algoGo/images/01/ds-algGo%20-%2022.png" width="80%" height="80%">
</a>
