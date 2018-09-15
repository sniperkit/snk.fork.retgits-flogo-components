# GitHub Issues

Get the GitHub issues assigned to an authenticated user

## Installation

```bash
flogo install github.com/sniperkit/snk.fork.retgits-flogo-components/activity/githubissues
```
Link for flogo web:
```
https://github.com/sniperkit/snk.fork.retgits-flogo-components/activity/githubissues
```

## Schema
Inputs and Outputs:

```json
{
"inputs": [
        {
            "name": "token",
            "type": "string",
            "required": true
        },
        {
            "name": "timeInterval",
            "type": "integer",
            "required": true
        }
    ],
    "outputs": [
        {
            "name": "result",
            "type": "any"
        }
    ]
}
```
## Inputs
| Input        | Description                                                                                                                         |
|:-------------|:------------------------------------------------------------------------------------------------------------------------------------|
| token        | Your Personal Access Token from GitHub                                                                                              |
| timeInterval | The timeinterval in minutes to check GitHub issues (setting this to `120` will get issues assigned to the user in the past 2 hours) |

## Ouputs
| Output      | Description                                                                                                                                                     |
|:------------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| result      | An array of issues assigned to the user in the past x minutes. The data structure for the response can be found [here](https://developer.github.com/v3/issues/) |