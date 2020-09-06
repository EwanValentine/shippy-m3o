package handler

import (
	"context"
	"errors"
	pb "vessels/proto"
	"vessels/repository"
)

// Vessels handler
type Vessels struct {
	repository *repository.MicroRepository
}

func (v *Vessels) find(req *pb.Specification) (*pb.Vessel, error) {
	vessels, err := v.repository.List()
	if err != nil {
		return nil, err
	}

	for _, val := range vessels {
		if req.Capacity <= val.Capacity && req.MaxWeight <= val.MaxWeight {
			return &pb.Vessel{
				Id:        val.ID,
				Name:      val.Name,
				Capacity:  val.Capacity,
				MaxWeight: val.MaxWeight,
				Available: val.Available,
				OwnerId:   val.OwnerID,
			}, nil
		}
	}

	return nil, errors.New("no valid vessel found")
}

// FindAvailable vessels
func (v *Vessels) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := v.find(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}

func (v *Vessels) create(req *pb.Vessel) error {
	return v.repository.Save(&repository.Vessel{
		ID:        req.Id,
		Name:      req.Name,
		Capacity:  req.Capacity,
		MaxWeight: req.MaxWeight,
		Available: req.Available,
		OwnerID:   req.OwnerId,
	})
}

// Create a new vessel
func (v *Vessels) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := v.create(req); err != nil {
		return err
	}
	res.Vessel = req
	return nil
}
