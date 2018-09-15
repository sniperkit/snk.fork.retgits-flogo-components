/*
Sniperkit-Bot
- Status: analyzed
*/

// Package githubstarred implements activities to get GitHub starred repositories for a me/specific user.
package ghstarred

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

// Update these variables before testing to match your own GitHub account
const (
	githubToken = ""
)

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

// Test for a card at the top of the list
func TestEvalGetIssues(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	// Set required attributes
	tc.SetInput("token", githubToken)

	// tc.SetInput("username", githubUsername)
	// tc.SetInput("page", 1)
	// tc.SetInput("per_page", 100)
	// tc.SetInput("direction", githubListOptsDirection)
	// tc.SetInput("sort", githubListOptsSort)

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")
	fmt.Printf("The result is:\n[%v]\n", result)

	// The below statement can print the result as a JSON object
	//enc := json.NewEncoder(os.Stdout)
	//enc.Encode(result)
}
