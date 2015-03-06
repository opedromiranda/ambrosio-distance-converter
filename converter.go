package ambrosio_distance_converter

import (
	"bytes"
	"github.com/opedromiranda/ambrosio"
	"strconv"
)

var conversions = map[string]float64{
	"metermile":       0.000621371,
	"meterkilometer":  0.001,
	"metercentimeter": 100,

	"kilometermile":       0.621371,
	"kilometermeter":      1000,
	"kilometercentimeter": 100000,

	"milekilometer":  1.60934,
	"milemeter":      1609.34,
	"milecentimeter": 160934,
}

var metrics = map[string]string{
	"km":         "kilometer",
	"kilometer":  "kilometer",
	"kilometers": "kilometer",

	"m":      "meter",
	"meters": "meter",
	"meter":  "meter",

	"mile":  "mile",
	"miles": "mile",
}

var Converter = ambrosio.Behaviour{
	"^([0-9][\\.[0-9]*]?)*[ ]?(km|mile|miles|meter|m|meters) (to|2) (kilometer|kilometers|km|mile|miles)$",
	func(matches []string) (string, error) {
		var buffer bytes.Buffer

		distance, _ := strconv.ParseFloat(matches[1], 64)
		originMetric := matches[2]
		destMetric := matches[4]

		parsedOriginMetric, _ := metrics[originMetric]
		parsedDestMetric, _ := metrics[destMetric]

		buffer.WriteString(parsedOriginMetric)
		buffer.WriteString(parsedDestMetric)

		result := distance * conversions[buffer.String()]
		return strconv.FormatFloat(result, 'f', 5, 64), nil
	},
}
