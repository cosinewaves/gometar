package gometar

import "testing"

func TestDecodeMETAR(t *testing.T) {
	sample := "EGLL 231651Z 18005KT 9999 FEW030 22/17 Q1018"

	data, err := DecodeMETAR(sample)
	if err != nil {
		t.Fatalf("DecodeMETAR returned error: %v", err)
	}

	if data.Station != "EGLL" {
		t.Errorf("Expected Station 'EGLL', got '%s'", data.Station)
	}

	if data.Day != "23" {
		t.Errorf("Expected Day '23', got '%s'", data.Day)
	}

	if data.Time != "16:51Z" {
		t.Errorf("Expected Time '16:51Z', got '%s'", data.Time)
	}

	if data.Wind != "180° at 05 kt" {
		t.Errorf("Expected Wind '180° at 05 kt', got '%s'", data.Wind)
	}

	if data.Visibility != "9999 m" {
		t.Errorf("Expected Visibility '9999 m', got '%s'", data.Visibility)
	}

	if data.Temperature != "22°C" {
		t.Errorf("Expected Temperature '22°C', got '%s'", data.Temperature)
	}

	if data.Pressure != "1018 hPa" {
		t.Errorf("Expected Pressure '1018 hPa', got '%s'", data.Pressure)
	}
}