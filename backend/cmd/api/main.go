package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func main() {
    e := echo.New()

    e.POST("/checks", func(c echo.Context) error {
        headers := simplifyArrays(c.Request().Header)
        urlParams := simplifyArrays(c.QueryParams())
        body, err := io.ReadAll(c.Request().Body)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, echo.Map{
                "error": "Failed to read request body",
            })
        }

        resp := echo.Map{
            "source_ip": c.RealIP(),
            "headers": headers,
            "query_params": urlParams,
            "body": echo.Map{
                "size": len(body),
                "content": string(body),
            },
        }
        return c.JSON(http.StatusOK, resp)
    })

    if err := e.Start(":8080"); err != nil {
        log.Fatalf("failed to start echo server: %v\n", err)
    }
}

func simplifyArrays(values map[string][]string) map[string]string {
    res := make(map[string]string, len(values))
    for k, v := range values {
        res[k] = strings.Join(v, " ")
    }
    return res
}
