package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/clivern/hamster/internal/app/sender"
    "github.com/clivern/hamster/pkg"
)

func CreateCommentTest(c *gin.Context) {

    var message sender.Comment
    message.Body = "Hello World!"

    err, data := message.ConvertToJSON()

    if err == nil {
        response, err := pkg.Request(
            "POST",
            "https://api.github.com/repos/Clivern/Hamster/issues/1/comments",
            data,
            "bbb...",
        )

        if err != nil{
            c.JSON(200, gin.H{
                "status": "Request Error!",
            })
        }else{
            c.JSON(200, gin.H{
                "status": response,
            })
        }
    }else{
        c.JSON(200, gin.H{
            "status": "Not Ok",
        })
    }
}