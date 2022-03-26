package autoMigrate

import "os"

func IsProd() bool {
	prod := os.Getenv("HDUHELP_PROD")
	return prod == "HDUHELP"
}
