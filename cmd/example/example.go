package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/TDAF/gologops"
)

type complexErr struct {
	Text  string
	Cause *complexErr
}

func (ce complexErr) Error() string { return "uno complejito" }

type procInfo struct {
	hostname string
	pid      int32
}

func f() interface{} {
	return procInfo{
		hostname: "localhost",
		pid:      3452354,
	}
}

type NotJSONableNError struct{}

func (NotJSONableNError) Error() string { return "a very strange error" }
func (NotJSONableNError) MarshalJSON() ([]byte, error) {
	return nil, errors.New("JSON not supported")
}

func main() {
	l := gologops.NewLogger()
	l.AddFlags(gologops.Lmethod)
	l.AddFlags(gologops.Llongfile)
	l.AddFlags(gologops.Lshortfile)

	l.Infof("%d y %d son %d", 2, 2, 4)
	l.Info("y ocho dieciséis")
	l.SetContext(gologops.C{"prefix": "prefijo"})
	l.SetContextFunc(func() gologops.C {
		hostname, _ := os.Hostname()
		pid := strconv.Itoa(os.Getpid())
		return gologops.C{
			"hostname": hostname,
			"pid":      pid,
		}
	})
	l.InfoC(gologops.C{"local": "España y olé"}, "%d y %d son %d", 2, 2, 4)
	l.InfoC(gologops.C{"local": `{"json":"pompón"}`}, "y ocho dieciséis")
	l.DebugC(gologops.C{"local": `{"json":"pompón con debug"}`}, "más o menos")

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

	otherErr := complexErr{"The 1 is another err...", &complexErr{"that nests the number 2 err", nil}}
	gologops.FatalE(otherErr, gologops.C{"msisdn": "+34677876568", "center": "5.5"}, "con más mensaje")

}
