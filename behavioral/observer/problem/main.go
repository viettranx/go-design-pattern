package main

type Developer struct{}

type Job struct {
	Title string
}

type ITJobsCompany struct {
	jobs []Job
}

func DeveloperAskForANewJob(dev Developer, comp ITJobsCompany) {
	// Developer: Hey, do you have any new job for me?
	// Company: Huh? Who are you? These are all jobs we have.
	// [...company display all jobs]
	// Developer: No interesting job. Can you notify me when you have a new job (or a specific one)?
	// Company: We're developing this feature.
	// Developer: ...
}
