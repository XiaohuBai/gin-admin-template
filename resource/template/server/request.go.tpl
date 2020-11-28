package request

import "gin-admin-template/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}