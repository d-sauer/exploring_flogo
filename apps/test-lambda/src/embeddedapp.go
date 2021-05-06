// Do not change this file, it has been generated using flogo-cli
// If you change it and rebuild the application your changes might get lost
package main

// embedded flogo app descriptor file
const flogoJSON string = `{
  "name": "Tutorial",
  "type": "flogo:app",
  "version": "0.0.1",
  "description": "",
  "appModel": "1.0.0",
  "imports": [
    "github.com/project-flogo/legacybridge"
  ],
  "triggers": [
    {
      "id": "start_flow_as_a_function_in_lambda",
      "ref": "github.com/TIBCOSoftware/flogo-contrib/trigger/lambda",
      "settings": {},
      "handlers": [
        {
          "settings": null,
          "actions": [
            {
              "ref": "github.com/TIBCOSoftware/flogo-contrib/action/flow"
            }
          ]
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:lambda_flow",
      "data": {
        "name": "LambdaFlow",
        "metadata": {
          "input": [
            {
              "name": "name",
              "type": "string"
            }
          ],
          "output": [
            {
              "name": "greeting",
              "type": "any"
            }
          ]
        },
        "tasks": [
          {
            "id": "log_2",
            "name": "Log Message",
            "description": "Simple Log Activity",
            "activity": {
              "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/log",
              "input": {
                "message": "",
                "flowInfo": "false",
                "addToFlow": "false"
              },
              "mappings": {
                "input": [
                  {
                    "type": "expression",
                    "value": "string.concat(\"Hello \", $flow.name)",
                    "mapTo": "message"
                  }
                ]
              }
            }
          },
          {
            "id": "actreturn_3",
            "name": "Return",
            "description": "Simple Return Activity",
            "activity": {
              "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/actreturn",
              "input": {
                "mappings": [
                  {
                    "mapTo": "greeting",
                    "type": "object",
                    "value": {
                      "Hello": "{{$flow.name}}"
                    }
                  }
                ]
              }
            }
          }
        ],
        "links": [
          {
            "from": "log_2",
            "to": "actreturn_3"
          }
        ]
      }
    }
  ]
}`

func init () {
	cfgJson = flogoJSON
}
