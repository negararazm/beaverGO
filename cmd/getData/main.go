package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"

	myTypes "github.com/mikleing/beaverGO/cmd/getData/data"
)

//Unit is an apartment unit consists of id, zone which is another entty, and etc
type Unit = myTypes.Unit

//Zone is a group of buildings or storages or garages
type Zone = myTypes.Zone

//Community is a group of buildings
type Community = myTypes.Community

//FloorPlan includes number of bedrooms and bathrooms, and etc
type FloorPlan = myTypes.FloorPlan

//Property is a building
type Property = myTypes.Property

//Amenity includes id and name and type
type Amenity = myTypes.Amenity

//AvailableUnitsView includes dates and units
type AvailableUnitsView = myTypes.AvailableUnitsView

//FloorPlanAmenity consists of Amenity ids and FloorPlan ids
type FloorPlanAmenity = myTypes.FloorPlanAmenity

//District includes id and name
type District = myTypes.District

//Lead includes dates, eail, phone number, and etc
type Lead = myTypes.Lead

//Lease includes only Lease
type Lease = myTypes.Lease

//LeaseEndDateView includes Lease and EndDate
type LeaseEndDateView = myTypes.LeaseEndDateView

//LeaseExpirationDateView includes Lease and ExpirationDate
type LeaseExpirationDateView = myTypes.LeaseExpirationDateView

//LeasePlace includes Lease, Place, and dates
type LeasePlace = myTypes.LeasePlace

//LeasePrimaryPlaceView has lease id and unit id
type LeasePrimaryPlaceView = myTypes.LeasePrimaryPlaceView

//MoveIn include name of the community, number of bedrooms and bathrooms and moveInDate
type MoveIn = myTypes.MoveIn

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
var leasePrimaryPlaceViews []LeasePrimaryPlaceView
var moveIns []MoveIn

func getUnits(w http.ResponseWriter, r *http.Request) {
	//---------------------------------------------------
	f, err := os.OpenFile("units.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer f.Close()
	//-----------------------------------------------------
	results, err := dbCon.Query("SELECT Name, Floor, HEX(Id), HEX(Zone), UnitType, UnitNumber, HEX(FloorPlan), SquareFeet from Unit")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var unit Unit
		err = results.Scan(&unit.Name, &unit.Floor, &unit.ID, &unit.Zone, &unit.UnitType, &unit.UnitNumber, &unit.FloorPlan, &unit.SquareFeet)
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
	results, err := dbCon.Query("SELECT HEX(Id), Name, HEX(Property) from Zone")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var zone Zone
		err = results.Scan(&zone.ID, &zone.Name, &zone.Property)
		if err != nil {
			panic(err.Error())
		}
		zones = append(zones, zone)
	}
	fmt.Println("Endpoint Hit: All Zones Endpoint")
	json.NewEncoder(w).Encode(zones)
}

func getProperties(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT HEX(Id), Name, PrimaryUnitType, DateBuilt, DateRemodeled, IdentificationCode, HEX(Community)  from Property")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var property Property
		err = results.Scan(&property.ID, &property.Name, &property.PrimaryUnitType, &property.DateBuilt, &property.DateRemodeled, &property.IdentificationCode, &property.Community)
		if err != nil {
			panic(err.Error())
		}
		properties = append(properties, property)
	}
	fmt.Println("Endpoint Hit: All Properties Endpoint")
	json.NewEncoder(w).Encode(properties)
}

func getCommunities(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT HEX(Id), Name from Community")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var community Community
		err = results.Scan(&community.ID, &community.Name)
		if err != nil {
			panic(err.Error())
		}
		communities = append(communities, community)
	}
	fmt.Println("Endpoint Hit: All Communities Endpoint")
	json.NewEncoder(w).Encode(communities)
}

func getFloorPlans(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT HEX(Id), Name, Bedrooms, Bathrooms, HEX(PropertyMarketing) from FloorPlan")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var floorPlan FloorPlan
		err = results.Scan(&floorPlan.ID, &floorPlan.Name, &floorPlan.Bedrooms, &floorPlan.Bathrooms, &floorPlan.PropertyMarketing)
		if err != nil {
			panic(err.Error())
		}
		floorPlans = append(floorPlans, floorPlan)
	}
	fmt.Println("Endpoint Hit: All FloorPlans Endpoint")
	json.NewEncoder(w).Encode(floorPlans)
}

func getAmenities(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT HEX(Id), Name, AmenityType from Amenity")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var amenity Amenity
		err = results.Scan(&amenity.ID, &amenity.Name, &amenity.AmenityType)
		if err != nil {
			panic(err.Error())
		}
		amenities = append(amenities, amenity)
	}
	fmt.Println("Endpoint Hit: All Amenities Endpoint")
	json.NewEncoder(w).Encode(amenities)
}

func getAvailableUnitsViews(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT HEX(Unit), MoveOutDate, FinancialEndDate, VacantDate, MoveInDate from AvailableUnitsView")
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
	results, err := dbCon.Query("SELECT HEX(Amenity), HEX(FloorPlan) from FloorPlanAmenity")
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
	results, err := dbCon.Query("SELECT HEX(Id), Name from District")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var district District
		err = results.Scan(&district.ID, &district.Name)
		if err != nil {
			panic(err.Error())
		}
		districts = append(districts, district)
	}
	fmt.Println("Endpoint Hit: All Districts Endpoint")
	json.NewEncoder(w).Encode(districts)
	test()
}

func getLeads(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT DateReceived, PhoneNumber, EmailAddress, TrackingType, HEX(District) from Lead")
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
	results, err := dbCon.Query("SELECT HEX(Id) from Lease")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var lease Lease
		err = results.Scan(&lease.ID)
		if err != nil {
			panic(err.Error())
		}
		leases = append(leases, lease)
	}
	fmt.Println("Endpoint Hit: All Leases Endpoint")
	json.NewEncoder(w).Encode(leases)
}

func getLeaseEndDateViews(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT HEX(Lease), EndDate from LeaseEndDateView")
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
	results, err := dbCon.Query("SELECT HEX(Lease), ExpirationDate from LeaseExpirationDateView")
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
	results, err := dbCon.Query("SELECT StartDate, HEX(Lease), HEX(Place), FinancialEndDate, Status from LeasePlace")
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

func getLeasePrimaryPlaceView(w http.ResponseWriter, r *http.Request) {
	results, err := dbCon.Query("SELECT HEX(Lease), HEX(PrimaryPlace) from LeasePrimaryPlaceView")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var leasePrimaryPlaceView LeasePrimaryPlaceView
		err = results.Scan(&leasePrimaryPlaceView.Lease, &leasePrimaryPlaceView.PrimaryPlace)
		if err != nil {
			panic(err.Error())
		}
		leasePrimaryPlaceViews = append(leasePrimaryPlaceViews, leasePrimaryPlaceView)
	}
	fmt.Println("Endpoint Hit: All LeasePlaces Endpoint")
	json.NewEncoder(w).Encode(leasePrimaryPlaceViews)
}

//------------------
func getMoveIns(w http.ResponseWriter, r *http.Request) {
	query := `select c.name, f.bedrooms, f.bathrooms, min(p.startDate) 
	from Lease l 
	join LeasePlace p on p.lease = l.id 
	join LeasePrimaryPlaceView v on v.lease = l.id 
	join Unit u on v.primaryPlace = u.id 
	join FloorPlan f on f.id = u.floorPlan 
	join Community c on f.propertyMarketing = c.id 
	where (u.unitType like "APARTMENT" or u.unitType like "TOWNHOME" OR u.unitType like "HOUSE")
	group by l.id
	order by c.name, f.bedrooms, f.bathrooms, min(p.startDate);`

	results, err := dbCon.Query(query)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var moveIn MoveIn
		err = results.Scan(&moveIn.Name, &moveIn.Bedrooms, &moveIn.Bathrooms, &moveIn.MoveInDate)
		if err != nil {
			panic(err.Error())
		}
		moveIns = append(moveIns, moveIn)
	}

	fmt.Println("Endpoint Hit: All MoveIns Endpoint")
	json.NewEncoder(w).Encode(moveIns)
	test()
}

//--------------------

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
	//test()
	handleRequests()
}

func test() {
	n := 3.0
	m := make(map[int]int)
	loc, _ := time.LoadLocation("UTC")
	dateNow := time.Now().In(loc)
	i := 0
	for _, element := range moveIns {
		//fmt.Println(element.Name)
		i++
		if element.MoveInDate.Valid == true {
			diff := dateNow.Sub(element.MoveInDate.Time).Hours() / 24
			if element.MoveInDate.Time.Before(dateNow) && diff < (365*n) {
				//if element.Name == "Emberwood Apartments" && element.Bedrooms.Float64 == 1.0 {
				//fmt.Println(m[int(element.MoveInDate.Time.Month())][2])
				//}
				m[int(element.MoveInDate.Time.Month())] = m[int(element.MoveInDate.Time.Month())] + 1
				//fmt.Println(m[int(element.MoveInDate.Time.Month())])
				//fmt.Println(diff)
				//fmt.Println(int(element.MoveInDate.Time.Month()))
			}

		}
	}
	fmt.Println(m)
	fmt.Println(i)
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
	myRouter.HandleFunc("/leasePrimaryPlaceViews", getLeasePrimaryPlaceView).Methods("GET")
	myRouter.HandleFunc("/moveIns", getMoveIns).Methods("GET")
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
