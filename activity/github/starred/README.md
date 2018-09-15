# GitHub Starred

Get the GitHub starred repositories for me/specific user.

## Installation

```bash
flogo install github.com/sniperkit/snk.fork.retgits-flogo-components/activity/github/starred
```
Link for flogo web:
```
https://github.com/sniperkit/snk.fork.retgits-flogo-components/activity/github/starred
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
            "name": "username",
            "type": "string",
            "required": false
        },
        {
            "name": "per_page",
            "type": "integer",
            "required": false
        },
        {
            "name": "page",
            "type": "integer",
            "required": false
        },
        {
            "name": "sort",
            "type": "string",
            "required": false
        },
        {
            "name": "direction",
            "type": "string",
            "required": false
        }
    ],
    "outputs": [
        {
            "name": "result",
            "type": "array"
        }
    ]
}
```
## Inputs
| Input        | Description                                                                                                                         |
|:-------------|:------------------------------------------------------------------------------------------------------------------------------------|
| token        | Your Personal Access Token from GitHub                                                                                              |

## Ouputs
| Output      | Description                                                                                                                                                     |
|:------------|:----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| result      | An array of starred assigned to the user in the past x minutes. The data structure for the response can be found [here](https://developer.github.com/v3/issues/) |