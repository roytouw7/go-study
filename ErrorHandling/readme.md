# Error Handling
> To practise and think about error handling and error handling patterns.

## Nil Pointer Receiver
In Go, unlike languages like Java, the receiver of a method can be nil. 
```go
(p *structPointer) name() string {
    return p.name   // panics if p is nil
}
```

In the Go standard packages they do not handle these kind of errors inside the methods.
**The caller is responsible** for ensuring the callee is not *nil* and handling panics if it is *nil*.

## Specific Error Handling
> todo add section


## Error Monad
> todo add section