module github.com/wuapp/wuapp

go 1.15

require (
	github.com/wuapp/rj v0.0.0-20201015064226-b86c7f42f931
	github.com/wuapp/util v0.0.0-20201019074138-16240c3fd303
	github.com/wuapp/log v0.0.0
)

replace (
    github.com/wuapp/log => ../log
)
