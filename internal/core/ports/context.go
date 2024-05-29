package ports

type HTTPContext interface {
	JSON(int, interface{})
}
