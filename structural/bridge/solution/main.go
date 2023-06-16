package main

type Data struct{}

// Bridge Abstraction and Implementation

//////// Abstraction ////////

type DataParser interface {
	Parse() (*Data, error)
}

type DataPersistent interface {
	Save(*Data) error
}

func parseAndSaveData(parser DataParser, storage DataPersistent) error {
	data, err := parser.Parse()

	if err != nil {
		return err
	}

	if err := storage.Save(data); err != nil {
		return err
	}

	return nil
}

////////////////

//////// Implementation ////////

type MySQLParser struct{}
type MongoParser struct{}
type FileParser struct{}

func (MySQLParser) Parse() (*Data, error) { return &Data{}, nil }
func (MongoParser) Parse() (*Data, error) { return &Data{}, nil }
func (FileParser) Parse() (*Data, error)  { return &Data{}, nil }

type JSONFilePersistent struct{}
type RPCServicePersistent struct{}
type AWSS3Persistent struct{}

func (JSONFilePersistent) Save(*Data) error   { return nil }
func (RPCServicePersistent) Save(*Data) error { return nil }
func (AWSS3Persistent) Save(*Data) error      { return nil }

////////////////

func main() {
	_ = parseAndSaveData(MySQLParser{}, JSONFilePersistent{})
	_ = parseAndSaveData(MongoParser{}, RPCServicePersistent{})
	_ = parseAndSaveData(FileParser{}, AWSS3Persistent{})
}
