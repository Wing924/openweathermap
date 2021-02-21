package openweathermap

// RoadRisk
type RoadRisk struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Dt  int64   `json:"dt"`
}

// RoadRiskRequest
type RoadRiskRequest struct {
	RoadRisk []RoadRisk `json:"roadrisk"`
}

// [
//   {
//     "dt": 1602702000,
//     "coord": [
//       7.27,
//       44.04
//     ],
//     "weather": {
//       "temp": 278.44,
//       "wind_speed": 2.27,
//       "wind_deg": 7,
//       "precipitation_intensity": 0.38,
//       "dew_point": 276.13
//     },
//     "alerts": [
//       {
//         "sender_name": "METEO-FRANCE",
//         "event": "Moderate thunderstorm warning",
//         "event_level": 2
//       }
//     ]
//   },
//   {
//     "dt": 1602702400,
//     "coord": [
//       7.37,
//       45.04
//     ],
//     "weather": {
//       "temp": 282.44,
//       "wind_speed": 1.84,
//       "wind_deg": 316,
//       "dew_point": 275.99
//     },
//     "alerts": [

//     ]
//   }
// ]
type RoadRiskResponse struct {
	Alerts []Alert `json:"alerts"`
}
