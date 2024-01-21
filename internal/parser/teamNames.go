package parser

import "strings"

func TeamNameOptionValidator(name string) string {
	name = strings.Replace(name, "State", "St.", 1)

	if name == "Miami" {
		return "Miami FL"
	}
	if name == "Ole Miss" {
		return "Mississippi"
	}
	if name == "Nebraska Cornhuskers" {
		return "Nebraska"
	}
	if name == "Pennsylvania" {
		return "Penn"
	}
	if name == "CSU Northridge" {
		return "Cal St. Northridge"
	}
	if name == "Long Island" {
		return "LIU"
	}
	if name == "California Baptist" {
		return "Cal Baptist"
	}
	if name == "UIC" {
		return "Illinois Chicago"
	}
	if name == "UL Monroe" {
		return "Louisiana Monroe"
	}
	if name == "Arkansas-Pine Bluff" {
		return "Arkansas Pine Bluff"
	}
	if name == "SE Louisiana" {
		return "Southeastern Louisiana"
	}
	if name == "NC St." {
		return "N.C. State"
	}
	if name == "UT Martin" {
		return "Tennessee Martin"
	}
	if name == "Florida International" {
		return "FIU"
	}
	if name == "McNeese" {
		return "McNeese St."
	}
	if name == "WV Mountaineers" {
		return "West Virginia"
	}
	if name == "Central Conn. St." {
		return "Central Connecticut"
	}
	if name == "Southern Methodist" {
		return "SMU"
	}
	if name == "Charlotte 49ers" {
		return "Charlotte"
	}
	if name == "St Josephs" {
		return "Saint Joseph's"
	}
	if name == "Citadel" {
		return "The Citadel"
	}
	if name == "Gardner-Webb" {
		return "Gardner Webb"
	}
	if name == "Nicholls" {
		return "Nicholls St."
	}
	if name == "CSU Bakersfield" {
		return "Cal St. Bakersfield"
	}
	if name == "San Jose St" {
		return "San Jose St."
	}
	if name == "CSU Fullerton" {
		return "Cal St. Fullerton"
	}
	if name == "UMass" {
		return "Massachusetts"
	}
	if name == "Nebraska Omaha Mavericks" {
		return "Nebraska Omaha"
	}
	if name == "Youngstown St" {
		return "Youngstown St."
	}
	if name == "GW Revolutionaries" {
		return "George Washington"
	}
	if name == "SE Missouri St." {
		return "Southeast Missouri St."
	}
	if name == "Mt. St. Mary's" {
		return "Mount St. Mary's"
	}
	if name == "Mt. St. Marys" {
		return "Mount St. Mary's"
	}
	if name == "Hawai'i" {
		return "Hawaii"
	}
	if name == "Stephen F Austin" {
		return "Stephen F. Austin"
	}
	if name == "UW Milwaukee" {
		return "Milwaukee"
	}
	if name == "Queen's University" {
		return "Queens"
	}
	if name == "GW Colonials" {
		return "George Washington"
	}
	if name == "Lafayette College" {
		return "Lafayette"
	}
	if name == "Louisiana Lafayette" {
		return "Louisiana"
	}
	if name == "Arkansas Little Rock" {
		return "Little Rock"
	}
	if name == "UW Green Bay" {
		return "Green Bay"
	}
	if name == "Utah Valley St." {
		return "Utah Valley"
	}
	if name == "San Diego Tritons" {
		return "UC San Diego"
	}
	return name
}
