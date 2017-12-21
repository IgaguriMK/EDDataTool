package journal

import "time"

type Scan struct {
	AbsoluteMagnitude     float64 `json:"AbsoluteMagnitude"`
	AgeMY                 int64   `json:"Age_MY"`
	Atmosphere            string  `json:"Atmosphere"`
	AtmosphereComposition []struct {
		Name    string  `json:"Name"`
		Percent float64 `json:"Percent"`
	} `json:"AtmosphereComposition"`
	AtmosphereType        string  `json:"AtmosphereType"`
	AxialTilt             float64 `json:"AxialTilt"`
	BodyName              string  `json:"BodyName"`
	DistanceFromArrivalLS float64 `json:"DistanceFromArrivalLS"`
	Eccentricity          float64 `json:"Eccentricity"`
	Landable              bool    `json:"Landable"`
	Luminosity            string  `json:"Luminosity"`
	MassEM                float64 `json:"MassEM"`
	Materials             []struct {
		Name    string  `json:"Name"`
		Percent float64 `json:"Percent"`
	} `json:"Materials"`
	OrbitalInclination float64 `json:"OrbitalInclination"`
	OrbitalPeriod      float64 `json:"OrbitalPeriod"`
	Periapsis          float64 `json:"Periapsis"`
	PlanetClass        string  `json:"PlanetClass"`
	Radius             float64 `json:"Radius"`
	ReserveLevel       string  `json:"ReserveLevel"`
	Rings              []struct {
		InnerRad  float64 `json:"InnerRad"`
		MassMT    float64 `json:"MassMT"`
		Name      string  `json:"Name"`
		OuterRad  float64 `json:"OuterRad"`
		RingClass string  `json:"RingClass"`
	} `json:"Rings"`
	RotationPeriod     float64   `json:"RotationPeriod"`
	SemiMajorAxis      float64   `json:"SemiMajorAxis"`
	StarType           string    `json:"StarType"`
	StellarMass        float64   `json:"StellarMass"`
	SurfaceGravity     float64   `json:"SurfaceGravity"`
	SurfacePressure    float64   `json:"SurfacePressure"`
	SurfaceTemperature float64   `json:"SurfaceTemperature"`
	TerraformState     string    `json:"TerraformState"`
	TidalLock          bool      `json:"TidalLock"`
	Volcanism          string    `json:"Volcanism"`
	Event              string    `json:"event"`
	Timestamp          time.Time `json:"timestamp"`
}

func (e Scan) GetEvent() string {
	return e.Event
}

func (e Scan) GetTimestamp() time.Time {
	return e.Timestamp
}
