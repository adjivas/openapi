package EIRSelection

type APIClient struct {
	cfg    *Configuration
	common service

	// API Services
	EIREquipmentStatusApi *EIRApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.EIREquipmentStatusApi = (*EIRApiService)(&c.common)

	return c
}
