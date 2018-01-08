package handlers

import (
	"net/http"
)

func GetJobs(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func CreateJob(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func GetJob(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func UpdateJob(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func DeleteJob(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func StartProcessingProvidedJobs(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func StopProcessingAllJobs(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func StartProcessingJob(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}

func StopProcessingJob(w http.ResponseWriter, r *http.Request) () {
	CallReceived(r)

	type Response struct {
	}
	Respond(w, Response{})
}
