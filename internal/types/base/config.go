package base

type Config struct {
	RegisterTypes []string `json:"register_types"`
	LoginTypes    []string `json:"login_types"`
	LoginGuests   []Guest  `json:"login_guests"`
}

type Guest struct {
	Type     string `json:"type"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
