<!--
SPDX-FileCopyrightText: 2026 Kalle Fagerberg

SPDX-License-Identifier: CC-BY-4.0
-->

# Game math library for Firefly Go SDK

[![NO AI](https://raw.githubusercontent.com/nuxy/no-ai-badge/master/badge.svg)](https://github.com/nuxy/no-ai-badge)
[![Go Reference](https://pkg.go.dev/badge/github.com/applejag/firefly-go-math.svg)](https://pkg.go.dev/github.com/applejag/firefly-go-math)
[![REUSE status](https://api.reuse.software/badge/github.com/applejag/firefly-go-math)](https://api.reuse.software/info/github.com/applejag/firefly-go-math)

Complementary library to [firefly-go](https://github.com/firefly-zero/firefly-go)
and [tinymath](https://github.com/orsinium-labs/tinymath)
containing math utilities not found in the other packages.

Contains tools like `ffrand`, a reimplementation of the Go [`math/rand`](https://pkg.go.dev/math/rand)
that is built upon the [`firefly.GetRandom()`](https://pkg.go.dev/github.com/firefly-zero/firefly-go/firefly#GetRandom)
function instead.

The `ffmath` package contains utility functions, and are using generics
extensively. Mostly as an experiment on generics usage, and because it's fun :)
I have compared the compiled output with non-generic variants to confirm
that the generic ones does not cause any regressions, even though they use
quirky hacks like `switch x := any(a).(type)`

## Installation

```bash
go get github.com/applejag/firefly-go-math
```

## License

This project conforms to the [REUSE](https://reuse.software/) standard.
Different parts of the code base use different licenses.

In general:

- Code: [MIT](./LICENSES/MIT.txt)
- Config files (e.g `go.mod`): [CC0-1.0](./LICENSES/CC0-1.0.txt)
- Docs: [CC-BY-4.0](./LICENSES/CC-BY-4.0.txt)
