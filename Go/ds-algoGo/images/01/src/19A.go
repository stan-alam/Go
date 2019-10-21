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
