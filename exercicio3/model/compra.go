package model

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Items   []string
}

func New(mercado string, data time.Time, itens []string) *Compra {
	c := &Compra{Mercado: mercado, Data: data, Items: itens}
	return c

}

func (c *Compra) Update(mercado string, data time.Time, itens []string) {
	c.Data = data
	c.Items = itens
	c.Mercado = mercado

}

func NewSemPonteiro(mercado string, data time.Time, itens []string) Compra {
	c := Compra{Mercado: mercado, Data: data, Items: itens}
	return c

}

func NewComTratamentoDeErro(mercado string, data time.Time, itens []string) (*Compra, error) {
	if mercado == "" {
		return nil, errors.New("mercado precisa ser informado")
	}

	return &Compra{Mercado: mercado, Data: data, Items: itens}, nil

}
