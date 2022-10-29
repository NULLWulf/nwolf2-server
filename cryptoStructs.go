package main

// CmpResponse Response body of CMP api, only interested in data array of crypto objects
type CmpResponse struct {
	TimeBlockUTC string
	Partition    string
	Data         []CryptoElement `dynamodbav:"data"`
}

// CryptoElement data type containing nominal and statistical data
type CryptoElement struct {
	Name        string `dynamodbav:"name"`
	Symbol      string `dynamodbav:"symbol"`
	CmcRank     int    `dynamodbav:"cmc_rank"`
	CryptoQuote Quote  `dynamodbav:"quote"`
}

// Quote containing data relative data to respective queried currency in this case USD
type Quote struct {
	USDStats USDRelativeData `dynamodbav:"USD"`
}

// USDRelativeData United States Dollar Relative Data
type USDRelativeData struct {
	Price             float64 `dynamodbav:"price"`
	Volume24hr        float64 `dynamodbav:"volume_24h"`
	VolumeChange24hr  float64 `dynamodbav:"volume_change_24h"`
	PercentChange1hr  float64 `dynamodbav:"percent_change_1h"`
	PercentChange24hr float64 `dynamodbav:"percent_change_24h"`
	PercentChange7d   float64 `dynamodbav:"percent_change_7d"`
	PercentChange30d  float64 `dynamodbav:"percent_change_30d"`
	PercentChange60d  float64 `dynamodbav:"percent_change_60d"`
	PercentChange90d  float64 `dynamodbav:"percent_change_90d"`
}
