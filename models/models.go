// how data is presented in table
// golang does not have direct support with json - so need to work with encoding and decoding json
package models

// GOLANG vars - StockID, Name.. (caps)
// In JSON - informing Golang, will look like "stockid"
type Stock struct {
	StockID int64  `json:"stockid"`
	Name    string `json:"name"`
	Price   int64  `json:"price"`
	Company string `json:"company"`
}
