package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"

	myTypes "github.com/mikleing/beaverGO/data"
)

type Unit = myTypes.Unit
type Zone = myTypes.Zone
type Community = myTypes.Community
type FloorPlan = myTypes.FloorPlan
type Property = myTypes.Property
type Amenity = myTypes.Amenity
type AvailableUnitsView = myTypes.AvailableUnitsView
type FloorPlanAmenity = myTypes.FloorPlanAmenity
type District = myTypes.District
type Lead = myTypes.Lead
type Lease = myTypes.Lease
type LeaseEndDateView = myTypes.LeaseEndDateView
type LeaseExpirationDateView = myTypes.LeaseExpirationDateView
type LeasePlace = myTypes.LeasePlace

var dbCon *sql.DB
var err error
var units []Unit
var zones []Zone
var properties []Property
var communities []Community
var floorPlans []FloorPlan
var amenities []Amenity
var availableUnitsViews []AvailableUnitsView
var floorPlanAmenities []FloorPlanAmenity
var districts []District
var leads []Lead
var leases []Lease
var leaseEndDateViews []LeaseEndDateView
var leaseExpirationDateViews []LeaseExpirationDateView
var leasePlaces []LeasePlace

func getUnits(w http.ResponseWriter, r *http.Request) {
	//---------------------------------------------------
	f, err := os.OpenFile("units.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()
	//-----------------------------------------------------
	results, err := dbCon.Query("SELECT Name, Floor, Id, Zone, UnitType, UnitNumber, FloorPlan, SquareFeet from Unit")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var unit Unit
		err = results.Scan(&unit.Name, &unit.Floor, &unit.Id, &unit.Zone, &unit.UnitType, &unit.UnitNumber, &unit.FloorPlan, &unit.SquareFeet)
		if err != nil {
			panic(err.Error())
		}
		units = append(units, unit)
		//---------------------------------------------
		b, err := json.Marshal(unit)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		f.Write(b)

		//---------------------------------------------
	}
	//------
	f.Close()
	//------
	fmt.Println("Endpoint Hit: All Units Endpoint")
	json.NewEncoder(w).Encode(units)
}

func getZones(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Id, Name, Property from Zone")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var zone Zone
		err = results.Scan(&zone.Id, &zone.Name, &zone.Property)
		if err != nil {
			panic(err.Error())
		}
		zones = append(zones, zone)
	}
	fmt.Println("Endpoint Hit: All Zones Endpoint")
	json.NewEncoder(w).Encode(zones)
}

func getProperties(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Id, Name, PrimaryUnitType, DateBuilt, DateRemodeled, IdentificationCode, Community  from Property")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var property Property
		err = results.Scan(&property.Id, &property.Name, &property.PrimaryUnitType, &property.DateBuilt, &property.DateRemodeled, &property.IdentificationCode, &property.Community)
		if err != nil {
			panic(err.Error())
		}
		properties = append(properties, property)
	}
	fmt.Println("Endpoint Hit: All Properties Endpoint")
	json.NewEncoder(w).Encode(properties)
}

func getCommunities(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Id, Name from Community")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var community Community
		err = results.Scan(&community.Id, &community.Name)
		if err != nil {
			panic(err.Error())
		}
		communities = append(communities, community)
	}
	fmt.Println("Endpoint Hit: All Communities Endpoint")
	json.NewEncoder(w).Encode(communities)
}

func getFloorPlans(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Id, Name, Bedrooms, Bathrooms, PropertyMarketing from FloorPlan")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var floorPlan FloorPlan
		err = results.Scan(&floorPlan.Id, &floorPlan.Name, &floorPlan.Bedrooms, &floorPlan.Bathrooms, &floorPlan.PropertyMarketing)
		if err != nil {
			panic(err.Error())
		}
		floorPlans = append(floorPlans, floorPlan)
	}
	fmt.Println("Endpoint Hit: All FloorPlans Endpoint")
	json.NewEncoder(w).Encode(floorPlans)
}

func getAmenities(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Id, Name, AmenityType from Amenity")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var amenity Amenity
		err = results.Scan(&amenity.Id, &amenity.Name, &amenity.AmenityType)
		if err != nil {
			panic(err.Error())
		}
		amenities = append(amenities, amenity)
	}
	fmt.Println("Endpoint Hit: All Amenities Endpoint")
	json.NewEncoder(w).Encode(amenities)
}

func getAvailableUnitsViews(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Unit, MoveOutDate, FinancialEndDate, VacantDate, MoveInDate from AvailableUnitsView")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var availableUnitsView AvailableUnitsView
		err = results.Scan(&availableUnitsView.Unit, &availableUnitsView.MoveOutDate, &availableUnitsView.FinancialEndDate, &availableUnitsView.VacantDate, &availableUnitsView.MoveInDate)
		if err != nil {
			panic(err.Error())
		}
		availableUnitsViews = append(availableUnitsViews, availableUnitsView)
	}
	fmt.Println("Endpoint Hit: All AvailableUnitsViews Endpoint")
	json.NewEncoder(w).Encode(availableUnitsViews)
}

func getFloorPlanAmenities(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Amenity, FloorPlan from FloorPlanAmenity")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var floorPlanAmenity FloorPlanAmenity
		err = results.Scan(&floorPlanAmenity.Amenity, &floorPlanAmenity.FloorPlan)
		if err != nil {
			panic(err.Error())
		}
		floorPlanAmenities = append(floorPlanAmenities, floorPlanAmenity)
	}
	fmt.Println("Endpoint Hit: All FloorPlanAmenities Endpoint")
	json.NewEncoder(w).Encode(floorPlanAmenities)
}

func getDistricts(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Id, Name from District")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var district District
		err = results.Scan(&district.Id, &district.Name)
		if err != nil {
			panic(err.Error())
		}
		districts = append(districts, district)
	}
	fmt.Println("Endpoint Hit: All Districts Endpoint")
	json.NewEncoder(w).Encode(districts)
}

func getLeads(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT DateReceived, PhoneNumber, EmailAddress, TrackingType, District from Lead")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var lead Lead
		err = results.Scan(&lead.DateReceived, &lead.PhoneNumber, &lead.EmailAddress, &lead.TrackingType, &lead.District)
		if err != nil {
			panic(err.Error())
		}
		leads = append(leads, lead)
	}
	fmt.Println("Endpoint Hit: All Leads Endpoint")
	json.NewEncoder(w).Encode(leads)
}

func getLeases(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Id from Lease")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var lease Lease
		err = results.Scan(&lease.Id)
		if err != nil {
			panic(err.Error())
		}
		leases = append(leases, lease)
	}
	fmt.Println("Endpoint Hit: All Leases Endpoint")
	json.NewEncoder(w).Encode(leases)
}

func getLeaseEndDateViews(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Lease, EndDate from LeaseEndDateView")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var leaseEndDateView LeaseEndDateView
		err = results.Scan(&leaseEndDateView.Lease, &leaseEndDateView.EndDate)
		if err != nil {
			panic(err.Error())
		}
		leaseEndDateViews = append(leaseEndDateViews, leaseEndDateView)
	}
	fmt.Println("Endpoint Hit: All LeaseEndDateViews Endpoint")
	json.NewEncoder(w).Encode(leaseEndDateViews)
}

func getLeaseExpirationDateViews(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT Lease, ExpirationDate from LeaseExpirationDateView")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var leaseExpirationDateView LeaseExpirationDateView
		err = results.Scan(&leaseExpirationDateView.Lease, &leaseExpirationDateView.ExpirationDate)
		if err != nil {
			panic(err.Error())
		}
		leaseExpirationDateViews = append(leaseExpirationDateViews, leaseExpirationDateView)
	}
	fmt.Println("Endpoint Hit: All LeaseExpirationDateViews Endpoint")
	json.NewEncoder(w).Encode(leaseExpirationDateViews)
}

func getLeasePlaces(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT StartDate, Lease, Place, FinancialEndDate, Status from LeasePlace")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var leasePlace LeasePlace
		err = results.Scan(&leasePlace.StartDate, &leasePlace.Lease, &leasePlace.Place, &leasePlace.FinancialEndDate, &leasePlace.Status)
		if err != nil {
			panic(err.Error())
		}
		leasePlaces = append(leasePlaces, leasePlace)
	}
	fmt.Println("Endpoint Hit: All LeasePlaces Endpoint")
	json.NewEncoder(w).Encode(leasePlaces)
}

func dbConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "negar:E%ycRw4MxjR6u!M2YpvDN9Cq6d^tT58n@tcp(production.cm4fwnwaa3mf.us-east-1.rds.amazonaws.com:3306)/production?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db, err
}

func main() {
	dbCon, err = dbConnect()
	defer dbCon.Close()
	fmt.Println("connect success")
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "How you doin?")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/units", getUnits).Methods("GET")
	myRouter.HandleFunc("/zones", getZones).Methods("GET")
	myRouter.HandleFunc("/properties", getProperties).Methods("GET")
	myRouter.HandleFunc("/communities", getCommunities).Methods("GET")
	myRouter.HandleFunc("/floorPlans", getFloorPlans).Methods("GET")
	myRouter.HandleFunc("/amenities", getAmenities).Methods("GET")
	myRouter.HandleFunc("/availableUnitsViews", getAvailableUnitsViews).Methods("GET")
	myRouter.HandleFunc("/floorPlanAmenities", getFloorPlanAmenities).Methods("GET")
	myRouter.HandleFunc("/districts", getDistricts).Methods("GET")
	myRouter.HandleFunc("/leads", getLeads).Methods("GET")
	myRouter.HandleFunc("/leases", getLeases).Methods("GET")
	myRouter.HandleFunc("/leaseEndDateViews", getLeaseEndDateViews).Methods("GET")
	myRouter.HandleFunc("/leaseExpirationDateViews", getLeaseExpirationDateViews).Methods("GET")
	myRouter.HandleFunc("/leasePlaces", getLeasePlaces).Methods("GET")
	//myRouter.HandleFunc("/units", postUnits).Methods("POST")
	//myRouter.HandleFunc("/units/{name}", oneUnit).Methods("GET")
	log.Fatal(http.ListenAndServe(":8083", myRouter))
}

//func postUnits(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "Test POST endpoint worked")
//}

//func oneUnit(w http.ResponseWriter, r *http.Request) {
//    params := mux.Vars(r)
//    for _, unit := range units {
//        if unit.Name == params["name"] {
//            fmt.Println("Endpoint Hit: One Unit Endpoint")
//            json.NewEncoder(w).Encode(unit)
//        }
//    }
//}
