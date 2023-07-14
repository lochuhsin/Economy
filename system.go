package main

func PeopleMovement(entityManager *EntityManager) {
	_ = entityManager.World
	// for i := 0; i < len(entityManager.DrawPopulation); i++ {
	// 	entityManager.DrawPopulation[i].AccelerationX += (float64(rand.Intn(3)) - 1) * 0.1
	// 	entityManager.DrawPopulation[i].AccelerationY += (float64(rand.Intn(3)) - 1) * 0.1

	// 	entityManager.DrawPopulation[i].VelocityX += entityManager.DrawPopulation[i].AccelerationX
	// 	entityManager.DrawPopulation[i].VelocityY += entityManager.DrawPopulation[i].AccelerationY

	// 	if entityManager.DrawPopulation[i].VelocityX > 2 || entityManager.DrawPopulation[i].VelocityX < -2 {
	// 		entityManager.DrawPopulation[i].AccelerationX *= -1
	// 		entityManager.DrawPopulation[i].VelocityX = 2
	// 	}

	// 	if entityManager.DrawPopulation[i].VelocityY > 2 || entityManager.DrawPopulation[i].VelocityY > -2 {
	// 		entityManager.DrawPopulation[i].AccelerationY *= -1
	// 		entityManager.DrawPopulation[i].VelocityY = 2
	// 	}

	// 	entityManager.DrawPopulation[i].CenterX += entityManager.DrawPopulation[i].VelocityX
	// 	entityManager.DrawPopulation[i].CenterY += entityManager.DrawPopulation[i].VelocityY

	// 	if entityManager.DrawPopulation[i].CenterX >= world.Width || entityManager.DrawPopulation[i].CenterX <= 0 {
	// 		entityManager.DrawPopulation[i].VelocityX *= -1

	// 	}
	// 	if entityManager.DrawPopulation[i].CenterY >= world.Height || entityManager.DrawPopulation[i].CenterY <= 0 {
	// 		entityManager.DrawPopulation[i].VelocityY *= -1
	// 	}
	// }
}

func Transaction(entityManager *EntityManager) {}
