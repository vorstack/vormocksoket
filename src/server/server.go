package server

const MOCK_FILE_EXTENSION = ".mock"

type Server interface {
	Run(settings Settings)
}

type Settings struct {
	Name    string
	Enable  bool
	Port    int
	MockDir string
	//Params map[string]interface{}
}
