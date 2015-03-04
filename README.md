# Ambrosio Distance Converter

First ever plugin for Ambrosio.
Converts distance lengths to different units.

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

    steve.Behaviours = steve.NewBehaviour(ambrosio_distance_converter.Converter);

	steve.Listen(3000)
}

```

License
----

MIT
