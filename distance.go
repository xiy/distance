package main

import (
	"encoding/json"
	levenshtein "github.com/texttheater/golang-levenshtein/levenshtein"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// DistanceService provides an interface to determining string similarity using
// the Levenshtein algorithm.
type DistanceService interface {
	GetDistance(string, string) int
}

type distanceRequest struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type distanceResponse struct {
	Distance int    `json:"distance"`
	Err      string `json:"err,omitempty"`
}

type distanceService struct{}

func main() {
	ctx := context.Background()
	service := distanceService{}

	distanceHandler := httptransport.NewServer(
		ctx,
		makeDistanceEndpoint(service),
		decodeDistanceRequest,
		encodeDistanceResponse)

	http.Handle("/", distanceHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func makeDistanceEndpoint(service DistanceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(distanceRequest)
		distance := service.GetDistance(req.Source, req.Target)
		return distanceResponse{distance, ""}, nil
	}
}

func decodeDistanceRequest(r *http.Request) (interface{}, error) {
	var request distanceRequest

	log.Printf("[%s] request received.", r.Method)

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeDistanceResponse(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func (distanceService) GetDistance(source string, target string) int {
	return levenshtein.DistanceForStrings(
		[]rune(source),
		[]rune(target),
		levenshtein.DefaultOptions)
}
