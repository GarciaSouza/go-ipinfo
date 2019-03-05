package ipinfo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	geo "github.com/kellydunn/golang-geo"
)

// Result represents the result of the request.
type IPInfoResult struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	Org      string `json:"org"`
	Phone    string `json:"phone"`
	Postal   string `json:"postal"`
}

// URL represents the address of the API.
var URL = "http://ipinfo.io/"

// MyIP returns information from the requester's IP address.
func MyIP() IPInfoResult {
	return IPInformation(URL + "json")
}

// OtherIP returns information from another IP.
func OtherIP(address string) IPInfoResult {
	return IPInformation(URL + address + "/json")
}

// IPInformation returns information from the requester's IP address.
func IPInformation(url string) IPInfoResult {
	resp, errRequest := http.Get(url)

	token := os.Getenv("IPINFO_TOKEN")
	http.Header.Set("Authorization", strings.Join("Bearer", token))

	if errRequest != nil {
		panic(errRequest)
	}

	defer resp.Body.Close()

	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		panic(errBody)
	}

	var result IPInfoResult
	json.Unmarshal(body, &result)

	return result
}

// ExtractLatLng extracts the latitude and longitude coordinates.
func ExtractLatLng(ip IPInfoResult) (float64, float64) {
	coordinates := strings.Split(ip.Location, ",")

	lat, errLat := strconv.ParseFloat(coordinates[0], 64)
	if errLat != nil {
		panic(errLat)
	}

	lng, errLng := strconv.ParseFloat(coordinates[1], 64)
	if errLng != nil {
		panic(errLng)
	}

	return lat, lng
}

// Distance returns the distance in Km from the large circle between the
// location of MyIP and OtherIP.
func Distance(origin IPInfoResult, destiny IPInfoResult) float64 {
	latOrig, lngOrig := ExtractLatLng(origin)
	latDest, lngDest := ExtractLatLng(destiny)

	orig := geo.NewPoint(latOrig, lngOrig)
	dest := geo.NewPoint(latDest, lngDest)

	distance := orig.GreatCircleDistance(dest)
	return distance
}
