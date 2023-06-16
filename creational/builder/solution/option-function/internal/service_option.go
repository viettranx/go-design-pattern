package internal

func WithName(name string) Option {
	return func(s *complexService) {
		s.name = name
	}
}

func WithStdLogger() Option {
	return func(s *complexService) {
		s.logger = StdLogger{}
	}
}

func WithFileLogger() Option {
	return func(s *complexService) {
		s.logger = FileLogger{}
	}
}

func WithCustomLogger(logger Logger) Option {
	return func(s *complexService) {
		if logger == nil {
			s.logger = StdLogger{}
			return
		}

		s.logger = logger
	}
}

func WithMySQLDataLayer() Option {
	return func(s *complexService) {
		s.dataLayer = MySQLDataLayer{}
	}
}

func WithEmailNotifier() Option {
	return func(s *complexService) {
		s.notifier = EmailNotifier{}
	}
}

func WithAWSUploader() Option {
	return func(s *complexService) {
		s.uploader = AWSS3Uploader{}
	}
}
