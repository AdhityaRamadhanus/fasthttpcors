package main

import (
	"log"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/valyala/fasthttp"
)

func main() {
	withCors := cors.DefaultHandler()
	if err := fasthttp.ListenAndServe(":8080", withCors.CorsMiddleware(requestHandler)); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("plain/text")
	ctx.SetStatusCode(200)
	ctx.SetBodyString("OK")
}
