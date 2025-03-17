package cassandra

import (
    "log"
    "github.com/gocql/gocql"
)

func SeedData(session *gocql.Session) {
    categories := []struct {
        ID          int
        Name        string
        Description string
    }{
        {ID: 1, Name: "Italian", Description: "Italian cuisine dishes"},
		{ID: 2, Name: "Mexican", Description: "Spicy Mexican flavors"},
		{ID: 3, Name: "Chinese", Description: "Traditional Chinese meals"},
		{ID: 4, Name: "Indian", Description: "Rich Indian spices"},
		{ID: 5, Name: "French", Description: "Elegant French cuisine"},
		{ID: 6, Name: "Japanese", Description: "Sushi and more"},
		{ID: 7, Name: "Thai", Description: "Thai spicy dishes"},
		{ID: 8, Name: "Greek", Description: "Mediterranean Greek food"},
		{ID: 9, Name: "Spanish", Description: "Tapas and paella"},
		{ID: 10, Name: "Turkish", Description: "Kebabs and baklava"},
		{ID: 11, Name: "American", Description: "Burgers and fries"},
		{ID: 12, Name: "Korean", Description: "Kimchi and BBQ"},
		{ID: 13, Name: "Vietnamese", Description: "Pho and spring rolls"},
		{ID: 14, Name: "Brazilian", Description: "Churrasco and feijoada"},
		{ID: 15, Name: "Moroccan", Description: "Tagines and couscous"},
		{ID: 16, Name: "Russian", Description: "Borscht and dumplings"},
		{ID: 17, Name: "Lebanese", Description: "Hummus and falafel"},
		{ID: 18, Name: "Ethiopian", Description: "Injera and stews"},
		{ID: 19, Name: "German", Description: "Sausages and pretzels"},
		{ID: 20, Name: "Caribbean", Description: "Jerk chicken and rice"},
    }

    for _, cat := range categories {
        err := session.Query(
            `INSERT INTO category (categoryid, categoryname, categorydescription) VALUES (?, ?, ?)`,
            cat.ID, cat.Name, cat.Description,
        ).Exec()
        if err != nil {
            log.Printf("Error inserting category %d: %v", cat.ID, err)
        }
    }
}