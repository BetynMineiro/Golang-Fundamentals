package main

import (
	"exercicio3/model"
	"fmt"
	"time"
)

func main() {
	c := model.New("mercado 1", time.Now(), []string{"produto 1", "produto 2"})

	fmt.Println("compra 1 ->", c)

	c.Update("mercado 2", time.Now(), []string{"produto 3", "produto 4"})

	fmt.Println("compra 2 ->", c)

	d := model.NewSemPonteiro("mercado 3 do centro", time.Now(), []string{"produto 5", "produto 9"})

	fmt.Println("compra 3 sem ponteiro ->", d)

	d.Update("mercado 4 do leste", time.Now(), []string{"produto 13", "produto 45"})

	fmt.Println("compra 4 sem ponteiro ->", d)

	_, err := model.NewComTratamentoDeErro("", time.Now(), []string{"produto 58", "produto 99"})

	fmt.Println("compra Com Erro ->", err)

}
