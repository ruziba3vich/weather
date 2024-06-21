package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	genprotos "github.com/ruziba3vich/local_weather_gateway/genprotos/protos"
	"google.golang.org/grpc"
)

type handler struct {
	client genprotos.WeatherServiceClient
	logger *log.Logger
}

func New(host string, conn *grpc.ClientConn, logger *log.Logger) *handler {
	return &handler{
		client: genprotos.NewWeatherServiceClient(conn),
		logger: logger,
	}
}

func (h *handler) GetWeatherByCountryName(c *gin.Context) {
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Transfer-Encoding", "chunked")

	req := genprotos.GetWeatherByCountryNameReq{
		CountryName: c.Query("country_name"),
	}

	stream, err := h.client.GetWeatherByCountryName(c, &req)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			if err != nil {
				// c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
				h.logger.Println(err)
				return
			}

		}
		c.IndentedJSON(http.StatusOK, response)
		// _, err = c.Writer.Write(response.ProtoReflect().GetUnknown())
		// if err != nil {
		// 	c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		// 	h.logger.Println(err)
		// 	return
		// }

		// if flusher, ok := c.Writer.(http.Flusher); ok {
		// 	flusher.Flush()
		// } else {
		// 	h.logger.Println("Streaming not supported by the client")
		// }
	}
}
