/*
 * swaggerpetstore_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package models_pkg

import "time"


/*
 * Structure for the custom type Tag
 */
type Tag struct {
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //TODO: Write general description for this field
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type Category
 */
type Category struct {
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //TODO: Write general description for this field
    Name            *string         `json:"name,omitempty" form:"name,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type Pet
 */
type Pet struct {
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //TODO: Write general description for this field
    Category        Category        `json:"category,omitempty" form:"category,omitempty"` //TODO: Write general description for this field
    Name            string          `json:"name" form:"name"` //TODO: Write general description for this field
    PhotoUrls       []string        `json:"photoUrls" form:"photoUrls"` //TODO: Write general description for this field
    Tags            []*Tag          `json:"tags,omitempty" form:"tags,omitempty"` //TODO: Write general description for this field
    Status          Status1Enum     `json:"status,omitempty" form:"status,omitempty"` //pet status in the store
}

/*
 * Structure for the custom type User
 */
type User struct {
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //TODO: Write general description for this field
    Username        *string         `json:"username,omitempty" form:"username,omitempty"` //TODO: Write general description for this field
    FirstName       *string         `json:"firstName,omitempty" form:"firstName,omitempty"` //TODO: Write general description for this field
    LastName        *string         `json:"lastName,omitempty" form:"lastName,omitempty"` //TODO: Write general description for this field
    Email           *string         `json:"email,omitempty" form:"email,omitempty"` //TODO: Write general description for this field
    Password        *string         `json:"password,omitempty" form:"password,omitempty"` //TODO: Write general description for this field
    Phone           *string         `json:"phone,omitempty" form:"phone,omitempty"` //TODO: Write general description for this field
    UserStatus      *int64          `json:"userStatus,omitempty" form:"userStatus,omitempty"` //User Status
}

/*
 * Structure for the custom type ApiResponse
 */
type ApiResponse struct {
    Code            *int64          `json:"code,omitempty" form:"code,omitempty"` //TODO: Write general description for this field
    Type            *string         `json:"type,omitempty" form:"type,omitempty"` //TODO: Write general description for this field
    Message         *string         `json:"message,omitempty" form:"message,omitempty"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type Order
 */
type Order struct {
    Id              *int64          `json:"id,omitempty" form:"id,omitempty"` //TODO: Write general description for this field
    PetId           *int64          `json:"petId,omitempty" form:"petId,omitempty"` //TODO: Write general description for this field
    Quantity        *int64          `json:"quantity,omitempty" form:"quantity,omitempty"` //TODO: Write general description for this field
    ShipDate        *time.Time      `json:"shipDate,omitempty" form:"shipDate,omitempty"` //TODO: Write general description for this field
    Status          StatusEnum      `json:"status,omitempty" form:"status,omitempty"` //Order Status
    Complete        *bool           `json:"complete,omitempty" form:"complete,omitempty"` //TODO: Write general description for this field
}
