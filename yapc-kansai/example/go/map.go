package main

func main() {
	service := make(map[string]interface{})
	service["Exec"] = "echo 'This is cron!!'"
	service["Every"] = "30 * * * * *"
	service["Worker"] = 3
	service["Tag"] = "cron"
	Add(service)
}

func Add(service map[string]interface{}) {
	if v, ok := service["Exec"]; ok {
		if exec, typeok := v.(string); typeok && exec != "" {
			// Do somethibg
		}
	}

	if v, ok := service["Worker"]; ok {
		if worker, typeok := v.(int); typeok && worker > 0 {
			// Do somethibg
		}
	}
}
