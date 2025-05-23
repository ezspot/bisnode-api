package models

// DirectorySearchRequest represents the request body for directory search
type DirectorySearchRequest struct {
	Form struct {
		Type         string `json:"Type"`
		SearchString string `json:"Searchstring"`
	} `json:"Form"`
	Options struct {
		SearchMode       int  `json:"SearchMode"`
		OnlyFoundWords   bool `json:"OnlyFoundWords"`
		ListingType      int  `json:"ListingType"`
		ResultLimit      int  `json:"ResultLimitSearch"`
	} `json:"Options"`
}

// DirectorySearchResponse represents the response from directory search
type DirectorySearchResponse struct {
	Result  []DirectoryResult `json:"Result"`
	Service struct {
		Dataset       string `json:"dataset"`
		Documentation string `json:"documentation"`
		Version      string `json:"version"`
		Timestamp    string `json:"timestamp"`
		Message      string `json:"message"`
	} `json:"Service"`
}

// DirectoryResult represents a single result in the directory search
type DirectoryResult struct {
	Type              string `json:"type"`
	OrganizationNumber string `json:"organizationnumber,omitempty"`
	Born              string `json:"born,omitempty"`
	Dead              string `json:"dead,omitempty"`
	Age               struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	} `json:"Age,omitempty"`
	Gender     string `json:"gender,omitempty"`
	FirstName  string `json:"firstname,omitempty"`
	MiddleName string `json:"middlename,omitempty"`
	LastName   string `json:"lastname,omitempty"`
	StreetName string `json:"streetname,omitempty"`
	HouseNo    string `json:"houseno,omitempty"`
	Entrance   string `json:"entrance,omitempty"`
	ZipCode    string `json:"zipcode,omitempty"`
	City       string `json:"city,omitempty"`
	Longitude  float64 `json:"longitude"`
	Latitude   float64 `json:"latitude"`
	Telephone  string `json:"telephone,omitempty"`
	Mobile     string `json:"mobile,omitempty"`
	Reservation struct {
		DirectMail    bool `json:"directmail"`
		Telemarketing bool `json:"telemarketing"`
		Humanitarian  bool `json:"humanitarian"`
	} `json:"Reservation,omitempty"`
	Addresses []struct {
		Source      string `json:"source"`
		Type        string `json:"type"`
		Quality     string `json:"quality"`
		StreetName  string `json:"streetname"`
		HouseNo     string `json:"houseno"`
		Entrance    string `json:"entrance"`
		ZipCode     string `json:"zipcode"`
		City        string `json:"city"`
		Longitude   float64 `json:"longitude"`
		Latitude    float64 `json:"latitude"`
		Date        struct {
			FirstAcquired     string `json:"firstaquired"`
			LastAcquired      string `json:"lastaquired"`
			InformationChanged string `json:"informationchanged"`
		} `json:"Date"`
	} `json:"Address,omitempty"`
	Phones []struct {
		Source      string `json:"source"`
		Type        string `json:"type"`
		Quality     string `json:"quality"`
		Number      string `json:"number"`
		Date        struct {
			FirstAcquired     string `json:"firstaquired"`
			LastAcquired      string `json:"lastaquired"`
			InformationChanged string `json:"informationchanged"`
		} `json:"Date"`
	} `json:"Phone,omitempty"`
}

// SearchRequest represents a search request for directory search
type SearchRequest struct {
	// MobileNumber is used to search for a person by mobile number
	MobileNumber string `json:"mobileNumber,omitempty"`
	// OrganizationNumber is used to search for a company by organization number
	OrganizationNumber string `json:"organizationNumber,omitempty"`
}
