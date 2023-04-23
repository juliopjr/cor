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

func pPreto(texto string) string {
	return abertura + Preto + fechamento + texto
}

func pVermelho(texto string) string {
	return abertura + Vermelho + fechamento + texto
}

func pVerde(texto string) string {
	return abertura + Verde + fechamento + texto
}

func pAmarelo(texto string) string {
	return abertura + Amarelo + fechamento + texto
}

func pAzul(texto string) string {
	return abertura + Azul + fechamento + texto
}

func pMagenta(texto string) string {
	return abertura + Magenta + fechamento + texto
}

func pCiano(texto string) string {
	return abertura + Ciano + fechamento + texto
}

func pBranco(texto string) string {
	return abertura + Branco + fechamento + texto
}
