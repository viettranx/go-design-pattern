package internal

type ServiceDirector interface {
	BuildService(builder Builder) Service
}

type serviceBuilderDirector struct{}

// I am director, I know how to build a service step by step,
// but I don't want to do it myself.
// Hey builder, come to take your job.

func (sbd serviceBuilderDirector) BuildService(builder Builder) Service {
	builder.reset()
	builder.setName("Complex Service")
	builder.buildLogger(StdLogger{})
	builder.buildNotifier(EmailNotifier{})
	builder.buildDataLayer(MySQLDataLayer{})
	builder.buildUploader(AWSS3Uploader{})

	return builder.result()
}

func NewDirector() serviceBuilderDirector {
	return serviceBuilderDirector{}
}
