// +build linux

package main

import (
	"net/http"
)

const (
	Port       = 443
	SocketPath = "/dev/socket/"
	IsOnRobot  = true
)

func checkAuth(_ http.ResponseWriter, r *http.Request) (string, error) {
	return "WiFi User Auth Bypass", nil

	// if r.URL.EscapedPath() == "/Anki.Vector.external_interface.ExternalInterface/UserAuthentication" {
	// 	return "WiFi User Auth Bypass", nil
	// }
	// auth, ok := r.Header["Authorization"]

	// if !ok {
	// 	return "", grpc.Errorf(codes.Unauthenticated, "No auth token")
	// }
	// if len(auth) != 1 {
	// 	return "", grpc.Errorf(codes.Unauthenticated, "Too many auth tokens")
	// }
	// authHeader := auth[0]
	// if strings.HasPrefix(authHeader, "Basic ") {
	// 	_, err := base64.StdEncoding.DecodeString(authHeader[6:])
	// 	if err != nil {
	// 		return "", grpc.Errorf(codes.Unauthenticated, "Failed to decode auth token (Base64)")
	// 	}
	// 	// todo
	// } else if !strings.HasPrefix(authHeader, "Bearer ") {
	// 	return "", grpc.Errorf(codes.Unauthenticated, "Unknown auth header type")
	// }
	// clientToken := authHeader[7:]

	// log.Println("==[ Client token: ", clientToken)

	// return tokenManager.CheckToken(clientToken)
}
