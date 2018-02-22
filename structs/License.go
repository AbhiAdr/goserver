package structs

type LicenseJson struct {
	Demo   string `json:"demo"`
	Expiry string `json:"exp"`
}

type Licensedata struct {
	Data []string
}
