package register

type ServiceRegister interface {
	Register(scheme, serviceName, serviceAddr string) error
}
