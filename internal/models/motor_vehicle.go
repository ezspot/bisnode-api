package models

// MotorVehicleSearchResponse represents the response from the motor vehicle search API
type MotorVehicleSearchResponse struct {
	Result []MotorVehicle `json:"Result"`
	Service struct {
		Dataset       string `json:"dataset"`
		Documentation string `json:"documentation"`
		Version      string `json:"version"`
		Timestamp    string `json:"timestamp"`
		Message      string `json:"message"`
	} `json:"Service"`
}

// MotorVehicle represents a motor vehicle record
type MotorVehicle struct {
	RegNo          string `json:"regno"`
	PersonalPlates string `json:"personalplates"`
	ChassisNo      string `json:"chassisno"`
	RegYear        string `json:"regyear"`
	GroupNo        int    `json:"groupno"`
	ModelYear      string `json:"modelyear"`
	BrandNo        int    `json:"brandno"`
	BrandName      string `json:"brandname"`
	Model          string `json:"model"`
	RegDate        string `json:"regdate"`
	OrganizationNo string `json:"organizationno"`
	ReregDate      string `json:"reregdate"`
	UnregDate      string `json:"unregdate"`
	ScrapDate      string `json:"scrapdate"`
	FinalScrapDate string `json:"finalscrapdate"`
	LastInspectionDate string `json:"lastinspectiondate"`
	NextInspectionDate string `json:"nextinspectiondate"`
	Color          string `json:"color"`
	ColorText      string `json:"colortext"`
	UsedImport     int    `json:"usedimport"`
	SeatsTotal     int    `json:"seatstotal"`
	SeatsFront     int    `json:"seatsfront"`
	OwnerRegDate   string `json:"ownerregdate"`
	OwnerChangeDate string `json:"ownerchangedate"`
	OwnerUnregDate string `json:"ownerunregdate"`
	PlateColor     string `json:"platecolor"`
	RegStatus      string `json:"regstatus"`
	NumDoors       string `json:"numdoors"`
	NatureOfDriving string `json:"natureofdriving"`
	Supplement     string `json:"supplement"`
	
	EngineAndTransmission EngineAndTransmission `json:"EngineAndTransmission"`
	AxleTiresAndRims     AxleTiresAndRims     `json:"AxleTiresAndRims"`
	WeightsAndMeasures   WeightsAndMeasures   `json:"WeightsAndMeasures"`
	EU                   EU                   `json:"EU"`
	Owner                Owner                `json:"Owner"`
	CoOwner              Owner                `json:"CoOwner,omitempty"`
	LeasingUser          Owner                `json:"LeasingUser,omitempty"`
}

type EngineAndTransmission struct {
	Transmission  string `json:"transmission"`
	Fuel          string `json:"fuel"`
	FuelText      string `json:"fueltext"`
	EngineVolume  string `json:"enginevolume"`
	EnginePower   string `json:"enginepower"`
	MotorCode     string `json:"motorcode"`
	Hybrid        string `json:"hybrid"`
	HybridCat     string `json:"hybridcat"`
}

type AxleTiresAndRims struct {
	TyreDim1        string `json:"tyredim1"`
	TyreDim2        string `json:"tyredim2"`
	TyreDim3        string `json:"tyredim3,omitempty"`
	TyreLoadIndex1  string `json:"tyreloadindex1"`
	TyreLoadIndex2  string `json:"tyreloadindex2"`
	TyreLoadIndex3  string `json:"tyreloadindex3,omitempty"`
	TyreSpeedIndex1 string `json:"tyrespeedindex1"`
	TyreSpeedIndex2 string `json:"tyrespeedindex2"`
	TyreSpeedIndex3 string `json:"tyrespeedindex3,omitempty"`
	BimOffset1      string `json:"bimoffset1"`
	BimOffset2      string `json:"bimoffset2"`
	BimOffset3      string `json:"bimoffset3,omitempty"`
	AxleSpread1     string `json:"axlespread1,omitempty"`
	AxleSpread2     string `json:"axlespread2,omitempty"`
	Axles           int    `json:"axles"`
	AxleOperation   int    `json:"axleoperation"`
	RimDim1         string `json:"rimdim1"`
	RimDim2         string `json:"rimdim2"`
	RimDim3         string `json:"rimdim3,omitempty"`
	TrackWideAxle1  string `json:"trackwideaxle1"`
	TrackWideAxle2  string `json:"trackwideaxle2"`
	TrackWideAxle3  string `json:"trackwideaxle3,omitempty"`
	AirSusp1        string `json:"airsusp1,omitempty"`
	AirSusp2        string `json:"airsusp2,omitempty"`
	AirSusp3        string `json:"airsusp3,omitempty"`
}

type WeightsAndMeasures struct {
	TotalWeight               int    `json:"totalweight"`
	Status                    int    `json:"status"`
	AxleWeightLimit1          string `json:"axleweightlimit1"`
	AxleWeightLimit2          string `json:"axleweightlimit2"`
	AxleWeightLimit3          string `json:"axleweightlimit3,omitempty"`
	RoofWeightLimit           int    `json:"roofweightlimit"`
	CurbWeight                int    `json:"curbweight"`
	TrailerWeightWithBreaks   int    `json:"trailerweightwithbreaks"`
	TrailerWeightWithoutBreaks int    `json:"trailerweightwithoutbreaks"`
	CouplingLoad              int    `json:"couplingload"`
	MaxWeight                 int    `json:"maxweight"`
	Length                    int    `json:"length"`
	Width                     int    `json:"width"`
	StandNoice                string `json:"standnoice"`
	ParticleFilter            int    `json:"particlefilter"`
	NOXEmissionsGPRKWH        string `json:"nox_emissions_gprkwh,omitempty"`
	NOXEmissionsMGPRKH        string `json:"nox_emissions_mgprkh,omitempty"`
	ParticleEmissions         string `json:"particleemissions,omitempty"`
	MeasurementMethod         string `json:"measurementmethod"`
	CO2Emission               string `json:"co2_emission"`
	FuelEconomy               string `json:"fueleconomy"`
}

type EU struct {
	InUseComplianceNo string `json:"inusecomplianceno"`
	EUMainNo         string `json:"eu_mainno"`
	EUTypeCode       string `json:"eu_typecode"`
	TypeVariant      string `json:"typevariant"`
	TypeVersion      string `json:"typeversion"`
	EuronormNew      string `json:"euronormnew"`
	TekCode          string `json:"tekcode"`
	TekUndercode     string `json:"tekundercode"`
}

type Owner struct {
	OrganizationNumber string `json:"organizationnumber"`
	Born              string `json:"born"`
	Name              string `json:"name"`
	Address           string `json:"address"`
	Zipcode           string `json:"zipcode"`
	City              string `json:"city"`
	Municip           string `json:"municip,omitempty"`
}
