# Game math library for Firefly Go SDK

Complementary library to [firefly-go](https://github.com/firefly-zero/firefly-go)
and [tinymath](https://github.com/orsinium-labs/tinymath)
containing math utilities not found in the other packages.

Contains tools like `ffrand`, a reimplementation of the Go [`math/rand`](https://pkg.go.dev/math/rand)
that is built upon the [`firefly.GetRandom()`](https://pkg.go.dev/github.com/firefly-zero/firefly-go/firefly#GetRandom)
function instead.
