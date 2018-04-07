package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"googlemaps.github.io/maps"
)

type LatLongByAddressResponse struct {
	Address string
	Lat     float64
	Long    float64
}

type HTMLError struct {
	Err  string
	Code int
}

func LatLongByAddress(w http.ResponseWriter, req *http.Request) {

	client, _ := maps.NewClient(maps.WithAPIKey("{{sign up at https://developers.google.com/maps/documentation/geolocation/intro to get a key}}"))

	r := &maps.GeocodingRequest{
		Address: req.URL.Path[len("/latlongByAddress/"):],
	}

	resp, err := client.Geocode(context.Background(), r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil && !strings.Contains(err.Error(), "ZERO_RESULTS") {

		w.WriteHeader(http.StatusBadRequest)
		errResult := &HTMLError{
			Err:  err.Error(),
			Code: http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(errResult)
		log.Println("WARN: fatal error: %s", err)
	} else {

		if err != nil && strings.Contains(err.Error(), "ZERO_RESULTS") {

			w.WriteHeader(http.StatusBadRequest)
			errResult := &HTMLError{
				Err:  err.Error(),
				Code: http.StatusBadRequest,
			}
			json.NewEncoder(w).Encode(errResult)
		} else {

			result := &LatLongByAddressResponse{
				Address: resp[0].FormattedAddress,
				Lat:     resp[0].Geometry.Location.Lat,
				Long:    resp[0].Geometry.Location.Lng,
			}
			json.NewEncoder(w).Encode(result)
		}
	}
}
