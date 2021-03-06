/*
Sniperkit-Bot
- Status: analyzed
*/

package starred

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var jsonMetadata = `{
    "name": "starred",
    "type": "flogo:activity",
    "ref": "github.com/sniperkit/snk.fork.retgits-flogo-components/activity/github/starred",
    "version": "0.0.1",
    "title": "GitHub Starred Repositories",
    "description": "Get the GitHub starred repositories for me/specific user.",
    "author": "sniperkit",
    "homepage": "https://github.com/sniperkit/snk.fork.retgits-flogo-components/tree/sniperkit/activity/github/starred",
    "inputs": [
        {
            "name": "token",
            "type": "string",
            "required": true
        },
        {
            "name": "username",
            "type": "string",
            "required": false
        },
        {
            "name": "recursive",
            "type": "boolean",
            "value": true,
            "required": false
        },
        {
            "name": "per_page",
            "type": "integer",
            "required": false,
            "value": 25,
            "minimum": 1,
            "maximum": 100,
            "allowed" : [10, 25, 50, 75, 100]
        },
        {
            "name": "page",
            "type": "integer",
            "required": false,
            "value": 1,
            "minimum": 1,
            "maximum": 2500,
            "exclusiveMaximum": true,
            "allowed" : ["non-zero"]
        },
        {
            "name": "sort",
            "type": "string",
            "required": false,
            "allowed": ["created", "updated", "pushed", "full_name"]
        },
        {
            "name": "direction",
            "type": "string",
            "required": false,
            "allowed": ["asc", "desc"]
        }
    ],
    "outputs": [
        {
            "name": "result",
            "type": "array"
        }
    ]
}
`

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(NewActivity(md))
}
