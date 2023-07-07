package repositories

type SampleRepository struct {
}

func NewSampleRepository() *SampleRepository {
	return &SampleRepository{}
}

func (r *SampleRepository) Get() string {
	// values retrieve from database.
	return "sample repository"
}
