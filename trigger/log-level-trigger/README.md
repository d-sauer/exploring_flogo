# Test trigger

How to create trigger: https://tibcosoftware.github.io/flogo/labs/building-triggers/

Description for `metadata.json`

- name: The name of the activity (this should match the name of the folder the activity is in, like MyTimerTrigger)
- version: The version of the activity (it is recommended to use semantic versioning for your trigger)
- type: This describes to the Flogo engine what kind of contribution this is (this should be trigger in this case)
- ref: The Flogo engine is based on Go and this field is the “import” path for Go apps (generally speaking this should match your repository)
- description: A brief description of your activity (this is displayed in the Flogo Web UI)
- author: This is you!
- settings: An array of name/type pairs that describe global settings of the trigger (configuration that will be the same for every instance of this trigger used in your app)
- output: An array of name/type pairs that describe the output of the trigger (the data that gets sent to your flow)
- handler: An array of name/type pairs that describe flow specific settings of the trigger (configuration that will be unique for every instance of this trigger used in your app) 