package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
	LogEntry "userLoggingProject/internal/features/logs/entity"
	"userLoggingProject/internal/features/logs/repository"
)

func main() {

	logRepo := repository.New()

	for userId := 0; userId < 1; userId++ {

		for logId := 0; logId < 20; logId++ {
			log := makeRandomLog()
			_, err := logRepo.Save(strconv.Itoa(userId), log)
			if err != nil {
				continue
			}
		}
	}
	fmt.Println("Select action:")
	fmt.Println("0) exit")
	fmt.Println("1) generate users")
	fmt.Println("2) generate logs")
	fmt.Println("3) print log from user")
	fmt.Println("4) remove log from user")
	var input string

	for {

		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Println("You entered:", input)
		switch input {
		case "1":
			{
				break
			}
		case "2":
			{
				fmt.Println("Get user id:")
				_, err := fmt.Scan(&input)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				lCount := rand.Intn(5000)
				generateAndSaveLog(input, lCount, logRepo)
				fmt.Println(strconv.Itoa(lCount) + " logs has been generated")
			}
		case "3":
			{
				fmt.Println("Get user id:")
				_, err := fmt.Scan(&input)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				userLogs, _ := logRepo.LoadAll(input)
				fmt.Println("Log count is:" + strconv.Itoa(len(userLogs)))
				for i, log := range userLogs {
					fmt.Println("logIndex: " + strconv.Itoa(i) + " time :" + log.Time.String() + log.Message)
				}
			}
		case "4":
			{
				fmt.Println("Get user id:")
				_, err := fmt.Scan(&input)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				err = logRepo.RemoveAll(input)
				if err != nil {
					continue
				}
				all, err := logRepo.LoadAll(input)
				if err != nil {
					return
				}
				fmt.Println("Log count is:" + strconv.Itoa(len(all)))
				for i, log := range all {
					fmt.Println("logIndex: " + strconv.Itoa(i) + " time :" + log.Time.String() + log.Message)
				}
			}
		default:
			break
		}
	}
}

func generateAndSaveLog(userId string, count int, repo repository.LogsRepository) {
	if count < 2000 {
		sl := make([]LogEntry.LogEntry, count)
		for i := 0; i < count; i++ {
			sl = append(sl, *makeRandomLog())
		}
		err := repo.SaveAll(userId, sl)
		if err == nil {
			fmt.Println("saved element count is:" + strconv.Itoa(count))
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {

		go func() {
			id, err := repo.Save(userId, makeRandomLog())
			if err == nil {
				fmt.Println("saved LogId is:" + strconv.Itoa(id))
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func makeRandomLog() *LogEntry.LogEntry {
	return &LogEntry.LogEntry{Time: time.Now(), Message: strconv.Itoa(int(rand.Int63()))}
}
