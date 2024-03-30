this is a readmefile

export MYAPP_DEBUG=false
export MYAPP_PORT=8080
export MYAPP_USER=Kelsey
export MYAPP_RATE="0.5"
export MYAPP_TIMEOUT="3m"
export MYAPP_USERS="rob,ken,robert"
export MYAPP_COLORCODES="red:1,green:2,blue:3"

/order
├── api
├── bin
│   └── golangci-lint
├── cmd
│   └── order
│       └── main.go
├── configs
│   ├── config.go
│   └── config_test.go
├── deploy
├── go.mod
├── go.sum
├── internal
│   ├── endpoint
│   │   └── endpoint.go
│   ├── pkg
│   │   └── app
│   │       └── app.go
│   └── service
│       ├── readconfig
│       │   └── readconfig.go
│       └── timecalc
│           └── time-calc.go
├── Makefile
├── pkg
└── README.md

