package binance

type API struct {
	ApiKey    string `json:"apiKey"`
	ApiSecret string `json:"apiSecret"`
}

type Configs struct {
	Id     uint `json:"id"`
	Enable bool `json:"enable"`

	Api     *API   `json:"api,omitempty"`
	Message string `json:"message,omitempty"`
}
