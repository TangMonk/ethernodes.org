package main

type Response struct {
	Data []struct {
		Client        string `json:"client"`
		ClientVersion string `json:"clientVersion"`
		Country       string `json:"country"`
		Host          string `json:"host"`
		ID            string `json:"id"`
		InSync        int64  `json:"inSync"`
		Isp           string `json:"isp"`
		LastUpdate    string `json:"lastUpdate"`
		Os            string `json:"os"`
		Port          int64  `json:"port"`
	} `json:"data"`
	Draw            int64 `json:"draw"`
	RecordsFiltered int64 `json:"recordsFiltered"`
	RecordsTotal    int64 `json:"recordsTotal"`
}
