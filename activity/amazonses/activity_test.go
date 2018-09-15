/*
Sniperkit-Bot
- Status: analyzed
*/

// Package amazonses allows the user to send emails using Amazon Simple Email Service (SES)
package amazonses

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

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

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("to", "someone@somewhere.com")
	tc.SetInput("from", "me@somewhere.com")
	tc.SetInput("content", "Hello World!")
	tc.SetInput("subject", "Just saying hello...")
	tc.SetInput("awsAccessKeyID", "")
	tc.SetInput("awsSecretAccessKey", "")
	tc.SetInput("awsRegion", "")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	fmt.Printf("[%s]\n", result)
}
