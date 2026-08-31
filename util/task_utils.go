package util

import (
	"manta/models"
)

func ExtractGroupInformation(task models.Task) (string, string) {
	groupName := "Ungrouped"	
	groupColor := "#ff0000"

	value, exists := groups[t.Group]
	
	if(exists == true) {
		myGroup := value
		
		groupName = myGroup.Name
		groupColor = myGroup.Color
	}

	return groupName, groupColor
}