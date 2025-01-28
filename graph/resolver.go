package graph

import (
	"context"
	"encoding/hex"
	"fmt"
	"historical-shipping-reports/database"
	"historical-shipping-reports/graph/generated"
	graph "historical-shipping-reports/models"
)

// Resolver is the root resolver for the GraphQL API.
type Resolver struct{}

// GetAllShipments is the resolver for the getAllShipments field.
func (r *queryResolver) GetAllShipments(ctx context.Context) ([]*graph.Shipment, error) {
	var shipments []*graph.Shipment
	// Get all Shipments of database
	if err := database.DB.Find(&shipments).Error; err != nil {
		return nil, err
	}

	// Converter shipmentID of string to hexadecimal
	for i, shipment := range shipments {
		if shipment.ShipmentID != "" {
			// Converter of string to hexadecimal
			hexID := hex.EncodeToString([]byte(shipment.ShipmentID))

			// ID hexadecimal have 32 characters
			if len(hexID) == 32 {
				// Insert (-) of formated UUID
				shipments[i].ShipmentID = fmt.Sprintf("%s-%s-%s-%s-%s",
					hexID[:8], hexID[8:12], hexID[12:16], hexID[16:20], hexID[20:])
			} else {
				// If length is not adequate, resolve error
				return nil, fmt.Errorf("shipmentID hexadecimal invalid")
			}
		}
	}

	return shipments, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
