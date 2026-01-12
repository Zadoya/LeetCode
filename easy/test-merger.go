package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type MergerOld struct {
	usersSubjctsIgnore map[uint32]int
	ignoreSubjects     [][]uint32
	abUsers            map[uint32]struct{}
}

type userInfo struct {
	ignoreSubjectsIdx int
	abUser            bool
}

type MergerNew struct {
	userInfo       map[uint32]userInfo
	ignoreSubjects [][]uint32
}

func measureMemory(f func()) uint64 {
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)
	f()
	runtime.ReadMemStats(&m2)
	return m2.Alloc - m1.Alloc
}

func generateSubjects(numSets, subjectsPerSet int) [][]uint32 {
	subjects := make([][]uint32, numSets)
	for i := 0; i < numSets; i++ {
		subj := make([]uint32, subjectsPerSet)
		for j := range subj {
			subj[j] = uint32(rand.Intn(10000))
		}
		subjects[i] = subj
	}
	return subjects
}

func createMergerOld(numUsers, numUsersAb, subjectsPerGroup, stardId int) func() {
	return func() {
		m := &MergerOld{
			usersSubjctsIgnore: make(map[uint32]int, numUsers),
			ignoreSubjects:     generateSubjects(numUsers, subjectsPerGroup),
			abUsers:            make(map[uint32]struct{}, numUsersAb),
		}

		for i := 0; i < numUsers; i++ {
			userID := uint32(i)
			m.usersSubjctsIgnore[userID] = i

		}
		for i := 0; i < numUsersAb; i++ {
			userID := uint32(i)
			m.abUsers[userID] = struct{}{}
		}
	}
}

func createMergerNew(numUsers, numUsersAb, subjectsPerGroup, stardId int) func() {
	return func() {
		m := &MergerNew{
			userInfo:       make(map[uint32]userInfo, numUsers),
			ignoreSubjects: generateSubjects(numUsers, subjectsPerGroup),
		}

		for i := 0; i < numUsers; i++ {
			userID := uint32(i)
			m.userInfo[userID] = userInfo{
				ignoreSubjectsIdx: i,
			}
		}

		for i := 0 + stardId; i < numUsersAb + stardId; i++ {
			userID := uint32(i)
			if info, ok := m.userInfo[userID]; ok {
				info.abUser = true
				m.userInfo[userID] = info
			} else {
				m.userInfo[userID] = userInfo{
					abUser: true,
				}
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numUsersIgnore := 70_000_000
	numUsersAb := 70_000_000
	subjectsPerGroup := 10
	starIdForAbUsers := 0

	oldSize := measureMemory(createMergerOld(numUsersIgnore, numUsersAb, subjectsPerGroup, starIdForAbUsers))
	newSize := measureMemory(createMergerNew(numUsersIgnore, numUsersAb, subjectsPerGroup, starIdForAbUsers))

	fmt.Printf("Users ignore: %d | Users AB: %d | Subjects per Group: %d | AB Ratio: %d\n", numUsersIgnore, numUsersAb, subjectsPerGroup, starIdForAbUsers)
	fmt.Printf("MergerOld uses ~%.2f MB\n", float64(oldSize)/1024/1024.0)
	fmt.Printf("MergerNew uses ~%.2f MB\n", float64(newSize)/1024/1024.0)
}