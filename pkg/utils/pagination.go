package utils

import (
    "strconv"
    "github.com/gin-gonic/gin"
    "github.com/gocql/gocql"
)

type Pagination struct {
    Limit  int `form:"limit" json:"limit"`   // form tag ekleyin
    Page   int `form:"page" json:"page"`     // form tag ekleyin
    Offset int `json:"offset"`
}

func GetPagination(c *gin.Context) Pagination {
    limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
    if err != nil || limit <= 0 {
        limit = 10
    }

    page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
    if err != nil || page <= 0 {
        page = 1
    }

    offset := (page - 1) * limit

    return Pagination{
        Limit:  limit,
        Page:   page,
        Offset: offset,
    }
}

func Paginate(c *gin.Context, query *gocql.Query) *gocql.Query {
    pagination := GetPagination(c)
    return query.PageSize(pagination.Limit)
}