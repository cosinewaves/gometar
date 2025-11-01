package gometar

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/tabwriter"
)

// a fully parsed METAR output
type MetarData struct {
	Station     string
	Day         string
	Time        string
	Wind        string
	Visibility  string
	Temperature string
	Pressure    string
}

// main decoding function
func DecodeMETAR(metar string) (*MetarData, error) {
	fields := strings.Fields(metar)
	data := &MetarData{}

	// weather station identifier - like EGLL or KLAX
	data.Station = fields[0]

	// the timestamp
	reTime := regexp.MustCompile(`\b(\d{2})(\d{2})(\d{2})Z`)
	for _, f := range fields {
		if match := reTime.FindStringSubmatch(f); match != nil {
			data.Day = match[1]
			data.Time = match[2] + ":" + match[3] + "Z" // Z = zulu time
		}
	}

	// get wind 
	reWind := regexp.MustCompile(`\b(\d{3})(\d{2})KT`)
	for _, f := range fields {
		if match := reWind.FindStringSubmatch(f); match != nil {
			data.Wind = fmt.Sprintf("%s° at %s kt", match[1], match[2])
		}
	}

	// get the visibility
	reVis := regexp.MustCompile(`\b(\d{4})\b`)
	for _, f := range fields {
		if match := reVis.FindStringSubmatch(f); match != nil && len(match[1]) == 4 {
			data.Visibility = match[1] + " m"
		}
	}

	// temp and dew point
	reTemp := regexp.MustCompile(`\b(M?\d{2})/(M?\d{2})\b`)
	for _, f := range fields {
		if match := reTemp.FindStringSubmatch(f); match != nil {
			data.Temperature = match[1] + "°C"
		}
	}

	// atmospheric pressure
	rePres := regexp.MustCompile(`(?:Q|A)(\d{4})`)
	for _, f := range fields {
		if match := rePres.FindStringSubmatch(f); match != nil {
			if strings.HasPrefix(f, "A") {
				inHg := match[1][:2] + "." + match[1][2:]
				data.Pressure = inHg + " inHg"
			} else {
				data.Pressure = match[1] + " hPa"
			}
		}
	}

	return data, nil
}

// print parsed metar
func PrintReport(m *MetarData) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Field\tValue")
	fmt.Fprintln(w, "ICAO Station\t"+m.Station)
	fmt.Fprintln(w, "Date\t"+m.Day)
	fmt.Fprintln(w, "Time (UTC)\t"+m.Time)
	fmt.Fprintln(w, "Wind\t"+m.Wind)
	fmt.Fprintln(w, "Visibility\t"+m.Visibility)
	fmt.Fprintln(w, "Temperature\t"+m.Temperature)
	fmt.Fprintln(w, "Pressure\t"+m.Pressure)
	w.Flush()
}