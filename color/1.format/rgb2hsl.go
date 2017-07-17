package format

import (
	"errors"
	"fmt"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func Convert(c[]float64) (error,[]float64) {
if len(c)!=3{
return errors.New(fmt.Sprintf("color dimension=[%d] error", len(c))), nil
	}else{
c[0],c[1],c[2]=colorful.Color{c[0],c[1],c[2]}.Hsl()
return nil,c
	}
}
