package model

import "Library_project/repo"

func GetInstances(Instances []repo.Instance) []repo.Instance {
	Instances = []repo.Instance{}
	repo.GetInstancesFromDB(&Instances)
	return Instances
}

func GetInstancesWithPage(Instances []repo.Instance, page string) []repo.Instance {
	Instances = []repo.Instance{}
	repo.GetInstancesFromDBWithPage(&Instances, page)
	return Instances
}
