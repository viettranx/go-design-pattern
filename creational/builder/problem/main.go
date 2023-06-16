package main

type Logger interface {
	Log(...any)
}

type StdLogger struct{}
type FileLogger struct{}

func (StdLogger) Log(...any)  {}
func (FileLogger) Log(...any) {}

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}
type SMSNotifier struct{}

func (EmailNotifier) Send(message string) {}
func (SMSNotifier) Send(message string)   {}

type DataLayer interface {
	Save()
}

type MySQLDataLayer struct{}
type MongoDataLayer struct{}

func (MySQLDataLayer) Save() {}
func (MongoDataLayer) Save() {}

type Uploader interface {
	Upload()
}

type AWSS3Uploader struct{}
type GoogleDriveUploader struct{}

func (AWSS3Uploader) Upload()       {}
func (GoogleDriveUploader) Upload() {}

type ComplexService struct {
	name      string
	logger    Logger
	notifier  Notifier
	dataLayer DataLayer
	uploader  Uploader
}

func (s *ComplexService) SetLogger(l Logger)        { s.logger = l }
func (s *ComplexService) SetNotifier(n Notifier)    { s.notifier = n }
func (s *ComplexService) SetDataLayer(dl DataLayer) { s.dataLayer = dl }
func (s *ComplexService) SetUploader(u Uploader)    { s.uploader = u }

func (s *ComplexService) DoBusiness() {
	// use all components to do some business
	s.logger.Log(s.name)
	s.uploader.Upload()
	s.dataLayer.Save()
	s.notifier.Send("hello world")
}

// God constructor with too many parameters

func NewService(name string, logger Logger, notifier Notifier, dataLayer DataLayer, uploader Uploader) ComplexService {
	// A lot of parameter checking logics
	// to guarantee no component is `nil`

	if logger == nil {
		logger = StdLogger{}
	}

	if notifier == nil {
		notifier = SMSNotifier{}
	}

	if dataLayer == nil {
		dataLayer = MySQLDataLayer{}
	}

	if uploader == nil {
		uploader = AWSS3Uploader{}
	}

	return ComplexService{
		name:      name,
		logger:    logger,
		notifier:  notifier,
		dataLayer: dataLayer,
		uploader:  uploader,
	}
}

// Other convenient constructors... We also have a lot them

func NewServiceWithName(name string) ComplexService {
	return NewService(
		name,

		// default modules
		StdLogger{},
		EmailNotifier{},
		MongoDataLayer{},
		AWSS3Uploader{},
	)
}

func main() {

}
