package cor

// := "\033[<intensidadeCor>;<aparenciaTexto>;<corTexto>;<corFundo>m"
// fmt.Printf("\033[31mHello, World!\033[0m\n")

var cores = map[string]string{
	"Preto":    "30",
	"Vermelho": "31",
	"Verde":    "32",
	"Amarelo":  "33",
	"Azul":     "34",
	"Magenta":  "35",
	"Ciano":    "36",
	"Branco":   "37",

	"PretoPastel":    "90",
	"VermelhoPastel": "91",
	"VerdePastel":    "92",
	"AmareloPastel":  "93",
	"AzulPastel":     "94",
	"MagentaPastel":  "95",
	"CianoPastel":    "96",
	"BrancoPastel":   "97",
}

const (
	Preto    = "30"
	Vermelho = "31"
	Verde    = "32"
	Amarelo  = "33"
	Azul     = "34"
	Magenta  = "35"
	Ciano    = "36"
	Branco   = "37"

	PretoPastel    = "90"
	VermelhoPastel = "91"
	VerdePastel    = "92"
	AmareloPastel  = "93"
	AzulPastel     = "94"
	MagentaPastel  = "95"
	CianoPastel    = "96"
	BrancoPastel   = "97"
)

const (
	abertura   = "\033["
	fechamento = "m"
	reseta     = "\033[0m"
	// sep        = ";"
)

// const Nenhum = ""
