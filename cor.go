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
	preto    = "30"
	vermelho = "31"
	verde    = "32"
	amarelo  = "33"
	azul     = "34"
	magenta  = "35"
	ciano    = "36"
	branco   = "37"

	// pretoPastel    = "90"
	// vermelhoPastel = "91"
	// verdePastel    = "92"
	// amareloPastel  = "93"
	// azulPastel     = "94"
	// magentaPastel  = "95"
	// cianoPastel    = "96"
	// brancoPastel   = "97"
)

const (
	abertura   = "\033["
	fechamento = "m"
	reseta     = "\033[0m"
	// sep        = ";"
)

// const Nenhum = ""

func Pintar(texto, cor string) string {
	return abertura + cor + fechamento + texto + reseta
}

func ImprimirTodas() {
	for matiz, codigo := range cores {
		fmt.Println(Pintar(matiz, codigo))
	}
}

func Preto(texto string) string {
	return abertura + preto + fechamento + texto + reseta
}

func Vermelho(texto string) string {
	return abertura + vermelho + fechamento + texto + reseta
}

func Verde(texto string) string {
	return abertura + verde + fechamento + texto + reseta
}

func Amarelo(texto string) string {
	return abertura + amarelo + fechamento + texto + reseta
}

func Azul(texto string) string {
	return abertura + azul + fechamento + texto + reseta
}

func Magenta(texto string) string {
	return abertura + magenta + fechamento + texto + reseta
}

func Ciano(texto string) string {
	return abertura + ciano + fechamento + texto + reseta
}

func Branco(texto string) string {
	return abertura + branco + fechamento + texto + reseta
}
