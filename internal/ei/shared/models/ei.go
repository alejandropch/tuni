package ei

import (
	"time"
)

type EI interface { // Educational Institution
	New()
	Get() *EIMetadata
}
type EIMetadata struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

func (m *EIMetadata) GetMetadata() *EIMetadata {
	return &EIMetadata{
		ID:        "123",
		Name:      "Some Institution",
		CreatedAt: time.Now(),
	}
}

/*func (ei *EducationalInstitution) New() *EducationalInstitution {
	fmt.Printf("this is new")
	return &EducationalInstitution{
		ID:        "123",
		Name:      "Some Institution",
		CreatedAt: time.Now(),
	}
}*/
