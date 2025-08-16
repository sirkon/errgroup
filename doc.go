// Package errgroup is an alternative to golang.org/x/sync/errgroup where the
// inner group's context is delivered directly to group's closures instead of
// using an error-prone side context like it is done with in the original library.
package errgroup
