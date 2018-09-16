# Execute Command flogo activity
This activity allows your flogo application to execute commands. e.g. ls, ps or batch(.bat) files(on windows)


## Installation

```bash
flogo install github.com/pawarvishal123/executecmd
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
    {
      "name": "command",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting   | Description    |
|:----------|:---------------|
| command   | Input command - ps, ls, batch(.bat) file path |         
| result | The execution result of the command  |


## Configuration Examples
### Simple
Configure a task to execute a script:

```json
{
  "id": 3,
  "type": 1,
  "activityType": "executecmd",
  "name": "Execute Command",
  "attributes": [
    { "name": "command", "value": "ps" }
  ]
}
```
