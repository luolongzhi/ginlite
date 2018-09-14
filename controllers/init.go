package controllers

type ControllerGroup struct {
    TestController *TestController
}

func Init() *ControllerGroup {
    //gen global ctrlGroup
    ctrlGroup := new(ControllerGroup)

    //regist each controller, such as TestController ...
    ctrlGroup.TestController = new(TestController)

    return ctrlGroup;
}

