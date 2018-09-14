package controllers

import (
    "fmt"
)

//共用controller的属性和方法在此文件中定义
//所有的子controller继承该公用方法

type Controller struct {
}

func (ctrl *Controller) TestPrint() {
    fmt.Println("this is public test controller")
}



