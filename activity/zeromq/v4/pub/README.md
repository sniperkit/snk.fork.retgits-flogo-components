# ZeroMQ v4 - Publish

Publish zmq message

## Installation

```bash
flogo install github.com/sniperkit/snk.fork.retgits-flogo-components/activity/zeromq/v4/pub
```

Link for flogo web:
```
https://github.com/sniperkit/snk.fork.retgits-flogo-components/activity/zeromq/v4/pub
```

## Schema
Inputs and Outputs:

```json
{
  "input":[
    {
      "name": "URI",
      "type": "string",
      "required": true
    },
    {
      "name": "Topic",
      "type": "string",
      "required": true
    },
    {
      "name": "Message",
      "type": "string",
      "required": true
    }
  ],
  "output": [
    {
      "name": "output",
      "type": "any"
    }
  ]
}
```
## Inputs
| Input       | Description                                                                |
|:------------|:---------------------------------------------------------------------------|
| URI         | Server uri.                                                                |
| Topic       | Topic name                                                                 |
| Message     | Message Content                                                            |

## Ouputs
| Output      | Description                                                                |
|:------------|:---------------------------------------------------------------------------|
| result      |                                                                            |