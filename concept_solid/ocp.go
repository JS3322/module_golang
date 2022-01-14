package main

// open / closed principle

type DataSender interface {
	SendData(*Data)
}

type FtpDataSender struct {
}

type HttpDataSender struct {
}

func (s *FtpDataSender) SendData(r *Data) {
	// ... send ftp data
}

func (s *HttpDataSender) SendData(r *Data) {
	// ... send http data
}
