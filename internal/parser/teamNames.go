package parser

import "strings"

func TeamNameOptionValidator(name string) string {
	name = strings.Replace(name, "State", "St.", 1)
	name = strings.Replace(name, "&amp;", "&", -1)
	switch name {
	case "Miami":
		return "Miami FL"
	case "Miami (FL)":
		return "Miami FL"
	case "Ole Miss":
		return "Mississippi"
	case "Nebraska Cornhuskers":
		return "Nebraska"
	case "Pennsylvania":
		return "Penn"
	case "CSU Northridge":
		return "Cal St. Northridge"
	case "Long Island":
		return "LIU"
	case "Mercyhurst Lakers":
		return "Mercyhurst"
	case "McNeese St.":
		return "McNeese"
	case "Texas A&M–Commerce":
		return "Texas A&M Commerce"
	case "Texas A&M-Corpus Christi":
		return "Texas A&M Corpus Chris"
	case "Bethune-Cookman":
		return "Bethune Cookman"
	case "California Baptist":
		return "Cal Baptist"
	case "Mississippi Valley":
		return "Mississippi Valley St."
	case "UIC":
		return "Illinois Chicago"
	case "UL Monroe":
		return "Louisiana Monroe"
	case "Arkansas-Pine Bluff":
		return "Arkansas Pine Bluff"
	case "SE Louisiana":
		return "Southeastern Louisiana"
	case "UConn":
		return "Connecticut"
	case "College of Charleston":
		return "Charleston"
	case "NC St.":
		return "N.C. State"
	case "UT Martin":
		return "Tennessee Martin"
	case "Florida International":
		return "FIU"
	case "McNeese":
		return "McNeese St."
	case "WV Mountaineers":
		return "West Virginia"
	case "Central Conn. St.":
		return "Central Connecticut"
	case "Southern Methodist":
		return "SMU"
	case "Charlotte 49ers":
		return "Charlotte"
	case "St Josephs":
		return "Saint Joseph's"
	case "Citadel":
		return "The Citadel"
	case "Gardner-Webb":
		return "Gardner Webb"
	case "Nicholls":
		return "Nicholls St."
	case "CSU Bakersfield":
		return "Cal St. Bakersfield"
	case "San Jose St":
		return "San Jose St."
	case "CSU Fullerton":
		return "Cal St. Fullerton"
	case "UMass":
		return "Massachusetts"
	case "Nebraska Omaha Mavericks":
		return "Nebraska Omaha"
	case "Youngstown St":
		return "Youngstown St."
	case "GW Revolutionaries":
		return "George Washington"
	case "SE Missouri St.":
		return "Southeast Missouri St."
	case "Mt. St. Mary's":
		return "Mount St. Mary's"
	case "Mt. St. Marys":
		return "Mount St. Mary's"
	case "Hawai'i":
		return "Hawaii"
	case "Stephen F Austin":
		return "Stephen F. Austin"
	case "UW Milwaukee":
		return "Milwaukee"
	case "Queen's University":
		return "Queens"
	case "GW Colonials":
		return "George Washington"
	case "Lafayette College":
		return "Lafayette"
	case "Louisiana Lafayette":
		return "Louisiana"
	case "Arkansas Little Rock":
		return "Little Rock"
	case "UW Green Bay":
		return "Green Bay"
	case "Utah Valley St.":
		return "Utah Valley"
	case "San Diego Tritons":
		return "UC San Diego"
	case "Nebraska-Omaha":
		return "Nebraska Omaha"
	}

	return name
}
