# SDK

## Overview

A sample API that uses a petstore as an example to demonstrate features in the OpenAPI 3.0 specification

### Available Operations

* [AddPet](#addpet) - Creates a new pet in the store. Duplicates are allowed
* [DeletePet](#deletepet) - deletes a single pet based on the ID supplied
* [FindPetByID](#findpetbyid) - Returns a user based on a single ID, if the user does not have access to the pet
* [FindPets](#findpets) - Returns all pets from the system that the user has access to
Nam sed condimentum est. Maecenas tempor sagittis sapien, nec rhoncus sem sagittis sit amet. Aenean at gravida augue, ac iaculis sem. Curabitur odio lorem, ornare eget elementum nec, cursus id lectus. Duis mi turpis, pulvinar ac eros ac, tincidunt varius justo. In hac habitasse platea dictumst. Integer at adipiscing ante, a sagittis ligula. Aenean pharetra tempor ante molestie imperdiet. Vivamus id aliquam diam. Cras quis velit non tortor eleifend sagittis. Praesent at enim pharetra urna volutpat venenatis eget eget mauris. In eleifend fermentum facilisis. Praesent enim enim, gravida ac sodales sed, placerat id erat. Suspendisse lacus dolor, consectetur non augue vel, vehicula interdum libero. Morbi euismod sagittis libero sed lacinia.

Sed tempus felis lobortis leo pulvinar rutrum. Nam mattis velit nisl, eu condimentum ligula luctus nec. Phasellus semper velit eget aliquet faucibus. In a mattis elit. Phasellus vel urna viverra, condimentum lorem id, rhoncus nibh. Ut pellentesque posuere elementum. Sed a varius odio. Morbi rhoncus ligula libero, vel eleifend nunc tristique vitae. Fusce et sem dui. Aenean nec scelerisque tortor. Fusce malesuada accumsan magna vel tempus. Quisque mollis felis eu dolor tristique, sit amet auctor felis gravida. Sed libero lorem, molestie sed nisl in, accumsan tempor nisi. Fusce sollicitudin massa ut lacinia mattis. Sed vel eleifend lorem. Pellentesque vitae felis pretium, pulvinar elit eu, euismod sapien.


## AddPet

Creates a new pet in the store. Duplicates are allowed

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PetStore"
	"PetStore/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.SDK.AddPet(ctx, shared.NewPet{
        Name: "Dallas Kassulke",
        Tag: sdk.String("suscipit"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

## DeletePet

deletes a single pet based on the ID supplied

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PetStore"
	"PetStore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.SDK.DeletePet(ctx, operations.DeletePetRequest{
        ID: 437587,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## FindPetByID

Returns a user based on a single ID, if the user does not have access to the pet

### Example Usage

```go
package main

import(
	"context"
	"log"
	"PetStore"
	"PetStore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.SDK.FindPetByID(ctx, operations.FindPetByIDRequest{
        ID: 297534,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

## FindPets

Returns all pets from the system that the user has access to
Nam sed condimentum est. Maecenas tempor sagittis sapien, nec rhoncus sem sagittis sit amet. Aenean at gravida augue, ac iaculis sem. Curabitur odio lorem, ornare eget elementum nec, cursus id lectus. Duis mi turpis, pulvinar ac eros ac, tincidunt varius justo. In hac habitasse platea dictumst. Integer at adipiscing ante, a sagittis ligula. Aenean pharetra tempor ante molestie imperdiet. Vivamus id aliquam diam. Cras quis velit non tortor eleifend sagittis. Praesent at enim pharetra urna volutpat venenatis eget eget mauris. In eleifend fermentum facilisis. Praesent enim enim, gravida ac sodales sed, placerat id erat. Suspendisse lacus dolor, consectetur non augue vel, vehicula interdum libero. Morbi euismod sagittis libero sed lacinia.

Sed tempus felis lobortis leo pulvinar rutrum. Nam mattis velit nisl, eu condimentum ligula luctus nec. Phasellus semper velit eget aliquet faucibus. In a mattis elit. Phasellus vel urna viverra, condimentum lorem id, rhoncus nibh. Ut pellentesque posuere elementum. Sed a varius odio. Morbi rhoncus ligula libero, vel eleifend nunc tristique vitae. Fusce et sem dui. Aenean nec scelerisque tortor. Fusce malesuada accumsan magna vel tempus. Quisque mollis felis eu dolor tristique, sit amet auctor felis gravida. Sed libero lorem, molestie sed nisl in, accumsan tempor nisi. Fusce sollicitudin massa ut lacinia mattis. Sed vel eleifend lorem. Pellentesque vitae felis pretium, pulvinar elit eu, euismod sapien.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"PetStore"
	"PetStore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.SDK.FindPets(ctx, operations.FindPetsRequest{
        Limit: sdk.Int(891773),
        Tags: []string{
            "delectus",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pets != nil {
        // handle response
    }
}
```
