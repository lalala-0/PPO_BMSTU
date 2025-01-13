package trace

//func NewTraceProvider(exp tracesdk.SpanExporter, ServiceName string) (*tracesdk.TracerProvider, error) {
//	// Ensure default SDK resources and the required service name are set.
//	r, err := resource.Merge(
//		resource.Default(),
//		resource.NewWithAttributes(
//			semconv.SchemaURL,
//			semconv.ServiceNameKey.String(ServiceName),
//		),
//	)
//	if err != nil {
//		return nil, err
//	}
//
//	return tracesdk.NewTracerProvider(
//		tracesdk.WithBatcher(exp),
//		tracesdk.WithResource(r),
//	), nil
//}
