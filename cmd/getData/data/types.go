package data

import (
	"database/sql"
	"database/sql/driver"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/google/uuid"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

type UnitType string

const (
	APARTMENT           UnitType = "APARTMENT"
	TOWNHOME            UnitType = "TOWNHOME"
	HOUSE               UnitType = "HOUSE"
	CONDOMINIUM         UnitType = "CONDOMINIUM"
	COMMERCIAL          UnitType = "COMMERCIAL"
	STORAGE             UnitType = "STORAGE"
	GARAGE              UnitType = "GARAGE"
	UNDERGROUND_PARKING UnitType = "UNDERGROUND_PARKING"
	OUTDOOR_PARKING     UnitType = "OUTDOOR_PARKING"
	AMENITY             UnitType = "AMENITY"
	UTILITY             UnitType = "UTILITY"
)

//type UUID [16]byte

type Unit struct {
	Id         string         `json:"id"`
	Zone       string         `json:"zone"`
	Name       string         `json:"name"`
	Floor      sql.NullInt64  `json:"floor"`
	UnitType   UnitType       `json:"unitType"`
	UnitNumber string         `json:"unitNumber"`
	FloorPlan  sql.NullString `json="floorPlan"`
	SquareFeet sql.NullInt64  `json:"squareFeet"`
}

type Zone struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Property string `json:"property"`
}

type Property struct {
	Id                 string         `json:"id"`
	Name               string         `json:"name"`
	PrimaryUnitType    string         `json:"primaryUnitType"`
	DateBuilt          NullTime       `json:"dateBuilt"`
	DateRemodeled      NullTime       `json:"dateRemodeled"`
	IdentificationCode string         `json:"identificationCode"`
	Community          sql.NullString `json:"community"`
}

type Community struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type FloorPlan struct {
	Id                string          `json:"id"`
	Name              string          `json:"name"`
	Bedrooms          sql.NullFloat64 `json:"bedrooms"`
	Bathrooms         sql.NullFloat64 `json:"bathrooms"`
	PropertyMarketing sql.NullString  `json:"community"`
}

type AmenityType string

const (
	PROPERTY   AmenityType = "PROPERTY"
	FLOOR_PLAN AmenityType = "FLOOR_PLAN"
	UNIT       AmenityType = "UNIT"
)

type Amenity struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	AmenityType AmenityType `json:"amenityType"`
}

type AvailableUnitsView struct {
	Unit             string   `json:"unit"`
	MoveOutDate      NullTime `json:"moveOutDate"`
	FinancialEndDate NullTime `json:"financialEndDate"`
	VacantDate       NullTime `json:"vacantDate"`
	MoveInDate       NullTime `json:"moveInDate"`
}

type FloorPlanAmenity struct {
	Amenity   string `json:"amenity"`
	FloorPlan string `json:"floorPlan"`
}

type District struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Lead struct {
	Id           string         `json:"id"`
	DateReceived NullTime       `json:"dateReceived"`
	PhoneNumber  sql.NullString `json:"phoneNumber"`
	EmailAddress sql.NullString `json:"emailAddress"`
	TrackingType sql.NullString `json:"trackingType"`
	District     sql.NullString `json:"district"`
}

type Lease struct {
	Id string `json:"id"`
}

type LeaseEndDateView struct {
	Lease   string   `json:"lease"`
	EndDate NullTime `json:"endDate"`
}

type LeaseExpirationDateView struct {
	Lease          string   `json:"lease"`
	ExpirationDate NullTime `json:"expirationDate"`
}

type LeasePlace struct {
	Id               string   `json:"id"`
	StartDate        NullTime `json:"startDate"`
	Lease            string   `json:"lease"`
	Place            string   `json:"place"`
	FinancialEndDate NullTime `json:"financialEndDate"`
	Status           string   `json:"status"`
}

type LeasePrimaryPlaceView struct {
	Lease        string `json:"lease"`
	PrimaryPlace string `json:"unit"`
}

/*
type PricingGroup struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
*/
