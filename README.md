Это сервис order в составе приложения eshop

30.03.2024

В программе два функционала:
1. Чтение переменных окружения с помощью "github.com/kelseyhightower/envconfig"
экспорт тестовых переменных - ниже.
2. Пример структуры файлов для посторения тестового обработчика времени с функцией вычисления остатка часов до 1го апреля 2024 года.
Обработчик вызывается через интерфейс в endpoint.
Все объекты создаются через свои конструкторы


---------------------------------------------------
Экспорт тестовых переменных:

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

