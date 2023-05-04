package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
)

type Product struct {
    mgm.DefaultModel `bson:",inline"`
	Name string `json:"name" bson:"name"`
}

func main() {

    // app.Get("/product/:id", func(c *fiber.Ctx) error {
    //     product := &Product{}
    //     coll := mgm.Coll(product)

    //     // Find and decode the doc to a book model.
    //     err := coll.FindByID(c.Params("id"), product)

    //     if err != nil {
    //         fmt.Println(err)
    //     }
        
    //     return c.JSON(product)
    // });
    r := gin.Default()

    r.StaticFile("/loaderio-04cbc2e6e8994582817d57faa8742ee5", "./loaderio-04cbc2e6e8994582817d57faa8742ee5.html")

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
        })
    })

    r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}