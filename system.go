package main

import (
	"log"
	"math/rand"
	"time"
)

type Operations int

const (
	Save              Operations = 0
	Spend             Operations = 1
	InvestProductivty Operations = 2
)

type SystemManager struct {
	Entities EntityManager
}

func InitSystemManager(entitySettings EntitySettings) SystemManager {
	return SystemManager{
		InitEntityManager(entitySettings),
	}
}

func (s *SystemManager) Run(systemChannel chan EntityManager) {
	initSystemTime := time.Now().Nanosecond()
	for {
		currentSystemTime := time.Now().Nanosecond()
		duration := float64(currentSystemTime-initSystemTime) / 100000000
		PeopleMovement(&s.Entities, duration)
		Transaction(&s.Entities)
		GroupTotalWealth(&s.Entities)

		select {
		case systemChannel <- s.Entities:
			continue
		default:
			log.Println("blocked, channel is system channel is full")
		}
		initSystemTime = currentSystemTime
	}
}

func PeopleMovement(entities *EntityManager, duration float64) {
	world := entities.World
	for i := 0; i < len(entities.GroupPeople); i++ {
		entities.GroupPeople[i].VelocityX += entities.GroupPeople[i].Acceleration * Direction()
		entities.GroupPeople[i].VelocityY += entities.GroupPeople[i].Acceleration * Direction()

		entities.GroupPeople[i].CenterX += entities.GroupPeople[i].VelocityX * duration
		entities.GroupPeople[i].CenterY += entities.GroupPeople[i].VelocityY * duration

		if ExceedBoundary(entities.GroupPeople[i].CenterX, int(world.Width)) {
			entities.GroupPeople[i].VelocityX *= -1
		}

		if ExceedBoundary(entities.GroupPeople[i].CenterY, int(world.Height)) {
			entities.GroupPeople[i].VelocityY *= -1
		}
	}
}

func Transaction(entityManager *EntityManager) {
	population := entityManager.Population

	for i := 0; i < len(population); i++ {
		op := Operations(rand.Intn(3))
		salary := population[i].Productivity
		switch op {
		case Save:
			population[i].Savings += salary
		case InvestProductivty:
			population[i].Productivity += salary
		case Spend:
			continue
		default:
			continue
		}
		population[i].Operations = append(population[i].Operations, op)
	}
}

func GroupTotalWealth(manager *EntityManager) {
	var totalWealth float64
	groupWealth := map[string]float64{}

	for _, p := range manager.Population {
		totalWealth += p.Savings
		groupWealth[p.GroupId] += p.Savings
	}
	for i := 0; i < len(manager.GroupPeople); i++ {
		gid := manager.GroupPeople[i].Id
		manager.GroupPeople[i].GroupWealth = groupWealth[gid]
	}
	manager.World.WorldWealth = totalWealth
}

func ExceedBoundary(position float64, boundary int) bool {
	return position < 0 || position >= float64(boundary)
}

func Direction() float64 {
	if rand.Float32() < 0.5 {
		return -1
	}
	return 1
}
