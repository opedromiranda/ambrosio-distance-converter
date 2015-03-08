# Ambrosio Distance Converter

First ever plugin for Ambrosio.
Converts distance lengths to different units.

[![BuildStatus](https://travis-ci.org/opedromiranda/ambrosio-distance-converter.svg)](https://travis-ci.org/opedromiranda/ambrosio-distance-converter)


### Version
0.0.1

### Usage

Add the behaviour to Ambrosio:

```golang
import (
    "github.com/opedromiranda/ambrosio"
    "github.com/opedromiranda/ambrosio-distance-converter"
)

func main() {
	steve := ambrosio.NewAmbrosio("Steve")

    steve.NewBehaviour(ambrosio_distance_converter.Behaviours);

	steve.Listen(3000)
}

```

License
----

MIT
