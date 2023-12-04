package entities

type RequestMarketAll struct {
	InitDate string `json:"fec_pagoini"`
	LastDate string `json:"fec_pagofin"`
	Nemo     string `json:"nemo"`
}
