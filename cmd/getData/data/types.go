package data

import (
	"database/sql"
	"database/sql/driver"
	"time"

	//github Imports:
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/google/uuid"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//NullTime is to prevent getting an error when the time has null value
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

// UnitType can be apartment or garage or ...
type UnitType string

//types of Unit:
const (
	APARTMENT          UnitType = "APARTMENT"
	TOWNHOME           UnitType = "TOWNHOME"
	HOUSE              UnitType = "HOUSE"
	CONDOMINIUM        UnitType = "CONDOMINIUM"
	COMMERCIAL         UnitType = "COMMERCIAL"
	STORAGE            UnitType = "STORAGE"
	GARAGE             UnitType = "GARAGE"
	UndergroundParking UnitType = "UNDERGROUND_PARKING"
	OutdoorParking     UnitType = "OUTDOOR_PARKING"
	AMENITY            UnitType = "AMENITY"
	UTILITY            UnitType = "UTILITY"
)

//Unit is an apartment unit consists of id, zone which is another entty, and etc
type Unit struct {
	ID         string         `json:"id"`
	Zone       string         `json:"zone"`
	Name       string         `json:"name"`
	Floor      sql.NullInt64  `json:"floor"`
	UnitType   UnitType       `json:"unitType"`
	UnitNumber string         `json:"unitNumber"`
	FloorPlan  sql.NullString `json:"floorPlan"`
	SquareFeet sql.NullInt64  `json:"squareFeet"`
}

//Zone is a group of buildings or storages or garages
type Zone struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Property string `json:"property"`
}

//Property is a building
type Property struct {
	ID                 string         `json:"id"`
	Name               string         `json:"name"`
	PrimaryUnitType    string         `json:"primaryUnitType"`
	DateBuilt          NullTime       `json:"dateBuilt"`
	DateRemodeled      NullTime       `json:"dateRemodeled"`
	IdentificationCode string         `json:"identificationCode"`
	Community          sql.NullString `json:"community"`
}

//Community is a group of buildings
type Community struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//FloorPlan includes number of bedrooms and bathrooms, and etc
type FloorPlan struct {
	ID                string          `json:"id"`
	Name              string          `json:"name"`
	Bedrooms          sql.NullFloat64 `json:"bedrooms"`
	Bathrooms         sql.NullFloat64 `json:"bathrooms"`
	PropertyMarketing sql.NullString  `json:"community"`
}

//AmenityType shows what kind of Amenity is this
type AmenityType string

//different types of amenity
const (
	PROPERTY   AmenityType = "PROPERTY"
	FLOOR_PLAN AmenityType = "FLOOR_PLAN"
	UNIT       AmenityType = "UNIT"
)

//Amenity includes id and name and type
type Amenity struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	AmenityType AmenityType `json:"amenityType"`
}

//AvailableUnitsView includes dates and units
type AvailableUnitsView struct {
	Unit             string   `json:"unit"`
	MoveOutDate      NullTime `json:"moveOutDate"`
	FinancialEndDate NullTime `json:"financialEndDate"`
	VacantDate       NullTime `json:"vacantDate"`
	MoveInDate       NullTime `json:"moveInDate"`
}

//FloorPlanAmenity consists of Amenity ids and FloorPlan ids
type FloorPlanAmenity struct {
	Amenity   string `json:"amenity"`
	FloorPlan string `json:"floorPlan"`
}

//District includes id and name
type District struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Lead includes dates, eail, phone number, and etc
type Lead struct {
	ID           string         `json:"id"`
	DateReceived NullTime       `json:"dateReceived"`
	PhoneNumber  sql.NullString `json:"phoneNumber"`
	EmailAddress sql.NullString `json:"emailAddress"`
	TrackingType sql.NullString `json:"trackingType"`
	District     sql.NullString `json:"district"`
}

//Lease includes only Lease
type Lease struct {
	ID string `json:"id"`
}

//LeaseEndDateView includes Lease and EndDate
type LeaseEndDateView struct {
	Lease   string   `json:"lease"`
	EndDate NullTime `json:"endDate"`
}

//LeaseExpirationDateView includes Lease and ExpirationDate
type LeaseExpirationDateView struct {
	Lease          string   `json:"lease"`
	ExpirationDate NullTime `json:"expirationDate"`
}

//LeasePlace includes Lease, Place, and dates
type LeasePlace struct {
	ID               string   `json:"id"`
	StartDate        NullTime `json:"startDate"`
	Lease            string   `json:"lease"`
	Place            string   `json:"place"`
	FinancialEndDate NullTime `json:"financialEndDate"`
	Status           string   `json:"status"`
}

//LeasePrimaryPlaceView has lease id and unit id
type LeasePrimaryPlaceView struct {
	Lease        string `json:"lease"`
	PrimaryPlace string `json:"unit"`
}

//MoveIn include name of the community, number of bedrooms and bathrooms and moveInDate
type MoveIn struct {
	Name       string          `json:"name"`
	Bedrooms   sql.NullFloat64 `json:"bedrooms"`
	Bathrooms  sql.NullFloat64 `json:"bathrooms"`
	MoveInDate NullTime        `json:"moveInDate"`
}

/*
type PricingGroup struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
*/
