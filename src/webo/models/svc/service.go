package svc

type Param map[string] string

type method struct {
    Id string
}

type service struct {
    Id string
}



func Serve(name string, method string, param Param) interface{}{

    return nil
}
