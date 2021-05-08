// Do not change this file, it has been generated using flogo-cli
// If you change it and rebuild the application your changes might get lost
package main

// embedded flogo app descriptor file
const flogoJSON string = `{
  "name": "TestApp",
  "type": "flogo:app",
  "version": "0.0.1",
  "appModel": "1.1.0",
  "description": "",
  "imports": [
    "github.com/project-flogo/contrib/activity/actreturn",
    "github.com/project-flogo/contrib/activity/log",
    "github.com/project-flogo/contrib/function/datetime",
    "github.com/project-flogo/contrib/function/string",
    "github.com/project-flogo/contrib/trigger/rest",
    "github.com/project-flogo/flow"
  ],
  "triggers": [
    {
      "id": "receive_http_message",
      "ref": "#rest",
      "name": "Receive HTTP Message",
      "description": "test/log_messages",
      "settings": {
        "port": 6677
      },
      "handlers": [
        {
          "settings": {
            "method": "GET",
            "path": "/test/log_level"
          },
          "action": {
            "ref": "#flow",
            "settings": {
              "flowURI": "res://flow:change_log_level"
            },
            "output": {
              "data": "=$.output_message"
            }
          }
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:change_log_level",
      "data": {
        "name": "ChangeLogLevel",
        "description": "Testing interactive way of changing LOG level",
        "metadata": {
          "input": [
            {
              "name": "input_test",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "output_message",
              "type": "string"
            },
            {
              "name": "http_code",
              "type": "integer"
            }
          ]
        },
        "tasks": [
          {
            "id": "log_2",
            "name": "Log A",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "Log message level: DEFAULT",
                "addDetails": true,
                "usePrint": true
              }
            }
          },
          {
            "id": "log_3",
            "name": "Log",
            "description": "Logs a message",
            "activity": {
              "ref": "#log",
              "input": {
                "message": "Log level: undefined",
                "addDetails": false,
                "usePrint": false
              }
            }
          },
          {
            "id": "actreturn_4",
            "name": "Return",
            "description": "Return Activity",
            "activity": {
              "ref": "#actreturn",
              "settings": {
                "mappings": {
                  "output_message": "=string.concat(\"Return message: hello world\", datetime.currentDatetime())",
                  "http_code": 200
                }
              }
            }
          }
        ],
        "links": [
          {
            "from": "log_2",
            "to": "log_3"
          },
          {
            "from": "log_3",
            "to": "actreturn_4"
          }
        ]
      }
    }
  ]
}`
const engineJSON string = ``

func init () {
	cfgJson = flogoJSON
	cfgEngine = engineJSON
}
