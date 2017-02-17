quit := make(chan os.Signal)
signal.Notify(quit, os.Interrupt)

srv := &http.Server{Addr: ":8080", Handler: http.DefaultServeMux}
go func() {
	<-quit
	log.Println("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("could not shutdown: %v", err)
	}
}()
