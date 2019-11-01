# Gaia

Elasticsearch client package for Go.

[![GoDoc](https://godoc.org/github.com/procore/gaia?status.svg)](https://godoc.org/github.com/procore/gaia) [![Go Report Card](https://goreportcard.com/badge/github.com/procore/gaia)](https://goreportcard.com/report/github.com/procore/gaia) [![GitHub release](https://img.shields.io/github/release/procore/gaia.svg)](https://github.com/procore/gaia/releases)  [![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v1.4%20adopted-ff69b4.svg)](CODE_OF_CONDUCT.md)

## Usage

To use the gaia client:

```golang
package main

import (
	"fmt"

	"github.com/procore/gaia"
)

func main() {
	config := gaia.NewConfig()
	config.Debug = true
	client := gaia.NewClient(config)
	response := client.Cat("indices")
	fmt.Println(response)
}
```

## Testing

To run tests:

```golang
$ go test ./test/
ok  github.com/procore/gaia/test  0.046s
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/procore/gaia/tags).

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## About Procore

<img
  src="https://www.procore.com/images/procore_logo.png"
  alt="Procore Logo"
  width="250px"
/>

Gaia is maintained by Procore Technologies.

Procore - building the software that builds the world.

Learn more about the #1 most widely used construction management software at
[procore.com](https://www.procore.com/)
