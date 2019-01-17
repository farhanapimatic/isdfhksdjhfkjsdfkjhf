/*
 * swaggerpetstore_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
 */

package pet_pkg

import "swaggerpetstore_lib/models_pkg"
import "swaggerpetstore_lib/configuration_pkg"

/*
 * Interface for the PET_IMPL
 */
type PET interface {
    UploadFile (int64, *string, []byte) (*models_pkg.ApiResponse, error)

    UpdatePetWithForm (int64, *string, *string) (error)

    DeletePet (int64, *string) (error)

    AddPet (*models_pkg.Pet) (error)

    FindPetsByTags ([]string) ([]*models_pkg.Pet, error)

    GetPetById (int64) (*models_pkg.Pet, error)

    UpdatePet (*models_pkg.Pet) (error)

    FindPetsByStatus ([]models_pkg.Status2Enum) ([]*models_pkg.Pet, error)
}

/*
 * Factory for the PET interaface returning PET_IMPL
 */
func NewPET(config configuration_pkg.CONFIGURATION) *PET_IMPL {
    client := new(PET_IMPL)
    client.config = config
    return client
}