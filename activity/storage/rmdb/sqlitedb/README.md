# SQLite DB Activity
This activity allows you to perform create table, insert, update and delete oeprations on SQLite DB.
To build this activity on windows, gcc executable is required in PATH. MingW can be used.

## Installation

### Flogo CLI
```bash
flogo install github.com/sniperkit/snk.fork.retgits-flogo-components/activity/storage/rmdb/sqlitedb
```

## Schema
Inputs and Outputs:

```json
{
  "input":[
    {
      "name": "database",
      "type": "string",
      "required": true
    },
    {
      "name": "query",
      "type": "string",
      "required": true
    },
    {
      "name": "parameters",
      "type": "params"
    },
    {
      "name": "prefix_path",
      "type": "string",
      "required": false,
      "value": "~"
    }
  ],
  "output": [
    {
      "name": "result",
      "type": "any"
    }
  ]
}
```

## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| database  | True     | The name of SQLite DB file |
| query       | True     | SQL statement |
| parameters     | False     | Parameters used in SQL statement (Please refer below sample input)|
| prefix_path     | False     | Prefix Path for local database|

## Example Inputs
```json
{
            "id": "sqlitedb_3",
            "name": "SQLite DB Operations (2)",
            "description": "Create, insert, update & delete from SQLite DB",
            "activity": {
              "ref": "github.com/sniperkit/snk.fork.retgits-flogo-components/activity/storage/rmdb/sqlite3",
              "input": {
                "database": "test2",
                "query": "insert into emp(id, name) values(?id, ?name)",
                "parameters": {
                  "id": "11",
                  "name": "SNK"
                }
              }
            }
          }
```

## Third Party Library
SQLite API in Go - [https://github.com/mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
