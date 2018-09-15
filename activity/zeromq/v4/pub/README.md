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
      "name": "service",
      "type": "string",
      "required": true,
      "value": "tcp://*:14444"
    },
    {
      "name": "uri",
      "type": "string",
      "required": true,
      "value": "tcp:localhost:5555"
    },
    {
      "name": "topic",
      "type": "string",
      "required": true,
      "value": "zmq.REP",
    },
    {
      "name": "message",
      "type": "string",
      "required": true,
      "value": "testing zmq server for flogo"
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
| service     | service                                                                    |
| uri         | uri                                                                        |
| topic       | Topic name                                                                 |
| message     | Message Content                                                            |

## Ouputs
| Output      | Description                                                                |
|:------------|:---------------------------------------------------------------------------|
| result      |                                                                            |