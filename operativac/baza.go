package operativac

type Dopisnik struct {
	sanduce chan Message
}

type Dopis struct {
	adresaPrimaoca string
	poruka         any
}

type Baza struct {
	adresa   string
	dopisnik Dopisnik
}
