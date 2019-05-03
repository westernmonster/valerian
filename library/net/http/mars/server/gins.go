// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"html/template"
	"net/http"
	"sync"

	"valerian/library/net/http/mars"
)

var once sync.Once
var internalEnmarse *mars.Engine

func enmarse() *mars.Engine {
	once.Do(func() {
		internalEnmarse = mars.Default()
	})
	return internalEnmarse
}

// LoadHTMLGlob is a wrapper for Enmarse.LoadHTMLGlob.
func LoadHTMLGlob(pattern string) {
	enmarse().LoadHTMLGlob(pattern)
}

// LoadHTMLFiles is a wrapper for Enmarse.LoadHTMLFiles.
func LoadHTMLFiles(files ...string) {
	enmarse().LoadHTMLFiles(files...)
}

// SetHTMLTemplate is a wrapper for Enmarse.SetHTMLTemplate.
func SetHTMLTemplate(templ *template.Template) {
	enmarse().SetHTMLTemplate(templ)
}

// NoRoute adds handlers for NoRoute. It return a 404 code by default.
func NoRoute(handlers ...mars.HandlerFunc) {
	enmarse().NoRoute(handlers...)
}

// NoMethod is a wrapper for Enmarse.NoMethod.
func NoMethod(handlers ...mars.HandlerFunc) {
	enmarse().NoMethod(handlers...)
}

// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
// For example, all the routes that use a common middleware for authorization could be grouped.
func Group(relativePath string, handlers ...mars.HandlerFunc) *mars.RouterGroup {
	return enmarse().Group(relativePath, handlers...)
}

// Handle is a wrapper for Enmarse.Handle.
func Handle(httpMethod, relativePath string, handlers ...mars.HandlerFunc) mars.IRoutes {
	return enmarse().Handle(httpMethod, relativePath, handlers...)
}

// POST is a shortcut for router.Handle("POST", path, handle)
func POST(relativePath string, handlers ...mars.HandlerFunc) mars.IRoutes {
	return enmarse().POST(relativePath, handlers...)
}

// GET is a shortcut for router.Handle("GET", path, handle)
func GET(relativePath string, handlers ...mars.HandlerFunc) mars.IRoutes {
	return enmarse().GET(relativePath, handlers...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handle)
func DELETE(relativePath string, handlers ...mars.HandlerFunc) mars.IRoutes {
	return enmarse().DELETE(relativePath, handlers...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handle)
func PATCH(relativePath string, handlers ...mars.HandlerFunc) mars.IRoutes {
	return enmarse().PATCH(relativePath, handlers...)
}

// PUT is a shortcut for router.Handle("PUT", path, handle)
func PUT(relativePath string, handlers ...mars.HandlerFunc) mars.IRoutes {
	return enmarse().PUT(relativePath, handlers...)
}

// OPTIONS is a shortcut for router.Handle("OPTIONS", path, handle)
func OPTIONS(relativePath string, handlers ...mars.HandlerFunc) mars.IRoutes {
	return enmarse().OPTIONS(relativePath, handlers...)
}

// HEAD is a shortcut for router.Handle("HEAD", path, handle)
func HEAD(relativePath string, handlers ...mars.HandlerFunc) mars.IRoutes {
	return enmarse().HEAD(relativePath, handlers...)
}

// Any is a wrapper for Enmarse.Any.
func Any(relativePath string, handlers ...mars.HandlerFunc) mars.IRoutes {
	return enmarse().Any(relativePath, handlers...)
}

// StaticFile is a wrapper for Enmarse.StaticFile.
func StaticFile(relativePath, filepath string) mars.IRoutes {
	return enmarse().StaticFile(relativePath, filepath)
}

// Static serves files from the given file system root.
// Internally a http.FileServer is used, therefore http.NotFound is used instead
// of the Router's NotFound handler.
// To use the operating system's file system implementation,
// use :
//     router.Static("/static", "/var/www")
func Static(relativePath, root string) mars.IRoutes {
	return enmarse().Static(relativePath, root)
}

// StaticFS is a wrapper for Enmarse.StaticFS.
func StaticFS(relativePath string, fs http.FileSystem) mars.IRoutes {
	return enmarse().StaticFS(relativePath, fs)
}

// Use attaches a global middleware to the router. ie. the middlewares attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
// For example, this is the right place for a logger or error management middleware.
func Use(middlewares ...mars.HandlerFunc) mars.IRoutes {
	return enmarse().Use(middlewares...)
}

// Routes returns a slice of registered routes.
func Routes() mars.RoutesInfo {
	return enmarse().Routes()
}

// Run attaches to a http.Server and starts listening and serving HTTP requests.
// It is a shortcut for http.ListenAndServe(addr, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func Run(addr ...string) (err error) {
	return enmarse().Run(addr...)
}

// RunTLS attaches to a http.Server and starts listening and serving HTTPS requests.
// It is a shortcut for http.ListenAndServeTLS(addr, certFile, keyFile, router)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func RunTLS(addr, certFile, keyFile string) (err error) {
	return enmarse().RunTLS(addr, certFile, keyFile)
}

// RunUnix attaches to a http.Server and starts listening and serving HTTP requests
// through the specified unix socket (ie. a file)
// Note: this method will block the calling goroutine indefinitely unless an error happens.
func RunUnix(file string) (err error) {
	return enmarse().RunUnix(file)
}

// RunFd attaches the router to a http.Server and starts listening and serving HTTP requests
// through the specified file descriptor.
// Note: the method will block the calling goroutine indefinitely unless on error happens.
func RunFd(fd int) (err error) {
	return enmarse().RunFd(fd)
}
