# GRAPHQL in Go 

This project is a simple GRAPHQL created with languaje GO that allows managing the SHIPPMENT domain, specifically for the LIST for HISTORIC REPORT microservice. The API offers the basic operation such as list a historical shippments, displaying the GRAPHIQL documentation technology screen as the main page.
## Project Structure

- **`main.go`**: The main class that runs the Go application and defines the endpoint graphql.

- `POST /graphql`: Allows you to list historic shippmets, under the required query.

## Requirements

- **GO 1.19** o superior
- **Gestor de dependencias como go mod.**

## Installation

1. **Clone the repository**

    ```bash
    git clone <https://github.com/kevinseya/microservicio-logistic-historical-shipping-reports.git>
    ```

2. **Install dependencies**

    ```bash
    go mod tidy
    ```

3. The application run on: `http://localhost:8080`.

## Use of GraphQL

### 1. POST /graphql

List shippmenst. The request body must contain the user details in JSON format with query
POST request example:
```bash
POST /graphql Content-Type: application/json
    
    {
  "query": "query { getAllShipments { shipmentID orderID carrierID dateAsignment } }"
}

```
**Response:**
```json
   {
     "data": {
        "getAllShipments": [
            {
                "shipmentID": "282209f7-5c51-4ae1-9df5-9fb7919986c9",
                "orderID": 10000,
                "carrierID": 101,
                "dateAsignment": "2025-01-27T14:15:14.853-05:00"
            }]
        }
   }
```
**Response code:**
- **`200 OK:`** List shippments.
- **`400 Bad Request:`** Bad Request.
- **`500 Internal Server Error:`** Server error.
