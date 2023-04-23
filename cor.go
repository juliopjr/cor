package cor

import "fmt"

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

func Pintar(texto, cor string) string {
	return abertura + cor + fechamento + texto
}

func ImprimirTodas() {
	for matiz, codigo := range cores {
		fmt.Println(Pintar(matiz, codigo))
	}
}

func PPreto(texto string) string {
	return abertura + Preto + fechamento + texto
}

func PVermelho(texto string) string {
	return abertura + Vermelho + fechamento + texto
}

func PVerde(texto string) string {
	return abertura + Verde + fechamento + texto
}

func PAmarelo(texto string) string {
	return abertura + Amarelo + fechamento + texto
}

func PAzul(texto string) string {
	return abertura + Azul + fechamento + texto
}

func PMagenta(texto string) string {
	return abertura + Magenta + fechamento + texto
}

func PCiano(texto string) string {
	return abertura + Ciano + fechamento + texto
}

func PBranco(texto string) string {
	return abertura + Branco + fechamento + texto
}
