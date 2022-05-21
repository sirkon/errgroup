# errgroup
errgroup implementation without ctx override quirk of original errgroup package

## Installation

```shell
go get github.com/sirkon/errgroup
```

## Why this package?

The original [golan.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync@v0.0.0-20220513210516-0976fa681c29/errgroup)
has unfortunate API which causes unintended context override:

```go
eg, ctx := errgroup.WithContext(ctx) // This context is only really needed within the error group, not the outside
```

This package fixes it by not exposing its context outside, it is to be delivered as Go's closure argument:

```go
eg := errgroup.WithContext(ctx)

eg.Go(func(ctx context.Context) error {
	
})
```



