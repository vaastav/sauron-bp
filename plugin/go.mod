module github.com/vaastav/sauron-bp/plugin

go 1.21

require (
	github.com/blueprint-uservices/blueprint/blueprint v0.0.0-20240619221802-d064c5861c1e
	github.com/blueprint-uservices/blueprint/plugins v0.0.0-20240619221802-d064c5861c1e
	github.com/vaastav/sauron-bp/sauron_runtime v0.0.0
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56
	honnef.co/go/tools v0.4.7
)

require (
	github.com/blueprint-uservices/blueprint/runtime v0.0.0-20240619221802-d064c5861c1e // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/otiai10/copy v1.14.0 // indirect
	go.mongodb.org/mongo-driver v1.15.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.26.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.26.0 // indirect
	go.opentelemetry.io/otel/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/sdk v1.26.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.26.0 // indirect
	golang.org/x/exp/typeparams v0.0.0-20221208152030-732eee02a75a // indirect
	golang.org/x/mod v0.19.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/tools v0.23.0 // indirect
)

replace github.com/vaastav/sauron-bp/sauron_runtime => ../sauron_runtime
