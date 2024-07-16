package templates

import (
	"github.rbx.com/roblox/entity-schema-generator/models"
)

// ParseTemplateForEntity parses a template for a given entity.
func ParseTemplateForEntity(model *models.Entity) (bllText, dalText, platformInterface, platformImplementation string, err error) {
	switch *model.Version {
	case 1:
		dalText, err = parseTemplateForEntityDalV1(model)
	default:
		dalText, err = parseTemplateForEntityDalV2(model)
	}

	if err != nil {
		return "", "", "", "", err
	}

	bllText, err = parseTemplateForEntityBllV1(model)
	if err != nil {
		return "", "", "", "", err
	}

	platformInterface, err = parseTemplateForPlatformEntityInterface(model)
	if err != nil {
		return "", "", "", "", err
	}

	platformImplementation, err = parseTemplateForPlatformEntityImplementation(model)
	if err != nil {
		return "", "", "", "", err
	}

	return bllText, dalText, platformInterface, platformImplementation, nil
}
