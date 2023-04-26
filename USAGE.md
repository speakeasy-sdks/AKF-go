<!-- Start SDK Example Usage -->
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
    req := shared.NewPet{
        Name: "Terrence Rau",
        Tag: sdk.String("nulla"),
    }

    res, err := s.AddPet(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->