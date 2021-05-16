module main

go 1.16

replace github.com/d-sauer/exploring_flogo/trigger/log-level-trigger => /Users/davorsauer/development/repo/d-sauer/exploring-flogo/trigger/log-level-trigger

require (
	github.com/d-sauer/exploring_flogo/trigger/log-level-trigger v0.0.0-20210511211036-5f8a91960d2c
	github.com/project-flogo/contrib/activity/log v0.10.0
	github.com/project-flogo/core v1.4.0
	github.com/project-flogo/flow v1.4.0
)
