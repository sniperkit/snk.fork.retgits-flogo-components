/*
Sniperkit-Bot
- Status: analyzed
*/

// Package starred implements activities to get GitHub starred repositories for a me/specific user.
package starred

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

var githubToken = os.Getenv("GITHUB_TOKEN")

// Update these variables before testing to match your own GitHub account
const (
	githubUsername      = "roscopecoltran"
	githubPage          = 1
	githubPerPage       = 100
	githubListRecursive = true
	githubListSort      = "updated"
	githubListDirection = "desc"
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

// IsJSON check if the string is valid JSON (note: uses json.Unmarshal).
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
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
	tc.SetInput("username", githubUsername)
	tc.SetInput("page", githubPage)
	tc.SetInput("per_page", githubPerPage)
	tc.SetInput("recursive", githubListRecursive)
	tc.SetInput("direction", githubListDirection)
	tc.SetInput("sort", githubListSort)

	// Execute the activity
	act.Eval(tc)

	// Check the result
	result := tc.GetOutput("result")

	// The below statement can print the result as a JSON object
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(result)

	// fmt.Printf("The result is:\n[%v]\n", result)

}
