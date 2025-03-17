package hateoas

import "net/http"

type Link struct {
    Rel    string `json:"rel"`
    Href   string `json:"href"`
    Method string `json:"method"`
}

func GenerateCategoryLinks(r *http.Request, id string) []Link {
    baseURL := "http://" + r.Host + "/api/categories"
    return []Link{
        {Rel: "self", Href: baseURL + "/" + id, Method: "GET"},
        {Rel: "update", Href: baseURL + "/" + id, Method: "PUT"},
        {Rel: "delete", Href: baseURL + "/" + id, Method: "DELETE"},
    }
}