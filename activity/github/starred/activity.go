/*
Sniperkit-Bot
- Status: analyzed
*/

// Package ghstarred implements activities to get GitHub starred repositories for a me/specific user.
package ghstarred

// Imports
import (
	ctx "context"
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Refs:
// - https://developer.github.com/v3/#pagination
// - https://developer.github.com/v3/activity/starring/#list-repositories-being-starred

// Last Page, check how to validate the page offset first.
// Link: <https://api.github.com/user/repos?page=3&per_page=100>; rel="next",
//  <https://api.github.com/user/repos?page=50&per_page=100>; rel="last"

// Constants
const (
	ivGithubAccessToken       = "token"
	ivGithubUsername          = "username"
	ivGithubListOptsPage      = "page"      // default: 1, max: LinkHeader.LastPage
	ivGithubListOptsPerPage   = "per_page"  // default: 30, max: 100
	ivGithubListOptsSort      = "sort"      // created, updated, pushed, full_name. Default is "full_name".
	ivGithubListOptsDirection = "direction" // Default is "asc" when sort is "full_name", otherwise default is "desc".
	ovResult                  = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-github-starred")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the data
	githubToken := context.GetInput(ivGithubAccessToken).(string)
	githubUsername := context.GetInput(ivGithubUsername).(string)
	githubPage := context.GetInput(ivGithubListOptsPage).(int)
	githubPerPage := context.GetInput(ivGithubListOptsPerPage).(int)
	githubListSort := context.GetInput(ivGithubListOptsSort).(string)
	githubListDirection := context.GetInput(ivGithubListOptsDirection).(string)

	// Create a new GitHub client
	ctxt := ctx.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: githubToken,
		},
	)
	tc := oauth2.NewClient(ctxt, ts)
	client := github.NewClient(tc)

	log.Infof("Check GitHub starred repositories for the user [%s], page=%v, per_page=%v", githubUsername, githubPage, githubPerPage)

	// Get all the starred repositories for me/specific user
	starsOpts := &github.ActivityListStarredOptions{
		ListOptions: github.ListOptions{
			Page:    githubPage,
			PerPage: githubPerPage,
		},
	}

	if githubListSort != "" {
		starsOpts.Sort = githubListSort
	}

	if githubListDirection != "" {
		starsOpts.Direction = githubListDirection
	}

	stars, _, err := client.Activity.ListStarred(ctxt, false, starsOpts)
	if err != nil {
		log.Error(err.Error())
		return true, err
	}

	log.Infof("GitHub returned %v starred repositories for the page %v/%v", len(stars), githubPage, githubPerPage)
	datamap := make([]interface{}, len(stars))

	for idx, star := range stars {
		datamap[idx] = star
	}

	// Set the output value in the context
	context.SetOutput(ovResult, datamap)

	return true, nil
}
