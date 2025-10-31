package payment_service

func PaymentMethodList() []string {
	methods := []string{"Credit Card", "PayPal", "Bank Transfer", "Cryptocurrency"}
	return methods
}

func PaymentMethodDetail(name string) string {
	return name + " is a secure and reliable payment method."
}
