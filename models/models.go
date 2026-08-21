// ipaddr-api/models
//
// This package contains all models used in the ipaddr service.

package models

// IPAddress is a struct we use to represent JSON and XML API responses.
type IPAddress struct {
	IP string `json:"ip" xml:"ip"`
}
