package statsservice

import (
	"context"
	"fmt"
	"log"
	"sort"

	"github.com/aman-saraiya/remote-procedure-calls/grpc-protobuf/client/generated/statsservice"
)

// Define the StatsService gRPC service struct
type StatsServiceServer struct {
	statsservice.UnimplementedStatsServiceServer
	numbers []float64
}

// Function to create StatsServiceServer instance
func NewStatsServiceServer() *StatsServiceServer{
	statsSvcServer := &StatsServiceServer{}
	return statsSvcServer
}

// Handler for GetStats gRPC
func (s *StatsServiceServer) GetStats(ctx context.Context, req *statsservice.GetStatsArg) (*statsservice.GetStatsRet, error) {
	log.Println("Handling GetStats RPC with argument ", req)
	numbers := req.Numbers
	// Calculate mean
	mean := calculateMean(numbers)

	// Calculate median
	median := calculateMedian(numbers)

	return &statsservice.GetStatsRet{
		Mean:   mean,
		Median: median,
	}, nil
}

// Handler for RunningStats gRPC
func (s *StatsServiceServer) RunningStats(stream statsservice.StatsService_RunningStatsServer) error {
	log.Println("Handling RunningStats RPC")
	var runningNumbers []float64

	for {
		// Receiving the stream of numbers sent by the client
		req, err := stream.Recv()
		
		if err != nil {
			// EOF or other error. EOF indicates the client has closed the stream
			if err.Error() == "EOF" {
				fmt.Println("Client has closed the stream")
				break
			}
			fmt.Printf("Error receiving message: %v", err)
			return err
		}
		log.Println("Data received over stream ", req)

		runningNumbers = append(runningNumbers, req.Number)
		log.Println("Calculating mean and median for ", runningNumbers)

		// Calculate running mean
		runningMean := calculateMean(runningNumbers)

		// Calculate running median
		runningMedian := calculateMedian(runningNumbers)

		// Send the running stats back
		if err := stream.Send(&statsservice.RunningStatsRet{
			Mean:   runningMean,
			Median: runningMedian,
		}); err != nil {
			return err
		}
	}
	return nil
}

// Calculate the mean of a list of numbers
func calculateMean(numbers []float64) float64 {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

// Calculate the median of a list of numbers
func calculateMedian(numbers []float64) float64 {
	sort.Float64s(numbers)
	n := len(numbers)
	if n%2 == 0 {
		return (numbers[n/2-1] + numbers[n/2]) / 2
	}
	return numbers[n/2]
}