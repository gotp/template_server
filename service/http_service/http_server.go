package http_service

import (
    "crypto/tls"
    "log"
    "net/http"
)

const urlPathPrefix = "/OnlineConsultant/TemplateServer"

func CreateServer() (*http.Server) {
	return &http.Server{
        //TODO: config
        Addr: "127.0.0.1:9003",
        Handler: CreateHandler(),
        // Disable HTTP/2.
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
    }
}

func StartHttpServer(server *http.Server) {
    log.Fatal(server.ListenAndServe())
}

func StartHttpsServer(server *http.Server, pemPath string, keyPath string) {
	log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
}

func CreateHandler() (*http.ServeMux) {
    serverHandler := http.NewServeMux()
    /* 
    * !!!!!!!!!!!! WARNING !!!!!!!!!!!!
    * Please not remove or modify comment 
    * below, it's anchor for new code
    */
    // ############ SERVICE ############
    RegisterTemplateServiceHandler(serverHandler)

    return serverHandler
}

func GetHandlerUrl(path string) string {
	return urlPathPrefix + path;
}