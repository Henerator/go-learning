package robot

import "fmt"

type Robot struct {
	Model       string
	workCounter int
}

func (r Robot) String() string {
	return fmt.Sprintf("Robot %s", r.Model)
}

func (r *Robot) Work(tasks []string) string {
	res := r.Model + " work:"
	for _, task := range tasks {
		res += "\n I do " + task
	}

	r.workCounter += len(tasks)
	return res
}
