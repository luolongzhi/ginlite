package controllers

import (
    "fmt"
)

//all public share controller attributes and method can defind here
type Controller struct {
}

func (ctrl *Controller) TestPrint() {
    fmt.Println("this is public test controller")
}



