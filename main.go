package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
	"log"
)

func main() {
	fmt.Println("ALL GOOD")
	tracer.Start()
	defer tracer.Stop()

	// Create a gin.Engine
	r := gin.New()

	// Use the tracer middleware with your desired service name.
	r.Use(gintrace.Middleware("heimdall-test"))

	// Continue using the router as normal.
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	err := profiler.Start(
		profiler.WithService("<SERVICE_NAME>"),
		profiler.WithEnv("<ENVIRONMENT>"),
		profiler.WithVersion("<APPLICATION_VERSION>"),
		profiler.WithTags("<KEY1>:<VALUE1>,<KEY2>:<VALUE2>"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer profiler.Stop()

	r.Run(":8080")
}
