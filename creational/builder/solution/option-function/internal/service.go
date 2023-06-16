package internal

import "log"

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

type complexService struct {
	name      string
	logger    Logger
	notifier  Notifier
	dataLayer DataLayer
	uploader  Uploader
}

func (s *complexService) DoBusiness() {
	// use all components to do some business
	s.logger.Log(s.name)
	s.uploader.Upload()
	s.dataLayer.Save()
	s.notifier.Send("hello world")

	log.Println("service do business normally")
}

type Option func(*complexService)

// God constructor with too many parameters

func NewService(opts ...Option) complexService {

	// Init service with default components
	service := complexService{
		name:      "Service",
		logger:    StdLogger{},
		notifier:  SMSNotifier{},
		dataLayer: MySQLDataLayer{},
		uploader:  AWSS3Uploader{},
	}

	// Execute option functions to modify components
	for i := range opts {
		opts[i](&service)
	}

	return service
}
