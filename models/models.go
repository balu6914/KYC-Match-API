package models

import "time"

// KYCRequest represents the incoming request body
type KYCRequest struct {
	PhoneNumber          string `json:"phoneNumber"`
	IDDocument           string `json:"idDocument"`
	Name                 string `json:"name"`
	GivenName            string `json:"givenName"`
	FamilyName           string `json:"familyName"`
	NameKanaHankaku      string `json:"nameKanaHankaku"`
	NameKanaZenkaku      string `json:"nameKanaZenkaku"`
	MiddleNames          string `json:"middleNames"`
	FamilyNameAtBirth    string `json:"familyNameAtBirth"`
	Address              string `json:"address"`
	StreetName           string `json:"streetName"`
	StreetNumber         string `json:"streetNumber"`
	PostalCode           string `json:"postalCode"`
	Region               string `json:"region"`
	Locality             string `json:"locality"`
	Country              string `json:"country"`
	HouseNumberExtension string `json:"houseNumberExtension"`
	Birthdate            string `json:"birthdate"`
	Email                string `json:"email"`
	Gender               string `json:"gender"`
}

// MatchResult represents the result of a comparison
type MatchResult struct {
	Value  string `json:"value"`
	Score  int    `json:"score,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// KYCResponse represents the response body
type KYCResponse struct {
	IDDocumentMatch           MatchResult `json:"idDocumentMatch"`
	NameMatch                 MatchResult `json:"nameMatch"`
	GivenNameMatch            MatchResult `json:"givenNameMatch"`
	FamilyNameMatch           MatchResult `json:"familyNameMatch"`
	NameKanaHankakuMatch      MatchResult `json:"nameKanaHankakuMatch"`
	NameKanaZenkakuMatch      MatchResult `json:"nameKanaZenkakuMatch"`
	MiddleNamesMatch          MatchResult `json:"middleNamesMatch"`
	FamilyNameAtBirthMatch    MatchResult `json:"familyNameAtBirthMatch"`
	AddressMatch              MatchResult `json:"addressMatch"`
	StreetNameMatch           MatchResult `json:"streetNameMatch"`
	StreetNumberMatch         MatchResult `json:"streetNumberMatch"`
	PostalCodeMatch           MatchResult `json:"postalCodeMatch"`
	RegionMatch               MatchResult `json:"regionMatch"`
	LocalityMatch             MatchResult `json:"localityMatch"`
	CountryMatch              MatchResult `json:"countryMatch"`
	HouseNumberExtensionMatch MatchResult `json:"houseNumberExtensionMatch"`
	BirthdateMatch            MatchResult `json:"birthdateMatch"`
	EmailMatch                MatchResult `json:"emailMatch"`
	GenderMatch               MatchResult `json:"genderMatch"`
}

// Customer represents the entity stored in HarperDB
type Customer struct {
	PhoneNumber          string    `json:"phoneNumber"`
	IDDocument           string    `json:"idDocument"`
	Name                 string    `json:"name"`
	GivenName            string    `json:"givenName"`
	FamilyName           string    `json:"familyName"`
	NameKanaHankaku      string    `json:"nameKanaHankaku"`
	NameKanaZenkaku      string    `json:"nameKanaZenkaku"`
	MiddleNames          string    `json:"middleNames"`
	FamilyNameAtBirth    string    `json:"familyNameAtBirth"`
	Address              string    `json:"address"`
	StreetName           string    `json:"streetName"`
	StreetNumber         string    `json:"streetNumber"`
	PostalCode           string    `json:"postalCode"`
	Region               string    `json:"region"`
	Locality             string    `json:"locality"`
	Country              string    `json:"country"`
	HouseNumberExtension string    `json:"houseNumberExtension"`
	Birthdate            time.Time `json:"birthdate"`
	Email                string    `json:"email"`
	Gender               string    `json:"gender"`
}
