package models

type DBConfig struct {
	Host         string
	Port         int
	Database     string
	User         string
	Password     string
	MaxLifeTime  int
	MaxIdleConns int
	MaxOpenConn  int
	TimeOut      int
}
