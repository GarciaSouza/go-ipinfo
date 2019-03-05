package ipinfo

import (
	"testing"
)

var ipGoogle = OtherIP("8.8.8.8")
var ipOpenDNS = OtherIP("208.67.222.222")

func AssertStrEqual(t *testing.T, expected string, got string) {
	if got != expected {
		t.Errorf(`Expected '%s' got '%s'`, expected, got)
	}
}

func AssertFloatEqual(t *testing.T, expected float64, got float64) {
	if got != expected {
		t.Errorf(`Expected '%f' got '%f'`, expected, got)
	}
}

func TestCity(t *testing.T) {
	AssertStrEqual(t, "Mountain View", ipGoogle.City)
}

func TestRegion(t *testing.T) {
	AssertStrEqual(t, "California", ipGoogle.Region)
}

func TestCountry(t *testing.T) {
	AssertStrEqual(t, "US", ipGoogle.Country)
}

func TestLocation(t *testing.T) {
	AssertStrEqual(t, "37.3860,-122.0840", ipGoogle.Location)
}

func TestPhone(t *testing.T) {
	AssertStrEqual(t, "650", ipGoogle.Phone)
}

// Testar a localização em coordenadas no resultado da informação.
func TestPostal(t *testing.T) {
	AssertStrEqual(t, "94035", ipGoogle.Postal)
}

func TestExtractLatLng(t *testing.T) {
	lat, lng := ExtractLatLng(ipGoogle)

	AssertFloatEqual(t, lat, 37.386)
	AssertFloatEqual(t, lng, -122.084)
}
