package internal

type Service interface {
	DoBusiness()
}

type Builder interface {
	reset()
	setName(name string)
	buildLogger(logger Logger)
	buildNotifier(notifier Notifier)
	buildDataLayer(dataLayer DataLayer)
	buildUploader(uploader Uploader)
	result() Service
}

type serviceBuilder struct {
	service *complexService
}

func NewBuilder() *serviceBuilder {
	return &serviceBuilder{}
}

func (builder *serviceBuilder) reset() { builder.service = &complexService{} }

func (builder *serviceBuilder) setName(name string) { builder.service.setName(name) }

func (builder *serviceBuilder) buildLogger(logger Logger) {
	if logger == nil {
		logger = StdLogger{}
	}

	builder.service.setLogger(logger)
}

func (builder *serviceBuilder) buildNotifier(notifier Notifier) {
	if notifier == nil {
		notifier = SMSNotifier{}
	}

	builder.service.setNotifier(notifier)
}

func (builder *serviceBuilder) buildDataLayer(dataLayer DataLayer) {
	if dataLayer == nil {
		dataLayer = MySQLDataLayer{}
	}

	builder.service.setDataLayer(dataLayer)
}

func (builder *serviceBuilder) buildUploader(uploader Uploader) {
	if uploader == nil {
		uploader = AWSS3Uploader{}
	}

	builder.service.setUploader(uploader)
}

func (builder *serviceBuilder) result() Service {
	return builder.service
}
