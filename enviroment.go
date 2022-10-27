package config

import "os"

// Authentication
var SECRET = "my-secret"

// Database
// var DB_DRIVER = "mysql"
// var DB_USER = "bde51da67a0731"
// var DB_PASS = "1957c08f"
// var DB_PROTOCOL = "tcp"
// var DB_HOST = "us-cdbr-east-06.cleardb.net"
// var DB_PORT = "3306"
// var DB_NAME = "heroku_7edb9b73b104c5c"
// var DB_USERS_TABLE = "users"
// var DB_URL = `dbDriver, dbUser+":"+dbPass+"@"+dbProtocol+"("+dbHost+":"+dbPort+")/"+dbName`
var DB_DRIVER = os.Getenv("DB_DRIVER")
var DB_USER = os.Getenv("DB_USER")
var DB_PASS = os.Getenv("DB_PASS")
var DB_PROTOCOL = os.Getenv("DB_PROTOCOL")
var DB_HOST = os.Getenv("DB_HOST")
var DB_PORT = os.Getenv("DB_PORT")
var DB_NAME = os.Getenv("DB_NAME")
var DB_USERS_TABLE = os.Getenv("DB_USERS_TABLE")
var DB_URL = os.Getenv("DB_URL")

// Google Auth
var (
	// GOOGLE_CLIENT_ID     = "104234061712-n83vudbke4h53vu17tr3khorber0qg3i.apps.googleusercontent.com"
	// GOOGLE_CLIENT_SECRET = "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"
	// GOOGLE_REDIRECT_URL  = SERVER_PROTOCOL + "://" + SERVER_HOST + ":" + SERVER_PORT + "/google/callback"
	GOOGLE_CLIENT_ID     = os.Getenv("GOOGLE_CLIENT_ID")
	GOOGLE_CLIENT_SECRET = os.Getenv("GOOGLE_CLIENT_SECRET")
	GOOGLE_REDIRECT_URL  = os.Getenv("GOOGLE_REDIRECT_URL")
)

// Server
var (
	// SERVER_PROTOCOL = "http"
	// SERVER_HOST     = "localhost"
	// SERVER_PORT     = "8080"
	SERVER_PROTOCOL = os.Getenv("SERVER_PROTOCOL")
	SERVER_HOST     = os.Getenv("SERVER_HOST")
	SERVER_PORT     = os.Getenv("SERVER_PORT")
)
