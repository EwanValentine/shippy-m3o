package repository

import (
	"encoding/json"
	"fmt"

	gostore "github.com/micro/go-micro/v3/store"
	"github.com/micro/micro/v3/service/store"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Vessel struct {
	ID        string `json:"id"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"max_weight"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerID   string `json:"owner_id"`
}

const (
	TableKey = "vessels"
)

func key(id string) string {
	return fmt.Sprintf("%s:%s", TableKey, id)
}

// MicroRepository - KeyValue store repository for the Micro platform
type MicroRepository struct{}

// Save a vessel
func (r *MicroRepository) Save(v *Vessel) error {
	v.ID = uuid.NewV4().String()
	body, err := json.Marshal(v)
	if err != nil {
		return errors.Wrap(err, "error marshalling vessel")
	}

	return store.Write(&gostore.Record{
		Key:   key(v.ID),
		Value: body,
	})
}

// Find a single vessel
func (r *MicroRepository) Find(id string) (*Vessel, error) {
	result, err := store.Read(id)
	if err != nil {
		return nil, errors.Wrap(err, "no item found with id "+id)
	}

	v := &Vessel{}
	if err := json.Unmarshal(result[0].Value, v); err != nil {
		return nil, errors.Wrap(err, "error unmarshalling result")
	}

	return v, nil
}

// List vessels
func (r *MicroRepository) List() ([]*Vessel, error) {
	key := fmt.Sprintf("%s:", TableKey)
	results, err := store.Read(key, gostore.ReadPrefix())
	if err != nil {
		return nil, errors.Wrap(err, "error listing vessels")
	}

	vessels := make([]*Vessel, len(results))
	for _, val := range results {
		vessel := &Vessel{}
		if err := json.Unmarshal(val.Value, vessel); err != nil {
			return nil, errors.Wrap(err, "error unmarshalling vessels from results")
		}
		vessels = append(vessels, vessel)
	}

	return vessels, nil
}
