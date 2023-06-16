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

func (s *complexService) setName(name string)       { s.name = name }
func (s *complexService) setLogger(l Logger)        { s.logger = l }
func (s *complexService) setNotifier(n Notifier)    { s.notifier = n }
func (s *complexService) setDataLayer(dl DataLayer) { s.dataLayer = dl }
func (s *complexService) setUploader(u Uploader)    { s.uploader = u }

func (s *complexService) DoBusiness() {
	// use all components to do some business
	s.logger.Log(s.name)
	s.uploader.Upload()
	s.dataLayer.Save()
	s.notifier.Send("hello world")

	log.Println("service do business normally")
}
