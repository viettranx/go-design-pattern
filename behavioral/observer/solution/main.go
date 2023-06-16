package main

import "fmt"

type Observer interface {
	ReceiveNotify(j Job)
}

type Developer struct{}

func (Developer) ReceiveNotify(j Job) { fmt.Println("Many thanks, I've received job:", j.Title) }

type Job struct {
	Title string
}

// Aka: Subject

type ITJobsCompany struct {
	jobs      []Job
	observers []Observer
}

func (comp *ITJobsCompany) AddObserver(o Observer) {
	comp.observers = append(comp.observers, o)
}

func (comp *ITJobsCompany) RemoveObserver(o Observer) {
	for i := range comp.observers {
		if comp.observers[i] == o {
			comp.observers = append(comp.observers[:i], comp.observers[i+1:]...)
			return
		}
	}
}

func (comp *ITJobsCompany) notifyToObservers(j Job) {
	for i := range comp.observers {
		comp.observers[i].ReceiveNotify(j)
	}
}

func (comp *ITJobsCompany) AddNewJob(j Job) {
	comp.jobs = append(comp.jobs, j)

	comp.notifyToObservers(j)
}

func main() {
	itComp := ITJobsCompany{}
	developer := Developer{}

	// Developer will be added as an observer
	itComp.AddObserver(developer)

	// Developer will receive new jobs
	itComp.AddNewJob(Job{Title: "Senior Go backend engineer"})
	itComp.AddNewJob(Job{Title: "Junior React developer"})

	// Developer no long receive new job
	itComp.RemoveObserver(developer)

	itComp.AddNewJob(Job{Title: "Some boring IT job"})
}
