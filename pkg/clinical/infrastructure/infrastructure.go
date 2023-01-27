package infrastructure

import (
	"github.com/savannahghi/clinical/pkg/clinical/application/common"
	"github.com/savannahghi/clinical/pkg/clinical/application/extensions"
	dataset "github.com/savannahghi/clinical/pkg/clinical/infrastructure/datastore/fhirdataset"
	"github.com/savannahghi/clinical/pkg/clinical/infrastructure/services/mycarehub"
	"github.com/savannahghi/clinical/pkg/clinical/infrastructure/services/openconceptlab"
	"github.com/savannahghi/clinical/pkg/clinical/repository"
)

// Infrastructure ...
type Infrastructure struct {
	FHIRRepo       dataset.FHIRRepository
	FHIR           repository.FHIR
	OpenConceptLab openconceptlab.ServiceOCL
	BaseExtension  extensions.BaseExtension
	MyCareHub      mycarehub.IServiceMyCareHub
}

// NewInfrastructureInteractor initializes a new Infrastructure
func NewInfrastructureInteractor(
	ext extensions.BaseExtension,
	fhirRepository dataset.FHIRRepository,
	fhir repository.FHIR,
	openconceptlab openconceptlab.ServiceOCL,
) Infrastructure {
	myCareHubClient := common.NewInterServiceClient("mycarehub", ext)
	mycarehub := mycarehub.NewServiceMyCareHub(myCareHubClient, ext)

	return Infrastructure{
		fhirRepository,
		fhir,
		openconceptlab,
		ext,
		mycarehub,
	}
}
