// handlers_test.go
package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestWeatherByLatLongAndDateHandler(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://maps.googleapis.com/maps/api/geocode/json?address=777&key=AIzaSyBKJ5hORqnuHg33bzcB7AyZBDr7jKzhwsc",
		httpmock.NewStringResponder(200, `{
   "results" : [
      {
         "address_components" : [
            {
               "long_name" : "State Highway 777",
               "short_name" : "State Hwy 777",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Sarasota County",
               "short_name" : "Sarasota County",
               "types" : [ "administrative_area_level_2", "political" ]
            },
            {
               "long_name" : "Florida",
               "short_name" : "FL",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "United States",
               "short_name" : "US",
               "types" : [ "country", "political" ]
            }
         ],
         "formatted_address" : "State Hwy 777, Florida, USA",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : 27.0430423,
                  "lng" : -82.2915476
               },
               "southwest" : {
                  "lat" : 26.9618842,
                  "lng" : -82.3373313
               }
            },
            "location" : {
               "lat" : 26.9908952,
               "lng" : -82.30378979999999
            },
            "location_type" : "GEOMETRIC_CENTER",
            "viewport" : {
               "northeast" : {
                  "lat" : 27.0430423,
                  "lng" : -82.2915476
               },
               "southwest" : {
                  "lat" : 26.9618842,
                  "lng" : -82.3373313
               }
            }
         },
         "place_id" : "ChIJm9px5GRWw4gRlKu_WEOorhI",
         "types" : [ "route" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "Kentucky 777",
               "short_name" : "KY-777",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Floyd County",
               "short_name" : "Floyd County",
               "types" : [ "administrative_area_level_2", "political" ]
            },
            {
               "long_name" : "Kentucky",
               "short_name" : "KY",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "United States",
               "short_name" : "US",
               "types" : [ "country", "political" ]
            }
         ],
         "formatted_address" : "KY-777, Kentucky, USA",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : 37.5325221,
                  "lng" : -82.779815
               },
               "southwest" : {
                  "lat" : 37.4677171,
                  "lng" : -82.834879
               }
            },
            "location" : {
               "lat" : 37.494712,
               "lng" : -82.784621
            },
            "location_type" : "GEOMETRIC_CENTER",
            "viewport" : {
               "northeast" : {
                  "lat" : 37.5325221,
                  "lng" : -82.779815
               },
               "southwest" : {
                  "lat" : 37.4677171,
                  "lng" : -82.834879
               }
            }
         },
         "place_id" : "ChIJv4PgeFkZRYgRNVFlWLNnH-U",
         "types" : [ "route" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "State Route 777",
               "short_name" : "State Rte 777",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Linville",
               "short_name" : "Linville",
               "types" : [ "administrative_area_level_3", "political" ]
            },
            {
               "long_name" : "Rockingham County",
               "short_name" : "Rockingham County",
               "types" : [ "administrative_area_level_2", "political" ]
            },
            {
               "long_name" : "Virginia",
               "short_name" : "VA",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "United States",
               "short_name" : "US",
               "types" : [ "country", "political" ]
            }
         ],
         "formatted_address" : "State Rte 777, Linville, VA, USA",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : 38.5501627,
                  "lng" : -78.89086279999999
               },
               "southwest" : {
                  "lat" : 38.5194464,
                  "lng" : -78.9344533
               }
            },
            "location" : {
               "lat" : 38.5325009,
               "lng" : -78.914486
            },
            "location_type" : "GEOMETRIC_CENTER",
            "viewport" : {
               "northeast" : {
                  "lat" : 38.5501627,
                  "lng" : -78.89086279999999
               },
               "southwest" : {
                  "lat" : 38.5194464,
                  "lng" : -78.9344533
               }
            }
         },
         "place_id" : "ChIJvz1gPRrstIkR7rAAQJGTauQ",
         "types" : [ "route" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "State Route 777",
               "short_name" : "State Rte 777",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Stuart",
               "short_name" : "Stuart",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "Patrick County",
               "short_name" : "Patrick County",
               "types" : [ "administrative_area_level_2", "political" ]
            },
            {
               "long_name" : "Virginia",
               "short_name" : "VA",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "United States",
               "short_name" : "US",
               "types" : [ "country", "political" ]
            },
            {
               "long_name" : "24171",
               "short_name" : "24171",
               "types" : [ "postal_code" ]
            }
         ],
         "formatted_address" : "State Rte 777, Stuart, VA 24171, USA",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : 36.6434629,
                  "lng" : -80.2168202
               },
               "southwest" : {
                  "lat" : 36.6176227,
                  "lng" : -80.23326829999999
               }
            },
            "location" : {
               "lat" : 36.6292108,
               "lng" : -80.2197977
            },
            "location_type" : "GEOMETRIC_CENTER",
            "viewport" : {
               "northeast" : {
                  "lat" : 36.6434629,
                  "lng" : -80.2168202
               },
               "southwest" : {
                  "lat" : 36.6176227,
                  "lng" : -80.23326829999999
               }
            }
         },
         "place_id" : "ChIJObNpSIRmUogRz4MNhvZetb8",
         "types" : [ "route" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "State Route 777",
               "short_name" : "State Rte 777",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Fort Blackmore",
               "short_name" : "Fort Blackmore",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "5",
               "short_name" : "5",
               "types" : [ "administrative_area_level_3", "political" ]
            },
            {
               "long_name" : "Scott County",
               "short_name" : "Scott County",
               "types" : [ "administrative_area_level_2", "political" ]
            },
            {
               "long_name" : "Virginia",
               "short_name" : "VA",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "United States",
               "short_name" : "US",
               "types" : [ "country", "political" ]
            },
            {
               "long_name" : "24250",
               "short_name" : "24250",
               "types" : [ "postal_code" ]
            }
         ],
         "formatted_address" : "State Rte 777, Fort Blackmore, VA 24250, USA",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : 36.7913489,
                  "lng" : -82.63536839999999
               },
               "southwest" : {
                  "lat" : 36.7603541,
                  "lng" : -82.6432519
               }
            },
            "location" : {
               "lat" : 36.7760564,
               "lng" : -82.6396834
            },
            "location_type" : "GEOMETRIC_CENTER",
            "viewport" : {
               "northeast" : {
                  "lat" : 36.7913489,
                  "lng" : -82.63536839999999
               },
               "southwest" : {
                  "lat" : 36.7603541,
                  "lng" : -82.6432519
               }
            }
         },
         "place_id" : "ChIJaaVDOXrBWogRXP9Mw0IGvNM",
         "types" : [ "route" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "State Route 777",
               "short_name" : "State Rte 777",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Blacksburg",
               "short_name" : "Blacksburg",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "Montgomery County",
               "short_name" : "Montgomery County",
               "types" : [ "administrative_area_level_2", "political" ]
            },
            {
               "long_name" : "Virginia",
               "short_name" : "VA",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "United States",
               "short_name" : "US",
               "types" : [ "country", "political" ]
            },
            {
               "long_name" : "24060",
               "short_name" : "24060",
               "types" : [ "postal_code" ]
            }
         ],
         "formatted_address" : "State Rte 777, Blacksburg, VA 24060, USA",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : 37.2804819,
                  "lng" : -80.4191335
               },
               "southwest" : {
                  "lat" : 37.274901,
                  "lng" : -80.4417467
               }
            },
            "location" : {
               "lat" : 37.2768529,
               "lng" : -80.43254
            },
            "location_type" : "GEOMETRIC_CENTER",
            "viewport" : {
               "northeast" : {
                  "lat" : 37.2804819,
                  "lng" : -80.4191335
               },
               "southwest" : {
                  "lat" : 37.274901,
                  "lng" : -80.4417467
               }
            }
         },
         "place_id" : "ChIJ6XX0172_TYgRnoOmsZDLX5g",
         "types" : [ "route" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "State Route 777",
               "short_name" : "State Rte 777",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Bedford",
               "short_name" : "Bedford",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "Center",
               "short_name" : "Center",
               "types" : [ "administrative_area_level_3", "political" ]
            },
            {
               "long_name" : "Bedford County",
               "short_name" : "Bedford County",
               "types" : [ "administrative_area_level_2", "political" ]
            },
            {
               "long_name" : "Virginia",
               "short_name" : "VA",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "United States",
               "short_name" : "US",
               "types" : [ "country", "political" ]
            },
            {
               "long_name" : "24523",
               "short_name" : "24523",
               "types" : [ "postal_code" ]
            }
         ],
         "formatted_address" : "State Rte 777, Bedford, VA 24523, USA",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : 37.3267811,
                  "lng" : -79.4632446
               },
               "southwest" : {
                  "lat" : 37.3190651,
                  "lng" : -79.4842304
               }
            },
            "location" : {
               "lat" : 37.32572469999999,
               "lng" : -79.47182529999999
            },
            "location_type" : "GEOMETRIC_CENTER",
            "viewport" : {
               "northeast" : {
                  "lat" : 37.3267811,
                  "lng" : -79.4632446
               },
               "southwest" : {
                  "lat" : 37.3190651,
                  "lng" : -79.4842304
               }
            }
         },
         "place_id" : "ChIJT3w6KPQxTYgR_kiaNTSoVjE",
         "types" : [ "route" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "Louisiana 777",
               "short_name" : "LA-777",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Jena",
               "short_name" : "Jena",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "8",
               "short_name" : "8",
               "types" : [ "administrative_area_level_3", "political" ]
            },
            {
               "long_name" : "La Salle Parish",
               "short_name" : "La Salle Parish",
               "types" : [ "administrative_area_level_2", "political" ]
            },
            {
               "long_name" : "Louisiana",
               "short_name" : "LA",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "United States",
               "short_name" : "US",
               "types" : [ "country", "political" ]
            },
            {
               "long_name" : "71342",
               "short_name" : "71342",
               "types" : [ "postal_code" ]
            }
         ],
         "formatted_address" : "LA-777, Jena, LA 71342, USA",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : 31.57604749999999,
                  "lng" : -92.1451419
               },
               "southwest" : {
                  "lat" : 31.5612252,
                  "lng" : -92.1514545
               }
            },
            "location" : {
               "lat" : 31.5687575,
               "lng" : -92.1480584
            },
            "location_type" : "GEOMETRIC_CENTER",
            "viewport" : {
               "northeast" : {
                  "lat" : 31.57604749999999,
                  "lng" : -92.1451419
               },
               "southwest" : {
                  "lat" : 31.5612252,
                  "lng" : -92.1514545
               }
            }
         },
         "place_id" : "ChIJ_YLjRtRfJYYRi5Lr8Ua6XMk",
         "types" : [ "route" ]
      },
      {
         "address_components" : [
            {
               "long_name" : "State Route 777",
               "short_name" : "State Rte 777",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Nathalie",
               "short_name" : "Nathalie",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "1",
               "short_name" : "1",
               "types" : [ "administrative_area_level_3", "political" ]
            },
            {
               "long_name" : "Halifax County",
               "short_name" : "Halifax County",
               "types" : [ "administrative_area_level_2", "political" ]
            },
            {
               "long_name" : "Virginia",
               "short_name" : "VA",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "United States",
               "short_name" : "US",
               "types" : [ "country", "political" ]
            },
            {
               "long_name" : "24577",
               "short_name" : "24577",
               "types" : [ "postal_code" ]
            }
         ],
         "formatted_address" : "State Rte 777, Nathalie, VA 24577, USA",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : 37.00617980000001,
                  "lng" : -78.99254590000001
               },
               "southwest" : {
                  "lat" : 37.0007656,
                  "lng" : -79.01392700000001
               }
            },
            "location" : {
               "lat" : 37.0027488,
               "lng" : -79.0031595
            },
            "location_type" : "GEOMETRIC_CENTER",
            "viewport" : {
               "northeast" : {
                  "lat" : 37.00617980000001,
                  "lng" : -78.99254590000001
               },
               "southwest" : {
                  "lat" : 37.0007656,
                  "lng" : -79.01392700000001
               }
            }
         },
         "place_id" : "ChIJDesr7DOWsokRAOMVTIlMh0I",
         "types" : [ "route" ]
      }
   ],
   "status" : "OK"
}`))

	httpmock.RegisterResponder("GET", "https://maps.googleapis.com/maps/api/geocode/json?address=777aaa&key=AIzaSyBKJ5hORqnuHg33bzcB7AyZBDr7jKzhwsc",
		httpmock.NewStringResponder(200, `{
   "results" : [],
   "status" : "ZERO_RESULTS"
}`))

	httpmock.RegisterResponder("GET", "https://maps.googleapis.com/maps/api/geocode/json?address=DifferentFailure&key=AIzaSyBKJ5hORqnuHg33bzcB7AyZBDr7jKzhwsc",
		httpmock.NewStringResponder(200, `{
   "results" : [],
   "status" : "NON ZERO RESULTS FAIL"
}`))

	req, err := http.NewRequest("GET", "/latlongByAddress/777", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LatLongByAddress)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `777`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// New test
	req, err = http.NewRequest("GET", "/latlongByAddress/777aaa", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(LatLongByAddress)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected = `ZERO_RESULTS`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// New test
	req, err = http.NewRequest("GET", "/latlongByAddress/DifferentFailure", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(LatLongByAddress)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected = `NON ZERO RESULTS FAIL`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
