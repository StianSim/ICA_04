package weather

import (
    "testing"
)

// Please note: We only test the degreeToName function, as the other functions are
// for the most part only simple mathematical operations.

var tests_degrees = []struct {
    value       float64
    expected    string
}{
    {0, "Nord"},
    {168.74, "Sør-sørøst"},
    {168.75, "Sør"},
    {400, "Nordøst"}, // "Overflows" and rolls around to 40 degrees
}

func TestDegreeToName(t *testing.T) {
    for _, v := range tests_degrees {
      val := degreeToName(v.value)
      if val != v.expected {
          t.Errorf("degreeToName(%.2f) returned %s, expected %s", v.value, val, v.expected)
      }
  }
}
