package app

type Photo struct {
	PhotoID  int    `json:"photo_id"`
	TITLE    string `json:"title"`
	CAPTION  string `json:"caption"`
	PHOTOURL string `json:"photourl"`
}
