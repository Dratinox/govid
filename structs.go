package govid

// Country stores info about individual country
type Country struct {
	Country string
	Slug    string
	ISO2    string
}

// CasesInfo containts the Covid-19 cases numbers
type CasesInfo struct {
	NewConfirmed   int
	TotalConfirmed int
	NewDeaths      int
	TotalDeaths    int
	NewRecovered   int
	TotalRecovered int
}

// Summary holds data for Global cases numbers and data for each Country
type Summary struct {
	Global    CasesInfo
	Countries []SummaryCountry
}

// SummaryCountry holds data for Country cases and basic info
type SummaryCountry struct {
	Country     string
	CountryCode string
	Slug        string
	CasesInfo
}
