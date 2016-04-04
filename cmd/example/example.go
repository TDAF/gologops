package main

import (
	"errors"
	"fmt"

	"github.com/TDAF/gologops"
)

type complexErr struct {
	Text  string
	Cause *complexErr
}

func (ce complexErr) Error() string { return "uno complejito" }

type NotJSONableNError struct{}

func (NotJSONableNError) Error() string { return "a very strange error" }
func (NotJSONableNError) MarshalJSON() ([]byte, error) {
	return nil, errors.New("JSON not supported")
}

func main() {
	l := gologops.NewLogger()
	l.SetFlags(gologops.Lmethod)
	l.SetFlags(gologops.Llongfile)
	l.SetFlags(gologops.Lshortfile)

	l.Infof("%d y %d son %d", 2, 2, 4)
	l.Info("y ocho dieciséis")
	l.SetContext(gologops.C{"prefix": "prefijo"})
	l.InfoC(gologops.C{"local": "España y olé"}, "%d y %d son %d", 2, 2, 4)
	l.InfoC(gologops.C{"local": `{"json":"pompón"}`}, "y ocho dieciséis")

	ce := complexErr{"1", &complexErr{"2", &complexErr{"3", nil}}}
	l.ErrorE(ce, nil, "esta sí que es buena")
	nj := NotJSONableNError{}
	l.ErrorE(nj, nil, "otro mejor")

	fmt.Println()

	fmt.Println("Con funciones del paquete")
	gologops.SetFlags(gologops.Lmethod)
	gologops.SetFlags(gologops.Lshortfile)
	gologops.Infof("%d y %d son %d", 2, 2, 4)
	gologops.Info("y ocho dieciséis")
	gologops.SetContext(gologops.C{"prefix": "prefijo"})
	gologops.InfoC(gologops.C{"local": "España y olé"}, "%d y %d son %d", 2, 2, 4)
	gologops.InfoC(gologops.C{"local": `{"json":"pompón"}`}, "y ocho dieciséis")

}
