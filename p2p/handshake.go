package p2p

type HandeshakeFunc func(any) error

func NOPHandshakeFunc(any)error{
	return nil
}