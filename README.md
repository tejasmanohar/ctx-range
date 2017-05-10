# ctx-range

An experimental, generic, [context]-aware `range` for Go.

# Why

Sometimes, you want to range over channels returned by functions and semantically

```
out := make(chan out)
Range(ctx, in, out)
for v := range out {

}
```

is semantically much clearer than `for { select { ... } } `. Plus, it's cleaner if
you have nesting out of the box.

# Should I use this?

Probably not. It's not type-safe because... Go. This also isn't something I'd actually
generate code for. I really just wanted an excuse to play with reflection. That said,
it would be awesome if this was magically a part of Go. The [context] package is
very powerful but feels like a 2nd-class citizen (at least, for now).


[context]: https://golang.org/pkg/context/
