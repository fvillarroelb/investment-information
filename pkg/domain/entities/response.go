package entities

type ListaResult struct {
	MERCADO    string `json:"MERCADO"`
	MONTO      int    `json:"MONTO"`
	N_NEGOCIOS int    `json:"n_NEGOCIOS"`
}

type ListaResultArray struct {
	ListaResult []ListaResult `json:"listaResult"`
}
