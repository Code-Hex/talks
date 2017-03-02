package main

type Service struct {
	Exec   string
	Every  string
	Worker int
	Tag    string
}

func main() {
	Add(Service{
		Exec:   "echo 'This is cron!!'",
		Every:  "30 * * * * *",
		Worker: 3,
		Tag:    "cron",
	})
}

func Add(service Service) {
	if service.Exec != "" {
		// Do something
	}
	if service.Worker > 0 {
		// Do something
	}
}

// END OMIT
