package main

// single responsibility principle

type DataSender interface {
	SendData(*Report)
}

type EmailDataSender struct {
}

type FileDataSender struct {
}

func (s *EmailDataSender) SendData(r *Report) {
	// ... Send Eamil
}

func (s *FileDataSender) SendData(r *Report) {
	// ... make file
}

func main() {

}
