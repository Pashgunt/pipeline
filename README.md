# Pipeline Golang Helper

**Pipeline** - is a tool that allows you to create and manage processing chains (pipelines) to perform
various operations sequentially or in parallel. It provides a convenient way to organize
and perform complex sequential or parallel tasks in an application.

This tool is especially useful in the following situations:

1. Request Processing Chain: **Pipeline** can be used to process incoming HTTP requests through a series of steps
processing such as authentication, validation, authorization, etc.
2. Multiple Data Transformation: It helps to apply multiple transformations to the data in the chain, avoiding
redundant code and making the process manageable and understandable.
3. Modularity and extensibility: **Pipeline** makes the code more modular and extensible, since each processing step can
be easily added, modified or deleted without changing the basic logic of the application.
4. Ease of debugging and testing: Using **Pipeline** simplifies debugging and testing because each stage
The processing can be tested separately, and the passage of data through the chain is also easily tracked.

The overall benefit of using **Pipeline** is to simplify the organization and execution of complex
data processing, increase code modularity and extensibility, and reduce code connectivity. But it can
have negative consequences, such as increased readability and understanding of the code.

### Methods

This helper provides the following methods:

```Go
Send(data config.SendDataType)
```

**Set the object being sent through the pipeline.**
This method allows you to specify which data set needs to be sent to pipes for processing and which can be digitized in
these pipes. The data must have a `config.SendDataType` type.

```Go
Through(structures []interface{})
```

**Set the array of pipes.**
This method allows you to specify which set of pipes the transmitted data will pass through. As an input parameter, the
method must accept a collection of structures, for example `&StructureName{}` or an anonymous
function `func(data config.SendDataType){}`, the
main
thing is that the passed structure has a public `Handle(data config.SendDataType)` method, it can be redefined using
the `Via(method string)`
method. This method or
anonymous function must accept the `config.SendDataType` data type as input and return it as well.

```Go
Via(method string)
```

**Set the method to call on the pipes.**
This method allows you to override the standard Handle method to be called in structures.

```Go
HandleException(callback func(ctx context.Context))
```

**Handle the given exception.** This method allows you to handle `panic` that
will occur during the execution of logic inside the pipe or when trying to run the helper `Pipeline`.

```Go
HandleCarry(callback func(ctx context.Context))
```

**Handle the value returned from each pipe before passing it to the next.**
This method allows you to process the result that the pipe returned before passing it to the next pipe.

```Go
Pipes() []interface{}
```

**Set the object being sent through the pipeline.** This method returns an array of pipes that were specified in the
method `Through(structures []interface{})`.

```Go
Then()
```

**Set the object being sent through the pipeline.** This method allows you to start processing all pipes specified in
the `Through(structure []interface{})` method and process the send data from the `Send(data config.SendDataType)` method
using the default `Handle(data config.SendDataType)` method or the one that was reinterpreted `Via(method string)`.

```Go
ThenReturn() config.SendDataType
```

**Set the object being sent through the pipeline.** This method works according to the same logic as `Then()`, but its
main
difference is that it returns the result after passing all pipes.

### Examples

1. **Simple Example**
```Go
func main() {
(pipeline.NewPipeline()).
Send(
config.SendDataType{
"1": 1,
},
).
Through(
[]interface{}{
&TestStructure1{},
func(data config.SendDataType) config.SendDataType {
data["3"] = "3"
return data
},
&TestStructure2{},
},
).
Via("Test").
Then()
}

type TestStructure1 struct {
}

func (test1 *TestStructure1) Test(data config.SendDataType) config.SendDataType {
data["2"] = "2"
return data
}

type TestStructure2 struct {
}

func (test2 *TestStructure2) Test(data config.SendDataType) config.SendDataType {
data["4"] = "4"
return data
}

```
2. **An example using a variety of optional methods**
```Go
func main() {
data := (pipeline.NewPipeline()).
Send(
config.SendDataType{
"1": 1,
},
).
Through(
[]interface{}{
&TestStructure1{},
func(data config.SendDataType) config.SendDataType {
data["3"] = "3"
return data
},
&TestStructure2{},
},
).
HandleException(
func(ctx context.Context) {
fmt.Println("Exception Handle")
fmt.Println(ctx.Value("err"))
},
).
HandleCarry(
func(ctx context.Context) {
fmt.Println("Carry Handle")
fmt.Println(ctx.Value("data"))
},
).
ThenReturn()
}

type TestStructure1 struct {
}

func (test1 *TestStructure1) Handle(data config.SendDataType) config.SendDataType {
data["2"] = "2"
return data
}

type TestStructure2 struct {
}

func (test2 *TestStructure2) Handle(data config.SendDataType) config.SendDataType {
data["4"] = "4"
return data
}
```