package models

import "go-todo/config"

func GetAllTasks(tasks *[]Task) error {
	if err := config.DB.Find(&tasks).Error; err != nil {
		return err
	}
	return nil
}

func CreateNewTask(task *Task) error {
	if err := config.DB.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func GetTask(task *Task, id string) error {
	if err := config.DB.Where("id = ?", id).First(task).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTask(task *Task) error {
	if err := config.DB.Save(task).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTask(task *Task) error {
	if err := config.DB.Delete(task).Error; err != nil {
		return err
	}
	return nil
}
