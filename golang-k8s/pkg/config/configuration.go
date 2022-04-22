package config

type Configuration struct {
	AppSettings AppSettings
	Database    Database
	CategoryApi ApiSettings
}

type AppSettings struct {
	HealthCheck string
}

type Database struct {
	Address      string
	Replicaset   string
	DatabaseName string
}

type ApiSettings struct {
	Url        string
}
