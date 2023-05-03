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
    res, err := s.AddPet(ctx, shared.NewPet{
        Name: "Terrence Rau",
        Tag: sdk.String("nulla"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->