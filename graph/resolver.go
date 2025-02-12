package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"fmt"
	"historical-shipping-reports/database"
	"historical-shipping-reports/graph/generated"
	graph "historical-shipping-reports/models"
	"log"

	"github.com/google/uuid"
)

type Resolver struct{}

// GetAll Shipments is the resolver for the getAll Shipments field.
func (r *queryResolver) GetAllShipments(ctx context.Context) ([]*graph.Shipments, error) {
	var shipments []*graph.Shipments
	// Get all shipments from the database
	if err := database.DB.Find(&shipments).Error; err != nil {
		return nil, err
	}
	log.Println("Shipments Data Received: ", shipments)

	// Validation of UUIDs: shipmentID, orderID, and carrierID
	for _, shipment := range shipments {
		// Validate shipmentID
		if shipment.ID != "" {
			if _, err := uuid.Parse(shipment.ID); err != nil {
				return nil, fmt.Errorf("shipmentID no es un UUID válido: %v", err)
			}
		}

		// Validate orderID
		if shipment.OrderID != "" {
			if _, err := uuid.Parse(shipment.OrderID); err != nil {
				return nil, fmt.Errorf("orderID no es un UUID válido: %v", err)
			}
		}

		// Validate user_carrier_id
		if shipment.UserCarrierID != "" {
			if _, err := uuid.Parse(shipment.UserCarrierID); err != nil {
				return nil, fmt.Errorf("user_carrier_id no es un UUID válido: %v", err)
			}
		}
	}

	return shipments, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
