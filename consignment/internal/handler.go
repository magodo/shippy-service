package internal

import (
	"context"
	"log"

	pb "github.com/magodo/shippy-service/consignment/proto/consignment"
	vesselPb "github.com/magodo/shippy-service/vessel/proto/vessel"
)

type Handler struct {
	Repo         repository
	VesselClient vesselPb.VesselService
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the rpc server.
func (s *Handler) Create(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	// find a vessel to ship the consignment
	vesselResp, err := s.VesselClient.FindAvailable(ctx, &vesselPb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if err != nil {
		return err
	}
	log.Printf("Found vessel: %s\n", vesselResp.Vessel.Name)

	// use the got vessel id for creating consignment
	req.VesselId = vesselResp.Vessel.Id

	// Save our consignment
	err = s.Repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	resp.Created = true
	resp.Consignment = req
	return nil
}

func (s *Handler) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) error {
	consignments, err := s.Repo.GetAll()
	if err != nil {
		return err
	}
	resp.Consignments = consignments
	return nil
}
